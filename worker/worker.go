package worker

import (
	"fmt"
	"strings"
	"time"

	"github.com/forbole/flowJuno/logging"
	"github.com/onflow/flow-go-sdk"

	tmjson "github.com/tendermint/tendermint/libs/json"

	"github.com/cosmos/cosmos-sdk/simapp/params"

	"github.com/forbole/flowJuno/modules/modules"

	"github.com/rs/zerolog/log"
	tmos "github.com/tendermint/tendermint/libs/os"
	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/forbole/flowJuno/client"
	"github.com/forbole/flowJuno/db"
	"github.com/forbole/flowJuno/types"
)

// Worker defines a job consumer that is responsible for getting and
// aggregating block and associated data and exporting it to a database.
type Worker struct {
	index int

	queue          types.HeightQueue
	encodingConfig *params.EncodingConfig
	cp             *client.Proxy
	db             db.Database
	modules        []modules.Module
	logger         logging.Logger
}

// NewWorker allows to create a new Worker implementation.
func NewWorker(index int, config *Config) Worker {
	return Worker{
		index:          index,
		encodingConfig: config.EncodingConfig,
		cp:             config.ClientProxy,
		queue:          config.Queue,
		db:             config.Database,
		modules:        config.Modules,
		logger:         config.Logger,
	}
}

// Start starts a worker by listening for new jobs (block heights) from the
// given worker queue. Any failed job is logged and re-enqueued.
func (w Worker) Start() {
	logging.WorkerCount.Inc()

	for i := range w.queue {
		if err := w.process(i); err != nil {
			// re-enqueue any failed job
			// TODO: Implement exponential backoff or max retries for a block height when the block is not generated.
			if err != nil && strings.Contains(err.Error(), "could not retrieve resource: key not found") {
				sleeptime:=1
				for err != nil && strings.Contains(err.Error(), "could not retrieve resource: key not found") {
					//If it cannot find the key, retry until can parse
						time.Sleep(time.Second*time.Duration(sleeptime))
						err = w.process(i)
						sleeptime=sleeptime*2
					}

			} else {
				WaitUntilQueueAvailable(w.queue)

				go func() {
					log.Error().Err(err).Int64("height", i).Msg("re-enqueueing failed block",)
					w.queue <- i
				}()
			}

			logging.WorkerHeight.WithLabelValues(fmt.Sprintf("%d", w.index)).Set(float64(i))
		}
	}
}

//nolint:gocyclo
// process defines the job consumer workflow. It will fetch a block for a given
// height and associated metadata and export it to a database. It returns an
// error if any export process fails.
// To get all transaction and event from the block, follow the order so that wont double call:
// block -> collection_grauntee -> transaction -> event
func (w Worker) process(height int64) error {
	exists, err := w.db.HasBlock(height)
	if err != nil {
		return err
	}

	if exists {
		log.Debug().Int64("height", height).Msg("skipping already exported block")
		return nil
	}

	// To get all transaction and event from the block, follow the order so that wont double call:
	// block -> collection_grauntee -> transaction -> event
	if height%int64(w.db.GetPartitionSize()) == 0 {
		err = w.db.CreatePartition("transaction", uint64(height))
		if err != nil {
			return err
		}
		err = w.db.CreatePartition("transaction_result", uint64(height))
		if err != nil {
			return err
		}
		err = w.db.CreatePartition("event", uint64(height))
		if err != nil {
			return err
		}
	}

	block, err := w.cp.Block(height)
	if err != nil {
		//log.Error().Err(err).Int64("height", height).Msg("failed to get block")
		return err
	}

	if height == int64(w.cp.GetGenesisHeight()) {
		return w.HandleGenesis(block)
	}

	txs, err := w.cp.Txs(block)
	if err != nil {
		log.Error().Err(err).Int64("height", height).Msg("failed to get transaction Result for block")
		return err
	}

	// Call the block handlers
	for _, module := range w.modules {
		if blockModule, ok := module.(modules.BlockModule); ok {
			err = blockModule.HandleBlock(block, &txs)
			if err != nil {
				w.logger.BlockError(module, block, err)
				return err
			}
		}
	}

	err = w.ExportBlock(block)
	if err != nil {
		return err
	}

	collections, err := w.ExportCollection(block)
	if err != nil {
		return err
	}

	err = w.ExportTx(&txs)
	if err != nil {
		return err
	}

	if len(collections) == 0 {
		return nil
	}

	var transactionIDs []flow.Identifier
	for _, collection := range collections {
		transactionIDs = append(transactionIDs, (collection.TransactionIds)...)
	}

	return w.ExportTransactionResult(transactionIDs, height)

}

