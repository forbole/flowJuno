package worker

import (
	"fmt"

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
	queue          types.HeightQueue
	encodingConfig *params.EncodingConfig
	cp             *client.Proxy
	db             db.Database
	modules        []modules.Module
}

// NewWorker allows to create a new Worker implementation.
func NewWorker(config *Config) Worker {
	return Worker{
		encodingConfig: config.EncodingConfig,
		cp:             config.ClientProxy,
		queue:          config.Queue,
		db:             config.Database,
		modules:        config.Modules,
	}
}

// Start starts a worker by listening for new jobs (block heights) from the
// given worker queue. Any failed job is logged and re-enqueued.
func (w Worker) Start() {

	for i := range w.queue {
		if err := w.process(i); err != nil {
			// re-enqueue any failed job
			// TODO: Implement exponential backoff or max retries for a block height.
			go func() {
				log.Error().Err(err).Int64("height", i).Msg("re-enqueueing failed block")
				w.queue <- i
			}()
		}
	}
}

// process defines the job consumer workflow. It will fetch a block for a given
// height and associated metadata and export it to a database. It returns an
// error if any export process fails.
func (w Worker) process(height int64) error {
	exists, err := w.db.HasBlock(height)
	if err != nil {
		return err
	}

	if exists {
		log.Debug().Int64("height", height).Msg("skipping already exported block")
		return nil
	}
	/*
		if height == 0 {
			cfg := types.Cfg.GetParsingConfig()
			var genesis *tmtypes.GenesisDoc
			if strings.TrimSpace(cfg.GetGenesisFilePath()) != "" {
				genesis, err = w.getGenesisFromFilePath(cfg.GetGenesisFilePath())
				if err != nil {
					return err
				}
			} else {
				genesis, err = w.getGenesisFromRPC()
				if err != nil {
					return err
				}
			}

			return w.HandleGenesis(genesis)
		}
	*/
	//log.Debug().Int64("height", height).Msg("processing block")

	block, err := w.cp.Block(height)
	if err != nil {
		log.Error().Err(err).Int64("height", height).Msg("failed to get block")
		return err
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
				fmt.Println(err)
			}
		}
	}


	return w.ExportBlock(block, &txs)
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

func (w Worker) SaveNodeInfos(vals []*types.NodeInfo) error {
	err := w.db.SaveNodeInfos(vals)
	if err != nil {
		return fmt.Errorf("error while saving node infos: %s", err)
	}
	return nil
}

// ExportBlock accepts a finalized block and a corresponding set of transactions
// and persists them to the database along with attributable metadata. An error
// is returned if the write fails.
func (w Worker) ExportBlock(b *flow.Block, txs *types.Txs) error {
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

	return w.ExportTx(txs)
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
	for _,tx:=range *txs{
		events, err := w.cp.EventsInTransaction(tx)
		if err != nil {
			log.Error().Err(err).Int64("height",int64(tx.Height)).Msg("failed to get events for block")
			return err
		}
		allEventInTx=append(allEventInTx, events...)

		//Handle event with associated tx
		for _, event := range events {
			for _, module := range w.modules {
				if messageModule, ok := module.(modules.MessageModule); ok {
					err = messageModule.HandleEvent(int(event.Height), event, &tx)
					if err != nil {
						return err
					}
				}
			}
		}
	}

	err=w.db.SaveEvents(allEventInTx)
	if err!=nil{
		return err
	}


	for _,tx:=range *txs{
	for _, module := range w.modules {
		if transactionModule, ok := module.(modules.TransactionModule); ok {
			err = transactionModule.HandleTx(int(tx.Height),&tx)
			if err != nil {
				//w.logger.TxError(module, tx, err)
				return fmt.Errorf("Cannot Parse Transaction")
			}
		}
	}
}

	return nil
}


// HandleGenesis accepts a GenesisDoc and calls all the registered genesis handlers
// in the order in which they have been registered.
func (w Worker) HandleGenesis(genesisHeight int64) error {
	// Call the genesis handlers
	for _, module := range w.modules {
		if genesisModule, ok := module.(modules.GenesisModule); ok {
			if err := genesisModule.HandleGenesis(genesisHeight); err != nil {
				//w.logger.GenesisError(module, err)
			}
		}
	}
	return nil
}