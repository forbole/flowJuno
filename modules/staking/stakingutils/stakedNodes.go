package stakingutils

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



func GetDataFromNodeID(block *flow.Block, db *database.Db, flowClient client.Proxy) error {
	log.Trace().Str("module", "staking").Int64("height", int64(block.Height)).
		Msg("getting staked node infos")

	ids, err := getCurrentTable(block, db, flowClient)
	if err != nil {
		return err
	}

	err = getNodeUnstakingTokens(ids, block, db, flowClient)
	if err != nil {
		return err
	}

	err = getNodeTotalCommitment(ids, block, db, flowClient)
	if err!=nil{
		return err
	}
	
	err = getNodeTotalCommitmentWithoutDelegators(ids, block, db, flowClient)
	if err!=nil{
		return err
	}

	err = getNodeStakingKey(ids, block, db, flowClient)
	if err!=nil{
		return err
	}

	err = getNodeStakedTokens(ids, block, db, flowClient)
	if err!=nil{
		return err
	}

	err = getNodeRole(ids, block, db, flowClient)
	if err!=nil{
		return err
	}

	err = getNodeRewardedTokens(ids, block, db, flowClient)
	if err!=nil{
		return err
	}

	err = getNodeNetworkingKey(ids, block, db, flowClient)
	if err!=nil{
		return err
	}

	err = getNodeNetworkingAddress(ids, block, db, flowClient)
	if err!=nil{
		return err
	}

	err = getNodeInitialWeight(ids, block, db, flowClient)
	if err!=nil{
		return err
	}

	err = getNodeCommittedTokens(ids, block, db, flowClient)
	if err!=nil{
		return err
	}

	return nil
}



func getNodeUnstakingTokens(nodeIds []string, block *flow.Block, db *database.Db, flowClient client.Proxy) error {
	log.Trace().Str("module", "staking").Int64("height", int64(block.Height)).
		Msg("updating node unstaking tokens")
	script := fmt.Sprintf(`
	import FlowIDTableStaking from %s
	pub fun main(nodeID: String): UFix64 {
	  let nodeInfo = FlowIDTableStaking.NodeInfo(nodeID: nodeID)
	  return nodeInfo.tokensUnstaking
  }`, flowClient.Contract().StakingTable)

	totalStakeArr := make([]types.NodeUnstakingTokens, len(nodeIds))
	for i, id := range nodeIds {
		nodeId := []cadence.Value{cadence.NewString(id)}
		value, err := flowClient.Client().ExecuteScriptAtLatestBlock(flowClient.Ctx(), []byte(script), nodeId)
		if err != nil {
			return err
		}

		tokensUnstaking, err := utils.CadenceConvertUint64(value)
		if err != nil {
			return err
		}

		totalStakeArr[i] = types.NewNodeUnstakingTokens(nodeIds[i], tokensUnstaking, int64(block.Height))
	}

	return db.SaveNodeUnstakingTokens(totalStakeArr)
}

func getNodeTotalCommitment(nodeIds []string, block *flow.Block, db *database.Db, flowClient client.Proxy) error {
	log.Trace().Str("module", "staking").Int64("height", int64(block.Height)).
		Msg("updating node unstaking tokens")
	script := fmt.Sprintf(`
	import FlowIDTableStaking from %s
	pub fun main(nodeID: String): UFix64 {
	  let nodeInfo = FlowIDTableStaking.NodeInfo(nodeID: nodeID)
	  return nodeInfo.totalCommittedWithDelegators()
  }`, flowClient.Contract().StakingTable)

	totalStakeArr := make([]types.NodeTotalCommitment, len(nodeIds))
	for i, id := range nodeIds {
		nodeId := []cadence.Value{cadence.NewString(id)}
		value, err := flowClient.Client().ExecuteScriptAtLatestBlock(flowClient.Ctx(), []byte(script), nodeId)
		if err != nil {
			return err
		}

		tokensUnstaking, err := utils.CadenceConvertUint64(value)
		if err != nil {
			return err
		}

		totalStakeArr[i] = types.NewNodeTotalCommitment(nodeIds[i], tokensUnstaking, int64(block.Height))
	}

	return db.SaveNodeTotalCommitment(totalStakeArr)
}

func getNodeTotalCommitmentWithoutDelegators(nodeIds []string, block *flow.Block, db *database.Db, flowClient client.Proxy) error {
	log.Trace().Str("module", "staking").Int64("height", int64(block.Height)).
		Msg("updating node unstaking tokens")
	script := fmt.Sprintf(`
	import FlowIDTableStaking from %s
	pub fun main(nodeID: String): UFix64 {
	  let nodeInfo = FlowIDTableStaking.NodeInfo(nodeID: nodeID)
	  return nodeInfo.totalCommittedWithoutDelegators()
  }`, flowClient.Contract().StakingTable)

	totalStakeArr := make([]types.NodeTotalCommitmentWithoutDelegators, len(nodeIds))
	for i, id := range nodeIds {
		nodeId := []cadence.Value{cadence.NewString(id)}
		value, err := flowClient.Client().ExecuteScriptAtLatestBlock(flowClient.Ctx(), []byte(script), nodeId)
		if err != nil {
			return err
		}

		tokensUnstaking, err := utils.CadenceConvertUint64(value)
		if err != nil {
			return err
		}

		totalStakeArr[i] = types.NewNodeTotalCommitmentWithoutDelegators(nodeIds[i], tokensUnstaking, int64(block.Height))
	}

	return db.SaveNodeTotalCommitmentWithoutDelegators(totalStakeArr)
}

