package accounts

import (
	"encoding/json"

	"github.com/forbole/flowJuno/db"
	"github.com/forbole/flowJuno/types"


	"github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/desmos-labs/juno/modules"
	"github.com/desmos-labs/juno/modules/messages"
	juno "github.com/desmos-labs/juno/types"
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
	db             *database.Db
}

// NewModule builds a new Module instance
func NewModule(
	messagesParser messages.MessageAddressesParser,
	authClient authtypes.QueryClient,
	encodingConfig *params.EncodingConfig, db *database.Db,
) *Module {
	return &Module{
		messagesParser: messagesParser,
		db:             db,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "auth"
}

// HandleMsg implements modules.MessageModule
func (m *Module) HandleEvent(_ int, msg types.Event, tx *juno.Tx) error {
	return HandleEvent(event types.Event, m.messagesParser, m.db, tx.Height)
}
