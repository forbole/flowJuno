package auth

import (
	"github.com/cosmos/cosmos-sdk/simapp/params"
	"github.com/forbole/flowJuno/modules/messages"
	"github.com/forbole/flowJuno/modules/modules"
	"github.com/forbole/flowJuno/types"

	"github.com/forbole/flowJuno/client"
	db "github.com/forbole/flowJuno/db/postgresql"
	"github.com/go-co-op/gocron"
)

var (
	_ modules.Module            = &Module{}
	_ modules.TransactionModule = &Module{}
)

// Module represents the x/auth module
type Module struct {
	messagesParser messages.MessageAddressesParser
	encodingConfig *params.EncodingConfig
	flowClient     client.Proxy
	db             *db.Db
}

// NewModule builds a new Module instance
func NewModule(
	messagesParser messages.MessageAddressesParser,
	flowClient client.Proxy,
	encodingConfig *params.EncodingConfig, db *db.Db,
) *Module {
	return &Module{
		messagesParser: messagesParser,
		encodingConfig: encodingConfig,
		flowClient:     flowClient,
		db:             db,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "auth"
}

// HandleEvent implements modules.MessageModule
func (m *Module) HandleTx(index int, tx *types.Tx) error {
	return HandleTxs(m.messagesParser, m.encodingConfig.Marshaler, m.db, m.flowClient, tx)
}

// RegisterPeriodicOperations implements modules.Module
func (m *Module) RegisterPeriodicOperations(scheduler *gocron.Scheduler) error {
	return nil//RegisterPeriodicOps(scheduler, m.db, m.flowClient)
}
