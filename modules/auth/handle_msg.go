package auth

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/forbole/flowJuno/modules/messages"
	"github.com/forbole/flowJuno/types"

	"github.com/forbole/flowJuno/client"
	"github.com/forbole/flowJuno/db"
	authutils "github.com/forbole/flowJuno/modules/auth/utils"
)

// HandleMsg handles any message updating the involved accounts
func HandleMsg(msg types.Event, getAddresses messages.MessageAddressesParser, cdc codec.Marshaler, db *db.Database, height int64, flowClient client.Proxy) error {
	address := msg.Value.Fields[0].String()
	addresses:=[]string{address}

	return authutils.UpdateAccounts(addresses, db, height, flowClient)

}