func getNodeStakingKey(nodeIds []string, block *flow.Block, db *database.Db, flowClient client.Proxy) error {
	log.Trace().Str("module", "staking").Int64("height", int64(block.Height)).
		Msg("updating node unstaking tokens")
	script := fmt.Sprintf(`
	import FlowIDTableStaking from %s
	pub fun main(nodeID: String): String {
	  let nodeInfo = FlowIDTableStaking.NodeInfo(nodeID: nodeID)
	  return nodeInfo.stakingKey
  }`, flowClient.Contract().StakingTable)

	totalStakeArr := make([]types.NodeStakingKey, len(nodeIds))
	for i, id := range nodeIds {
		nodeId := []cadence.Value{cadence.NewString(id)}
		value, err := flowClient.Client().ExecuteScriptAtLatestBlock(flowClient.Ctx(), []byte(script), nodeId)
		if err != nil {
			return err
		}

		stakingKey := value.String()

		totalStakeArr[i] = types.NewNodeStakingKey(nodeIds[i], stakingKey, int64(block.Height))
	}

	return db.SaveNodeStakingKey(totalStakeArr)
}

func getNodeStakedTokens(nodeIds []string, block *flow.Block, db *database.Db, flowClient client.Proxy) error {
	log.Trace().Str("module", "staking").Int64("height", int64(block.Height)).
		Msg("updating node unstaking tokens")
	script := fmt.Sprintf(`
	import FlowIDTableStaking from %s
	pub fun main(nodeID: String): UFix64 {
	  let nodeInfo = FlowIDTableStaking.NodeInfo(nodeID: nodeID)
	  return nodeInfo.tokensStaked
  }`, flowClient.Contract().StakingTable)

	totalStakeArr := make([]types.NodeStakedTokens, len(nodeIds))
	for i, id := range nodeIds {
		nodeId := []cadence.Value{cadence.NewString(id)}
		value, err := flowClient.Client().ExecuteScriptAtLatestBlock(flowClient.Ctx(), []byte(script), nodeId)
		if err != nil {
			return err
		}

		stakingKey, err := utils.CadenceConvertUint64(value)
		if err != nil {
			return err
		}

		totalStakeArr[i] = types.NewNodeStakedTokens(nodeIds[i], stakingKey, int64(block.Height))
	}

	return db.SaveNodeStakedTokens(totalStakeArr)
}

func getNodeRole(nodeIds []string, block *flow.Block, db *database.Db, flowClient client.Proxy) error {
	log.Trace().Str("module", "staking").Int64("height", int64(block.Height)).
		Msg("updating node unstaking tokens")
	script := fmt.Sprintf(`
	import FlowIDTableStaking from %s
	pub fun main(nodeID: String): UInt8 {
	  let nodeInfo = FlowIDTableStaking.NodeInfo(nodeID: nodeID)
	  return nodeInfo.role
  }`, flowClient.Contract().StakingTable)

	totalStakeArr := make([]types.NodeRole, len(nodeIds))
	for i, id := range nodeIds {
		nodeId := []cadence.Value{cadence.NewString(id)}
		value, err := flowClient.Client().ExecuteScriptAtLatestBlock(flowClient.Ctx(), []byte(script), nodeId)
		if err != nil {
			return err
		}

		stakingKey, err := utils.CadenceConvertUint8(value)
		if err != nil {
			return err
		}

		totalStakeArr[i] = types.NewNodeRole(nodeIds[i], stakingKey, int64(block.Height))
	}

	return db.SaveNodeRole(totalStakeArr)
}

func getNodeRewardedTokens(nodeIds []string, block *flow.Block, db *database.Db, flowClient client.Proxy) error {
	log.Trace().Str("module", "staking").Int64("height", int64(block.Height)).
		Msg("updating get node rewarded tokens")
	script := fmt.Sprintf(`
	import FlowIDTableStaking from %s
	pub fun main(nodeID: String): UFix64 {
	  let nodeInfo = FlowIDTableStaking.NodeInfo(nodeID: nodeID)
	  return nodeInfo.tokensRewarded
  }`, flowClient.Contract().StakingTable)

	totalStakeArr := make([]types.NodeRewardedTokens, len(nodeIds))
	for i, id := range nodeIds {
		nodeId := []cadence.Value{cadence.NewString(id)}
		value, err := flowClient.Client().ExecuteScriptAtLatestBlock(flowClient.Ctx(), []byte(script), nodeId)
		if err != nil {
			return err
		}

		stakingKey, err := utils.CadenceConvertUint64(value)
		if err != nil {
			return err
		}

		totalStakeArr[i] = types.NewNodeRewardedTokens(nodeIds[i], stakingKey, int64(block.Height))
	}

	return db.SaveNodeRewardedTokens(totalStakeArr)
}

