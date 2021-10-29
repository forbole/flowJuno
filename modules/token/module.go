package token

import (
	"github.com/cosmos/cosmos-sdk/simapp/params"
	"github.com/forbole/flowJuno/modules/messages"
	"github.com/forbole/flowJuno/modules/modules"
	"github.com/forbole/flowJuno/types"
	"github.com/onflow/flow-go-sdk"

	"github.com/forbole/flowJuno/client"
	db "github.com/forbole/flowJuno/db/postgresql"
)

var (
	_ modules.Module      = &Module{}
	_ modules.BlockModule = &Module{}
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
	return "token"
}

// HandleEvent implements modules.MessageModule
func (m *Module) HandleBlock(block *flow.Block, _ *types.Txs) error {
	return HandleBlock(m.db, block.Height, m.flowClient)
}
