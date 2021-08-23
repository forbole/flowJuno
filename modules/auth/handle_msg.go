package auth

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/forbole/flowJuno/modules/messages"
	"github.com/forbole/flowJuno/types"
	"github.com/rs/zerolog/log"

	"github.com/forbole/bdjuno/database"
	authutils "github.com/forbole/bdjuno/modules/auth/utils"
	"github.com/forbole/flowJuno/client"

)

// HandleMsg handles any message updating the involved accounts
func HandleMsg(msg types.Event, getAddresses messages.MessageAddressesParser, cdc codec.Marshaler, db *database.Db, height int64, flowClient client.Proxy) error {
	addresses := msg.Value.Fields[0].String()

	return authutils.UpdateAccounts(utils.FilterNonAccountAddresses(addresses), cdc, db, height, flowClient)

}