func getNodeNetworkingKey(nodeIds []string, block *flow.Block, db *database.Db, flowClient client.Proxy) error {
	log.Trace().Str("module", "staking").Int64("height", int64(block.Height)).
		Msg("updating get node networking key")
	script := fmt.Sprintf(`
	import FlowIDTableStaking from %s
	pub fun main(nodeID: String): String {
	  let nodeInfo = FlowIDTableStaking.NodeInfo(nodeID: nodeID)
	  return nodeInfo.networkingKey
  }`, flowClient.Contract().StakingTable)

	totalStakeArr := make([]types.NodeNetworkingKey, len(nodeIds))
	for i, id := range nodeIds {
		nodeId := []cadence.Value{cadence.NewString(id)}
		value, err := flowClient.Client().ExecuteScriptAtLatestBlock(flowClient.Ctx(), []byte(script), nodeId)
		if err != nil {
			return err
		}

		stakingKey := value.String()
		if err != nil {
			return err
		}

		totalStakeArr[i] = types.NewNodeNetworkingKey(nodeIds[i], stakingKey, int64(block.Height))
	}

	return db.SaveNodeNetworkingKey(totalStakeArr)
}

func getNodeNetworkingAddress(nodeIds []string, block *flow.Block, db *database.Db, flowClient client.Proxy) error {
	log.Trace().Str("module", "staking").Int64("height", int64(block.Height)).
		Msg("updating get node networking address")
	script := fmt.Sprintf(`
	import FlowIDTableStaking from %s
	pub fun main(nodeID: String): String {
	  let nodeInfo = FlowIDTableStaking.NodeInfo(nodeID: nodeID)
	  return nodeInfo.networkingAddress
  }`, flowClient.Contract().StakingTable)

	totalStakeArr := make([]types.NodeNetworkingAddress, len(nodeIds))
	for i, id := range nodeIds {
		nodeId := []cadence.Value{cadence.NewString(id)}
		value, err := flowClient.Client().ExecuteScriptAtLatestBlock(flowClient.Ctx(), []byte(script), nodeId)
		if err != nil {
			return err
		}

		stakingKey := value.String()
		if err != nil {
			return err
		}

		totalStakeArr[i] = types.NewNodeNetworkingAddress(nodeIds[i], stakingKey, int64(block.Height))
	}

	return db.SaveNodeNetworkingAddress(totalStakeArr)
}
func getNodeInitialWeight(nodeIds []string, block *flow.Block, db *database.Db, flowClient client.Proxy) error {
	log.Trace().Str("module", "staking").Int64("height", int64(block.Height)).
		Msg("updating get node networking address")
	script := fmt.Sprintf(`
	import FlowIDTableStaking from %s
	pub fun main(nodeID: String): UInt64{
	  let nodeInfo = FlowIDTableStaking.NodeInfo(nodeID: nodeID)
	  return nodeInfo.initialWeight
  }`, flowClient.Contract().StakingTable)

	totalStakeArr := make([]types.NodeInitialWeight, len(nodeIds))
	for i, id := range nodeIds {
		nodeId := []cadence.Value{cadence.NewString(id)}
		value, err := flowClient.Client().ExecuteScriptAtLatestBlock(flowClient.Ctx(), []byte(script), nodeId)
		if err != nil {
			return err
		}

		stakingKey, err := utils.CadenceConvertUint64(value)
		if err != nil {
			return err
		}

		totalStakeArr[i] = types.NewNodeInitialWeight(nodeIds[i], stakingKey, int64(block.Height))
	}

	return db.SaveNodeInitialWeight(totalStakeArr)
}


func getNodeCommittedTokens(nodeIds []string, block *flow.Block, db *database.Db, flowClient client.Proxy) error {
	log.Trace().Str("module", "staking").Int64("height", int64(block.Height)).
		Msg("updating get node networking address")
	script := fmt.Sprintf(`
	import FlowIDTableStaking from %s
	pub fun main(nodeID: String): UFix64{
	  let nodeInfo = FlowIDTableStaking.NodeInfo(nodeID: nodeID)
	  return nodeInfo.tokensCommitted
  }`, flowClient.Contract().StakingTable)

	totalStakeArr := make([]types.NodeCommittedTokens, len(nodeIds))
	for i, id := range nodeIds {
		nodeId := []cadence.Value{cadence.NewString(id)}
		value, err := flowClient.Client().ExecuteScriptAtLatestBlock(flowClient.Ctx(), []byte(script), nodeId)
		if err != nil {
			return err
		}

		stakingKey, err := utils.CadenceConvertUint64(value)
		if err != nil {
			return err
		}

		totalStakeArr[i] = types.NewNodeCommittedTokens(nodeIds[i], stakingKey, int64(block.Height))
	}

	return db.SaveNodeCommittedTokens(totalStakeArr)
}
