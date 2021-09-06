package auth

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/forbole/flowJuno/modules/messages"
	"github.com/forbole/flowJuno/types"

	"github.com/forbole/flowJuno/client"
	db "github.com/forbole/flowJuno/db/postgresql"
	authutils "github.com/forbole/flowJuno/modules/auth/utils"
)

// HandleMsg handles any message updating the involved accounts
func HandleMsg(msg types.Event, getAddresses messages.MessageAddressesParser, cdc codec.Marshaler, db *db.Db, height int64, flowClient client.Proxy,tx *types.Tx) error {
	addresses,err:=getAddresses(cdc,*tx)
	if err!=nil{
		return err
	}
	fmt.Println("HandleMsg")

	return authutils.UpdateAccounts(addresses, db, height, flowClient)

}
