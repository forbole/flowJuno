package db

import (
	"github.com/cosmos/cosmos-sdk/simapp/params"

	"github.com/forbole/flowJuno/types"
	"github.com/onflow/flow-go-sdk"
)

// Database represents an abstract database that can be used to save data inside it
type Database interface {
	// HasBlock tells whether or not the database has already stored the block having the given height.
	// An error is returned if the operation fails.
	HasBlock(height int64) (bool, error)

	// SaveBlock will be called when a new block is parsed, passing the block itself
	// and the transactions contained inside that block.
	// An error is returned if the operation fails.
	// NOTE. For each transaction inside txs, SaveTx will be called as well.
	SaveBlock(block *flow.Block) error

	// SaveTx will be called to save each transaction contained inside a block.
	// An error is returned if the operation fails.
	SaveTxs(txs types.Txs) error

	// HasValidator returns true if a given validator by consensus address exists.
	// An error is returned if the operation fails.
	HasValidator(address string) (bool, error)

	// SaveValidators stores a list of validators if they do not already exist.
	// An error is returned if the operation fails.
	SaveValidators(validators []*types.Validator) error

	// SaveNodeInfo stores a list of Node Operator Info if they do not already exist.
	// An error is returned if the operation fails.
	SaveNodeInfos(vals []*types.StakerNodeInfo) error

	// SaveCommitSignatures stores a  slice of validator commit signatures.
	// An error is returned if the operation fails.
	SaveCommitSignatures(signatures []*types.CommitSig) error

	// SaveMessage stores a single message.
	// An error is returned if the operation fails.
	SaveMessage(msg *types.Message) error

	// SaveEvent store an array of event emitted in a block.
	// An error is returned if the operation fails.
	SaveEvents(events []types.Event) error

	SaveCollection(collection []types.Collection) error

	SaveTransactionResult(txResults []types.TransactionResult, height uint64) error

	//Partition
	CreatePartition(table string, height int64, heightRange int64) error

	DropPartition(name string) error

	// Close closes the connection to the database
	Close()
}

// PruningDb represents a database that supports pruning properly
type PruningDb interface {
	// Prune prunes the data for the given height, returning any error
	Prune(height int64) error

	// StoreLastPruned saves the last height at which the database was pruned
	StoreLastPruned(height int64) error

	// GetLastPruned returns the last height at which the database was pruned
	GetLastPruned() (int64, error)
}

// Builder represents a method that allows to build any database from a given codec and configuration
type Builder func(cfg types.Config, encodingConfig *params.EncodingConfig) (Database, error)
