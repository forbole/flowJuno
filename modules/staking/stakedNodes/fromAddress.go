package staking

import (
	"fmt"

	"github.com/rs/zerolog/log"

	"github.com/forbole/flowJuno/types"
	"github.com/onflow/cadence"
	"github.com/onflow/flow-go-sdk"

	"github.com/forbole/flowJuno/client"

	database "github.com/forbole/flowJuno/db/postgresql"
)

func getNodeInfoFromNodeIDs(nodeIds []string, block *flow.Block, db *database.Db, flowClient client.Proxy) error {
	log.Trace().Str("module", "staking").Int64("height", int64(block.Height)).
		Msg("updating get node networking address")
	script := fmt.Sprintf(`
	import FlowIDTableStaking from %s
	pub fun main(address: Address): FlowIDTableStaking.NodeInfo {

	  let account = getAccount(address)
  
	  let nodeStaker = account.getCapability<&{FlowIDTableStaking.NodeStakerPublic}>(FlowIDTableStaking.NodeStakerPublicPath)!
		  .borrow() ?? panic("Could not borrow reference to node staker object")
  
	  return FlowIDTableStaking.NodeInfo(nodeID: nodeStaker.id)
  }`, flowClient.Contract().StakingTable)

	totalStakeArr := make([]types.NodeInfoFromAddress, len(nodeIds))
	for i, id := range nodeIds {
		nodeId := []cadence.Value{cadence.NewString(id)}
		value, err := flowClient.Client().ExecuteScriptAtLatestBlock(flowClient.Ctx(), []byte(script), nodeId)
		if err != nil {
			return err
		}

		stakingKey, err := types.NewStakerNodeInfoFromCadence(value)
		if err != nil {
			return err
		}

		totalStakeArr[i] = types.NewNodeInfoFromAddress(nodeIds[i], stakingKey, int64(block.Height))
	}

	return db.SaveNodeInfoFromAddresses(totalStakeArr)
}
