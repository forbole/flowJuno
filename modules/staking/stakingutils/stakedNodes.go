package stakingutils

import (
	"fmt"
	"strings"

	"github.com/rs/zerolog/log"

	"github.com/forbole/flowJuno/types"
	"github.com/onflow/cadence"
	"github.com/onflow/flow-go-sdk"

	"github.com/forbole/flowJuno/client"
	"github.com/forbole/flowJuno/modules/utils"

	database "github.com/forbole/flowJuno/db/postgresql"
)

func GetDataFromNodeID(nodeInfofromNodeId []types.StakerNodeInfo, block *flow.Block, db *database.Db, flowClient client.Proxy) error {
	log.Trace().Str("module", "staking").Int64("height", int64(block.Height)).
		Msg("getting staked node infos")

	err := getNodeUnstakingTokens(nodeInfofromNodeId, block, db, flowClient)
	if err != nil {
		return err
	}

	err = getNodeTotalCommitment(nodeInfofromNodeId, block, db, flowClient)
	if err != nil {
		return err
	}

	err = getNodeTotalCommitmentWithoutDelegators(nodeInfofromNodeId, block, db, flowClient)
	if err != nil {
		return err
	}

	err = getNodeStakingKey(nodeInfofromNodeId, block, db, flowClient)
	if err != nil {
		return err
	}

	err = getNodeStakedTokens(nodeInfofromNodeId, block, db, flowClient)
	if err != nil {
		return err
	}

	err = getNodeRole(nodeInfofromNodeId, block, db, flowClient)
	if err != nil {
		return err
	}

	err = getNodeRewardedTokens(nodeInfofromNodeId, block, db, flowClient)
	if err != nil {
		return err
	}

	err = getNodeNetworkingKey(nodeInfofromNodeId, block, db, flowClient)
	if err != nil {
		return err
	}

	err = getNodeNetworkingAddress(nodeInfofromNodeId, block, db, flowClient)
	if err != nil {
		return err
	}

	err = getNodeInitialWeight(nodeInfofromNodeId, block, db, flowClient)
	if err != nil {
		return err
	}

	err = getNodeCommittedTokens(nodeInfofromNodeId, block, db, flowClient)
	if err != nil {
		return err
	}

	return nil
}

func getNodeUnstakingTokens(nodeInfos []types.StakerNodeInfo, block *flow.Block, db *database.Db, flowClient client.Proxy) error {
	log.Trace().Str("module", "staking").Int64("height", int64(block.Height)).
		Msg("updating node unstaking tokens")

	totalStakeArr := make([]types.NodeUnstakingTokens, len(nodeInfos))
	for i, id := range nodeInfos {
		totalStakeArr[i] = types.NewNodeUnstakingTokens(id.Id, id.TokensUnstaking, int64(block.Height))
	}

	return db.SaveNodeUnstakingTokens(totalStakeArr)
}

func getNodeTotalCommitment(nodeInfos []types.StakerNodeInfo, block *flow.Block, db *database.Db, flowClient client.Proxy) error {
	log.Trace().Str("module", "staking").Int64("height", int64(block.Height)).
		Msg("updating node unstaking tokens")
	script := fmt.Sprintf(`
	import FlowIDTableStaking from %s
	pub fun main(nodeID: String): UFix64 {
	  let nodeInfo = FlowIDTableStaking.NodeInfo(nodeID: nodeID)
	  return nodeInfo.totalCommittedWithDelegators()
  }`, flowClient.Contract().StakingTable)

	totalStakeArr := make([]types.NodeTotalCommitment, len(nodeInfos))
	for i, id := range nodeInfos {
		nodeId := []cadence.Value{cadence.NewString(id.Id)}
		fmt.Println(id.Id)
		value, err := flowClient.Client().ExecuteScriptAtLatestBlock(flowClient.Ctx(), []byte(script), nodeId)
		if err != nil {
			// When validator exist 10000, cadence exceed computation limit. Wonfix atm
			if strings.Contains(err.Error(),"computation limited exceeded: 100000"){
				log.Trace().Str("module", "staking").Int64("height", int64(block.Height)).Err(err)
				continue
			}
			return err
		}

		tokensUnstaking, err := utils.CadenceConvertUint64(value)
		if err != nil {
			return err
		}

		totalStakeArr[i] = types.NewNodeTotalCommitment(id.Id, tokensUnstaking, int64(block.Height))
	}

	return db.SaveNodeTotalCommitment(totalStakeArr)
}

func getNodeTotalCommitmentWithoutDelegators(nodeInfos []types.StakerNodeInfo, block *flow.Block, db *database.Db, flowClient client.Proxy) error {
	log.Trace().Str("module", "staking").Int64("height", int64(block.Height)).
		Msg("updating node unstaking tokens")
	script := fmt.Sprintf(`
	import FlowIDTableStaking from %s
	pub fun main(nodeID: String): UFix64 {
	  let nodeInfo = FlowIDTableStaking.NodeInfo(nodeID: nodeID)
	  return nodeInfo.totalCommittedWithoutDelegators()
  }`, flowClient.Contract().StakingTable)

	totalStakeArr := make([]types.NodeTotalCommitmentWithoutDelegators, len(nodeInfos))
	for i, id := range nodeInfos {
		nodeId := []cadence.Value{cadence.NewString(id.Id)}
		value, err := flowClient.Client().ExecuteScriptAtLatestBlock(flowClient.Ctx(), []byte(script), nodeId)
		if err != nil {
			return err
		}

		tokensUnstaking, err := utils.CadenceConvertUint64(value)
		if err != nil {
			return err
		}

		totalStakeArr[i] = types.NewNodeTotalCommitmentWithoutDelegators(id.Id, tokensUnstaking, int64(block.Height))
	}

	return db.SaveNodeTotalCommitmentWithoutDelegators(totalStakeArr)
}

