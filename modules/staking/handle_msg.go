package staking

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/forbole/flowJuno/modules/messages"
	"github.com/forbole/flowJuno/types"

	"github.com/forbole/flowJuno/client"
	db "github.com/forbole/flowJuno/db/postgresql"
	authutils "github.com/forbole/flowJuno/modules/auth/utils"
)

// HandleEvent handles any message updating the involved accounts
func HandleEvent(getAddresses messages.MessageAddressesParser,db *db.Db,  event types.Event,tx *types.Tx, flowClient client.Proxy) error {
	event.Type
}
