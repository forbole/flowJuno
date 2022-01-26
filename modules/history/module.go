package history

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/forbole/flowJuno/modules"
	"github.com/forbole/flowJuno/modules/messages"
	"github.com/forbole/flowJuno/types/config"

	"github.com/forbole/flowJuno/db/db"
)

const (
	moduleName = "history"
)

var (
	_ modules.Module                   = &Module{}
	_ modules.PeriodicOperationsModule = &Module{}
	_ modules.GenesisModule            = &Module{}
	_ modules.MessageModule            = &Module{}
)

// Module represents the module that allows to store historic information
type Module struct {
	cfg config.ChainConfig
	cdc codec.Marshaler
	db  *database.Db

	getAddresses messages.MessageAddressesParser
}

// NewModule allows to build a new Module instance
func NewModule(cfg config.ChainConfig, messagesParser messages.MessageAddressesParser, cdc codec.Marshaler, db *database.Db) *Module {
	return &Module{
		cfg:          cfg,
		cdc:          cdc,
		db:           db,
		getAddresses: messagesParser,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return moduleName
}
