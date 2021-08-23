package auth

import (
	"encoding/json"

	database "github.com/forbole/flowJuno/db"

	"github.com/cosmos/cosmos-sdk/simapp/params"
	"github.com/forbole/flowJuno/modules/messages"
	"github.com/forbole/flowJuno/modules/modules"
	"github.com/forbole/flowJuno/types"

	"github.com/forbole/flowJuno/client"
	tmtypes "github.com/tendermint/tendermint/types"
)

var (
	_ modules.Module        = &Module{}
	_ modules.GenesisModule = &Module{}
	_ modules.MessageModule = &Module{}
	
)

// Module represents the x/auth module
type Module struct {
	messagesParser messages.MessageAddressesParser
	encodingConfig *params.EncodingConfig
	flowClient     client.Proxy
	db             *database.Database
}

// NewModule builds a new Module instance
func NewModule(
	messagesParser messages.MessageAddressesParser,
	flowClient client.Proxy,
	encodingConfig *params.EncodingConfig, db *database.Database,
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

// HandleGenesis implements modules.GenesisModule
func (m *Module) HandleGenesis(_ *tmtypes.GenesisDoc, appState map[string]json.RawMessage) error {
	return Handler(appState, m.encodingConfig.Marshaler, m.db)
}

// HandleMsg implements modules.MessageModule
func (m *Module) HandleMsg(index int, msg types.Event, tx *types.Tx) error {
	return HandleMsg(msg, m.messagesParser, m.encodingConfig.Marshaler, m.db, int64(tx.Height), m.flowClient)
}
