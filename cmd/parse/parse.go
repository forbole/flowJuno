package parse

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/go-co-op/gocron"

	"github.com/forbole/flowJuno/client"
	"github.com/forbole/flowJuno/logging"
	"github.com/forbole/flowJuno/modules/modules"
	"github.com/forbole/flowJuno/types"
	"github.com/forbole/flowJuno/worker"

	"github.com/forbole/flowJuno/db"

	"github.com/spf13/cobra"
)

var (
	waitGroup sync.WaitGroup
)

// ParseCmd returns the command that should be run when we want to start parsing a chain state.
func ParseCmd(cmdCfg *Config) *cobra.Command {
	return &cobra.Command{
		Use:     "parse",
		Short:   "Start parsing the blockchain data",
		PreRunE: types.ConcatCobraCmdFuncs(ReadConfig(cmdCfg)),
		RunE: func(cmd *cobra.Command, args []string) error {
			parserData, err := SetupParsing(cmdCfg)
			if err != nil {
				return err
			}

			// Run all the additional operations
			for _, module := range parserData.Modules {
				if module, ok := module.(modules.AdditionalOperationsModule); ok {
					err = module.RunAdditionalOperations()
					if err != nil {
						return err
					}
				}
			}

			return StartParsing(parserData)
		},
	}
}

// StartParsing represents the function that should be called when the parse command is executed
func StartParsing(data *ParserData) error {
	// Get the config
	cfg := types.Cfg.GetParsingConfig()
	logging.StartHeight.Add(float64(cfg.GetStartHeight()))

	// Start periodic operations
	scheduler := gocron.NewScheduler(time.UTC)
	for _, module := range data.Modules {
		if module, ok := module.(modules.PeriodicOperationsModule); ok {
			err := module.RegisterPeriodicOperations(scheduler)
			if err != nil {
				return err
			}
		}
	}
	scheduler.StartAsync()

	// Create a queue that will collect, aggregate, and export blocks and metadata
	exportQueue := types.NewQueue(25)

	// Create workers
	config := worker.NewConfig(exportQueue, data.EncodingConfig, data.Proxy, data.Database, data.Modules, data.Logger)
	workers := make([]worker.Worker, cfg.GetWorkers(), cfg.GetWorkers())
	for i := range workers {
		workers[i] = worker.NewWorker(i, config)
	}

	waitGroup.Add(1)

	// Run all the async operations
	for _, module := range data.Modules {
		if module, ok := module.(modules.AsyncOperationsModule); ok {
			go module.RunAsyncOperations()
		}
	}
	// Start each blocking worker in a go-routine where the worker consumes jobs
	// off of the export queue.
	for i, w := range workers {
		data.Logger.Debug("starting worker...", "number", i+1)
		go w.Start()
	}

	// Listen for and trap any OS signal to gracefully shutdown and exit
	trapSignal(data.Proxy, data.Database, data.Logger)

	if cfg.ShouldParseGenesis() {
		// Add the genesis to the queue if requested
		exportQueue <- 0
	}

	if cfg.ShouldParseOldBlocks() {
		go enqueueMissingBlocks(exportQueue, data)
	}

	if cfg.ShouldParseNewBlocks() {
		//go startNewBlockListener(exportQueue, data)
	}

	// Block main process (signal capture will call WaitGroup's Done)
	waitGroup.Wait()
	return nil
}

// enqueueMissingBlocks enqueues jobs (block heights) for missed blocks starting
// at the startHeight up until the latest known height.
func enqueueMissingBlocks(exportQueue types.HeightQueue, data *ParserData) {
	// Get the config
	cfg := types.Cfg.GetParsingConfig()

	// Get the latest height
	latestBlockHeight, err := data.Proxy.LatestHeight()
	if err != nil {
		panic(fmt.Errorf("failed to get last block from RPC client: %s", err))
	}

	if cfg.UseFastSync() {
		data.Logger.Info("fast sync is enabled, ignoring all previous blocks", "latest_block_height", latestBlockHeight)
		for _, module := range data.Modules {
			if mod, ok := module.(modules.FastSyncModule); ok {
				err := mod.DownloadState(latestBlockHeight)
				if err != nil {
					data.Logger.Error("error while performing fast sync",
						"err", err,
						"last_block_height", latestBlockHeight,
						"module", module.Name(),
					)
				}
			}
		}
	} else {
		data.Logger.Info("syncing missing blocks...", "latest_block_height", latestBlockHeight)
		for i := cfg.GetStartHeight(); i <= latestBlockHeight; {
			if len(exportQueue) < 25 {
				data.Logger.Debug("enqueueing missing block", "height", i)

				exportQueue <- i
				i++
				latestBlockHeight++
			}
		}
	}
}

/*
// startNewBlockListener subscribes to new block events via the Tendermint RPC
// and enqueues each new block height onto the provided queue. It blocks as new
// blocks are incoming.
func startNewBlockListener(exportQueue types.HeightQueue, data *ParserData) {
	eventCh, cancel, err := data.Proxy.SubscribeNewBlocks(types.Cfg.GetRPCConfig().GetClientName() + "-blocks")
	defer cancel()

	if err != nil {
		log.Fatal().Err(fmt.Errorf("failed to subscribe to new blocks: %s", err))
	}

	log.Info().Msg("listening for new block events...")

	for e := range eventCh {
		newBlock := e.Data.(tmtypes.EventDataNewBlock).Block
		height := newBlock.Header.Height

		log.Debug().Int64("height", height).Msg("enqueueing new block")
		exportQueue <- height
	}
} */

// trapSignal will listen for any OS signal and invoke Done on the main
// WaitGroup allowing the main process to gracefully exit.
func trapSignal(cp *client.Proxy, db db.Database, logger logging.Logger) {
	var sigCh = make(chan os.Signal)

	signal.Notify(sigCh, syscall.SIGTERM)
	signal.Notify(sigCh, syscall.SIGINT)

	go func() {
		sig := <-sigCh
		logger.Info("caught signal; shutting down...", "signal", sig.String())
		defer cp.Stop()
		defer db.Close()
		defer waitGroup.Done()
	}()
}
