package staking

import (
	"fmt"

	"github.com/rs/zerolog/log"

	"github.com/forbole/flowJuno/types"
	"github.com/onflow/cadence"
	"github.com/onflow/flow-go-sdk"

	"github.com/forbole/flowJuno/client"
	"github.com/forbole/flowJuno/modules/utils"

	database "github.com/forbole/flowJuno/db/postgresql"
)

func getDelegatorCommitted(nodeId string, delegatorID uint32, block *flow.Block, db *database.Db, flowClient client.Proxy) error {
	log.Trace().Str("module", "staking").Int64("height", int64(block.Height)).
		Msg("updating node unstaking tokens")
	script := fmt.Sprintf(`
	import FlowIDTableStaking from %s
	pub fun main(nodeID: String, delegatorID: UInt32): UFix64 {
	  let delInfo = FlowIDTableStaking.DelegatorInfo(nodeID: nodeID, delegatorID: delegatorID)
	  return delInfo.tokensCommitted
  }`, flowClient.Contract().StakingTable)

	arg := []cadence.Value{cadence.NewString(nodeId), cadence.NewUInt32(delegatorID)}
	value, err := flowClient.Client().ExecuteScriptAtLatestBlock(flowClient.Ctx(), []byte(script), arg)
	if err != nil {
		return err
	}

	committed, err := utils.CadenceConvertUint64(value)
	if err != nil {
		return err
	}

	return db.SaveDelegatorCommitted(types.NewDelegatorCommitted(committed, block.Height, nodeId, delegatorID))
}
