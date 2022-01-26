package history

import (
	"github.com/forbole/flowJuno/modules"
	"github.com/forbole/juno/v2/types/config"

	db "github.com/forbole/flowJuno/db/postgresql"
)

const (
	moduleName = "history"
)

var (
	_ modules.Module                   = &Module{}
)

// Module represents the module that allows to store historic information
type Module struct {
	cfg config.ChainConfig
	db  *db.Db
}

// NewModule allows to build a new Module instance
func NewModule(cfg config.ChainConfig,  db *db.Db) *Module {
	return &Module{
		cfg:          cfg,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return moduleName
}