func (w Worker) ExportTransactionResult(txids []flow.Identifier, height int64) error {
	txResults, err := w.cp.TransactionResult(txids)
	if len(txResults) == 0 {
		return nil
	}
	if err != nil {
		return err
	}
	return w.db.SaveTransactionResult(txResults, uint64(height))
}

func (w Worker) ExportCollection(block *flow.Block) ([]types.Collection, error) {
	collections := w.cp.Collections(block)
	if len(collections) == 0 {
		return nil, nil
	}
	err := w.db.SaveCollection(collections)
	if err != nil {
		return nil, err
	}
	return collections, nil
}

// getGenesisFromRPC returns the genesis read from the RPC endpoint
func (w Worker) getGenesisFromRPC() (*tmtypes.GenesisDoc, error) {
	return nil, fmt.Errorf("Not implenment genesisi from grpc")
	/* log.Debug().Msg("getting genesis")
	response, err := w.cp.Genesis()
	if err != nil {
		log.Error().Err(err).Msg("failed to get genesis")
		return nil, err
	}
	return response.Genesis, nil */
}

// getGenesisFromFilePath tries reading the genesis doc from the given path
func (w Worker) getGenesisFromFilePath(path string) (*tmtypes.GenesisDoc, error) {
	log.Debug().Str("path", path).Msg("reading genesis from file")

	bz, err := tmos.ReadFile(path)
	if err != nil {
		log.Error().Err(err).Msg("failed to read genesis file")
		return nil, err
	}

	var genDoc tmtypes.GenesisDoc
	err = tmjson.Unmarshal(bz, &genDoc)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal genesis doc")
		return nil, err
	}

	return &genDoc, nil
}

// ExportBlock accepts a finalized block and a corresponding set of transactions
// and persists them to the database along with attributable metadata. An error
// is returned if the write fails.
func (w Worker) ExportBlock(b *flow.Block) error {
	// Save all validators
	/* err := w.SaveNodeInfos(vals.NodeInfos)
	if err != nil {
		return err
	} */

	// Save the block
	err := w.db.SaveBlock(b)
	if err != nil {
		log.Error().Err(err).Int64("height", int64(b.BlockHeader.Height)).Msg("failed to persist block")
		return err
	}

	return nil
}

// ExportTxs accepts a slice of transactions and persists then inside the database.
// An error is returned if the write fails.
func (w Worker) ExportTx(txs *types.Txs) error {
	// Handle all the transactions inside the block
	err := w.db.SaveTxs(*txs)
	if err != nil {
		log.Error().Err(err).Int64("height", int64((*txs)[0].Height)).Msg("failed to export txs")
		return err
	}

	//Handle all event
	var allEventInTx []types.Event
	for _, tx := range *txs {
		events, err := w.cp.EventsInTransaction(tx)
		if err != nil {
			log.Error().Err(err).Int64("height", int64(tx.Height)).Msg("failed to get events for block")
			return err
		}
		allEventInTx = append(allEventInTx, events...)

		//Handle event with associated tx
		for _, event := range events {
			for _, module := range w.modules {
				if messageModule, ok := module.(modules.MessageModule); ok {
					err = messageModule.HandleEvent(event.Height, event, &tx)
					if err != nil {
						w.logger.EventsError(module, &event, err)
						return err
					}
				}
			}
		}
	}

	err = w.db.SaveEvents(allEventInTx)
	if err != nil {
		return err
	}

	for _, tx := range *txs {
		for _, module := range w.modules {
			if transactionModule, ok := module.(modules.TransactionModule); ok {
				err = transactionModule.HandleTx(int(tx.Height), &tx)
				if err != nil {
					w.logger.TxError(module, &tx, err)
					return err
				}
			}
		}
	}

	return nil
}

// HandleGenesis accepts a GenesisDoc and calls all the registered genesis handlers
// in the order in which they have been registered.
func (w Worker) HandleGenesis(block *flow.Block) error {
	// Call the genesis handlers
	for _, module := range w.modules {
		if genesisModule, ok := module.(modules.GenesisModule); ok {
			if err := genesisModule.HandleGenesis(block, w.cp.GetChainID()); err != nil {
				w.logger.GenesisError(module, err)
			}
		}
	}
	return nil
}
