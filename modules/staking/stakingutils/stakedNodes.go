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

	err := getNodeTotalCommitment(nodeInfofromNodeId, block, db, flowClient)
	if err != nil {
		return err
	}

	err = getNodeTotalCommitmentWithoutDelegators(nodeInfofromNodeId, block, db, flowClient)
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
			if strings.Contains(err.Error(), "computation limited exceeded: 100000") {
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
