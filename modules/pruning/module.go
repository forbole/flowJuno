package pruning

import (
	"fmt"

	"github.com/onflow/flow-go-sdk"
	"github.com/rs/zerolog/log"
	tmctypes "github.com/tendermint/tendermint/rpc/core/types"

	"github.com/forbole/flowJuno/db"
	"github.com/forbole/flowJuno/modules"
	"github.com/forbole/flowJuno/types"
)

var _ modules.Module = &Module{}

// Module represents the pruning module allowing to clean the database periodically
type Module struct {
	cfg types.PruningConfig
	db  db.Database
}

// NewModule builds a new Module instance
func NewModule(cfg types.PruningConfig, db db.Database) *Module {
	return &Module{
		cfg: cfg,
		db:  db,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "pruning"
}

// HandleBlock implements modules.BlockModule
func (m *Module) HandleBlock(block *flow.Block, _ *types.Txs, _ *tmctypes.ResultValidators) error {
	if m.cfg == nil {
		// Nothing to do, pruning is disabled
		return nil
	}

	if int64(block.Height)%m.cfg.GetInterval() != 0 {
		// Not an interval height, so just skip
		return nil
	}

	pruningDb, ok := m.db.(db.PruningDb)
	if !ok {
		return fmt.Errorf("pruning is enabled, but your database does not implement PruningDb")
	}

	// Get last pruned height
	var height, err = pruningDb.GetLastPruned()
	if err != nil {
		return err
	}

	// Iterate from last pruned height until (current block height - keep recent) to
	// avoid pruning the recent blocks that should be kept
	for ; height < int64(block.Height)-m.cfg.GetKeepRecent(); height++ {

		if height%m.cfg.GetKeepEvery() == 0 {
			// The height should be kept, so just skip
			continue
		}

		// Prune the height
		log.Debug().Str("module", "pruning").Int64("height", height).Msg("pruning")
		err = pruningDb.Prune(height)
		if err != nil {
			return fmt.Errorf("error while pruning height %d: %s", height, err.Error())
		}
	}

	return pruningDb.StoreLastPruned(height)
}
