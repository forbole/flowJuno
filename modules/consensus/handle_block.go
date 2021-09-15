package consensus

import (
	"github.com/cosmos/cosmos-sdk/simapp/params"
	"github.com/forbole/flowJuno/modules/messages"
	"github.com/forbole/flowJuno/modules/modules"
	"github.com/forbole/flowJuno/types"
	"github.com/onflow/flow-go-sdk"

	"github.com/forbole/flowJuno/client"
	db "github.com/forbole/flowJuno/db/postgresql"
)

func HandleBlock(block *flow.Block,_ messages.MessageAddressesParser,db *db.Db ,height int64, flowClient client.Proxy)error{
	

}