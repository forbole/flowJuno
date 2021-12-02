package stakingutils

import (
	"fmt"

	"github.com/rs/zerolog/log"

	"github.com/forbole/flowJuno/types"
	"github.com/onflow/cadence"
	"github.com/onflow/flow-go-sdk"

	"github.com/forbole/flowJuno/client"

	database "github.com/forbole/flowJuno/db/postgresql"
)

// GetDataFromAddresses getting delegator info and node info from address. However, it should be 
// replaced by auth that save the node id with addres
func GetDataFromAddresses(addresses []string, block *flow.Block, db *database.Db, flowClient client.Proxy) error {
	err := getNodeInfoFromAddress(addresses, block, db, flowClient)
	if err != nil {
		return err
	}

	err = getDelegatorInfoFromAddress(addresses, block, db, flowClient)
	if err != nil {
		return err
	}

	return nil
}

// getNodeInfoFromAddress get validator nodes that the addresses control 
func getNodeInfoFromAddress(addresses []string, block *flow.Block, db *database.Db, flowClient client.Proxy) error {
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

	totalStakeArr := make([]types.NodeInfoFromAddress, len(addresses))
	for i, id := range addresses {
		nodeId := []cadence.Value{cadence.BytesToAddress([]byte(id))}
		value, err := flowClient.Client().ExecuteScriptAtLatestBlock(flowClient.Ctx(), []byte(script), nodeId)
		if err != nil {
			return err
		}

		stakingKey, err := types.NewStakerNodeInfoFromCadence(value)
		if err != nil {
			return err
		}

		totalStakeArr[i] = types.NewNodeInfoFromAddress(addresses[i], stakingKey, int64(block.Height))
	}

	return db.SaveNodeInfoFromAddresses(totalStakeArr)
}

// getDelegatorInfoFromAddress get the delegators info that the addresses control
func getDelegatorInfoFromAddress(addresses []string, block *flow.Block, db *database.Db, flowClient client.Proxy) error {
	log.Trace().Str("module", "staking").Int64("height", int64(block.Height)).
		Msg("updating get node networking address")
	script := fmt.Sprintf(`
	import FlowIDTableStaking from %s
	pub fun main(address: Address): FlowIDTableStaking.DelegatorInfo {

	  let account = getAccount(address)
  
	  let delegator = account.getCapability<&{FlowIDTableStaking.NodeDelegatorPublic}>(/public/flowStakingDelegator)!
		  .borrow() ?? panic("Could not borrow reference to delegator object")
  
	  return FlowIDTableStaking.DelegatorInfo(nodeID: delegator.nodeID, delegatorID: delegator.id)
  }`, flowClient.Contract().StakingTable)

	totalStakeArr := make([]types.NodeInfoFromAddress, len(addresses))
	for i, id := range addresses {
		nodeId := []cadence.Value{cadence.BytesToAddress([]byte(id))}
		value, err := flowClient.Client().ExecuteScriptAtLatestBlock(flowClient.Ctx(), []byte(script), nodeId)
		if err != nil {
			return err
		}

		stakingKey, err := types.NewStakerNodeInfoFromCadence(value)
		if err != nil {
			return err
		}

		totalStakeArr[i] = types.NewNodeInfoFromAddress(addresses[i], stakingKey, int64(block.Height))
	}

	return db.SaveNodeInfoFromAddresses(totalStakeArr)
}