func getNodeStakingKey(nodeInfos []types.StakerNodeInfo, block *flow.Block, db *database.Db, flowClient client.Proxy) error {
	log.Trace().Str("module", "staking").Int64("height", int64(block.Height)).
		Msg("updating node unstaking tokens")

	totalStakeArr := make([]types.NodeStakingKey, len(nodeInfos))
	for i, id := range nodeInfos {
		stakingKey := id.StakingKey

		totalStakeArr[i] = types.NewNodeStakingKey(id.Id, stakingKey, int64(block.Height))
	}

	return db.SaveNodeStakingKey(totalStakeArr)
}

func getNodeStakedTokens(nodeInfos []types.StakerNodeInfo, block *flow.Block, db *database.Db, flowClient client.Proxy) error {
	log.Trace().Str("module", "staking").Int64("height", int64(block.Height)).
		Msg("updating node unstaking tokens")
	totalStakeArr := make([]types.NodeStakedTokens, len(nodeInfos))
	for i, id := range nodeInfos {
		totalStakeArr[i] = types.NewNodeStakedTokens(id.Id, id.TokensStaked, int64(block.Height))
	}

	return db.SaveNodeStakedTokens(totalStakeArr)
}

func getNodeRole(nodeInfos []types.StakerNodeInfo, block *flow.Block, db *database.Db, flowClient client.Proxy) error {
	log.Trace().Str("module", "staking").Int64("height", int64(block.Height)).
		Msg("updating node role")

	totalStakeArr := make([]types.NodeRole, len(nodeInfos))
	for i, id := range nodeInfos {
		totalStakeArr[i] = types.NewNodeRole(id.Id, id.Role, int64(block.Height))
	}

	return db.SaveNodeRole(totalStakeArr)
}

func getNodeRewardedTokens(nodeInfos []types.StakerNodeInfo, block *flow.Block, db *database.Db, flowClient client.Proxy) error {
	log.Trace().Str("module", "staking").Int64("height", int64(block.Height)).
		Msg("updating get node rewarded tokens")

	totalStakeArr := make([]types.NodeRewardedTokens, len(nodeInfos))
	for i, id := range nodeInfos {
		totalStakeArr[i] = types.NewNodeRewardedTokens(id.Id, id.TokensRewarded, int64(block.Height))
	}

	return db.SaveNodeRewardedTokens(totalStakeArr)
}

func getNodeNetworkingKey(nodeInfos []types.StakerNodeInfo, block *flow.Block, db *database.Db, flowClient client.Proxy) error {
	log.Trace().Str("module", "staking").Int64("height", int64(block.Height)).
		Msg("updating get node networking key")
	totalStakeArr := make([]types.NodeNetworkingKey, len(nodeInfos))
	for i, id := range nodeInfos {
		totalStakeArr[i] = types.NewNodeNetworkingKey(id.Id, id.NetworkingKey, int64(block.Height))
	}

	return db.SaveNodeNetworkingKey(totalStakeArr)
}

func getNodeNetworkingAddress(nodeInfos []types.StakerNodeInfo, block *flow.Block, db *database.Db, flowClient client.Proxy) error {
	log.Trace().Str("module", "staking").Int64("height", int64(block.Height)).
		Msg("updating get node networking address")
	totalStakeArr := make([]types.NodeNetworkingAddress, len(nodeInfos))
	for i, id := range nodeInfos {

		totalStakeArr[i] = types.NewNodeNetworkingAddress(id.Id, id.NetworkingAddress, int64(block.Height))
	}

	return db.SaveNodeNetworkingAddress(totalStakeArr)
}
func getNodeInitialWeight(nodeInfos []types.StakerNodeInfo, block *flow.Block, db *database.Db, flowClient client.Proxy) error {
	log.Trace().Str("module", "staking").Int64("height", int64(block.Height)).
		Msg("updating get node networking address")
	totalStakeArr := make([]types.NodeInitialWeight, len(nodeInfos))
	for i, id := range nodeInfos {
		totalStakeArr[i] = types.NewNodeInitialWeight(id.Id, id.InitialWeight, int64(block.Height))
	}

	return db.SaveNodeInitialWeight(totalStakeArr)
}

func getNodeCommittedTokens(nodeInfos []types.StakerNodeInfo, block *flow.Block, db *database.Db, flowClient client.Proxy) error {
	log.Trace().Str("module", "staking").Int64("height", int64(block.Height)).
		Msg("updating get node networking address")
	totalStakeArr := make([]types.NodeCommittedTokens, len(nodeInfos))
	for i, id := range nodeInfos {
		totalStakeArr[i] = types.NewNodeCommittedTokens(id.Id, id.TokensCommitted, int64(block.Height))
	}

	return db.SaveNodeCommittedTokens(totalStakeArr)
}
