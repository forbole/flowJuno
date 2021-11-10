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

	if len(totalStakeArr)==0{
		return nil
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
		value, err := flowClient.Client().ExecuteScriptAtLatestBlock(flowClient.Ctx(), []byte(script), nodeId)
		if err != nil {
			// When validator exist 10000, cadence exceed computation limit. It need to calculate in raw
			if strings.Contains(err.Error(), "computation limited exceeded: 100000") {
				fmt.Println(id.Id)
				nodeTotalCommitment, err := getNodeTotalCommitmentRaw(id, block, db, flowClient)
				if err != nil {
					return err
				}
				totalStakeArr[i] = *nodeTotalCommitment
				continue
			}
			return err
		}

		totalCommit, err := utils.CadenceConvertUint64(value)
		if err != nil {
			return err
		}

		totalStakeArr[i] = types.NewNodeTotalCommitment(id.Id, totalCommit, int64(block.Height))
	}

	if len(totalStakeArr)==0{
		return nil
	}

	return db.SaveNodeTotalCommitment(totalStakeArr)
}

// getNodeTotalCommitmentRaw add up all delegator's delegatorFullCommittedBalance in a node
func getNodeTotalCommitmentRaw(nodeInfo types.StakerNodeInfo, block *flow.Block, db *database.Db, flowClient client.Proxy) (*types.NodeTotalCommitment, error) {
	script := fmt.Sprintf(`
	import FlowIDTableStaking from %s
	pub fun main(node:String,begin:UInt32,end:UInt32): UFix64{
		var totalCommit:UFix64=UFix64(0)
		
		var i:UInt32 = begin

		if (begin==0){
			let nodeInfo = FlowIDTableStaking.NodeInfo(nodeID: node)
			totalCommit=totalCommit+nodeInfo.totalCommittedWithoutDelegators()
		}

		while i<=end {
			i=i+1

			var delegator=FlowIDTableStaking.DelegatorInfo(nodeID: node, delegatorID: i)
			if (delegator.tokensCommitted + delegator.tokensStaked) < delegator.tokensRequestedToUnstake{
				totalCommit=totalCommit+UFix64(0)
			}else{
				totalCommit=totalCommit+delegator.tokensCommitted+delegator.tokensStaked-delegator.tokensRequestedToUnstake
			}
		}
			
		return totalCommit
	
  }`, flowClient.Contract().StakingTable)

	var i uint32
	delegatorNum := nodeInfo.DelegatorIDCounter - 1
	var totalCommit uint64
	for i = 0; i < delegatorNum; i = i + 4001 {
		end := i + 4000
		if end > delegatorNum {
			end = delegatorNum
		}
		if end == 0 {
			end = 1
		}
		args := []cadence.Value{cadence.NewString(nodeInfo.Id), cadence.NewUInt32(i), cadence.NewUInt32(end)}
		value, err := flowClient.Client().ExecuteScriptAtLatestBlock(flowClient.Ctx(), []byte(script), args)
		if err != nil {
			return nil, err
		}

		val, err := utils.CadenceConvertUint64(value)
		if err != nil {
			return nil, err
		}

		totalCommit += val

	}
	nodeTotalCommitment := types.NewNodeTotalCommitment(nodeInfo.Id, totalCommit, int64(block.Height))

	return &nodeTotalCommitment, nil
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

	if len(totalStakeArr)==0{
		return nil
	}

	return db.SaveNodeTotalCommitmentWithoutDelegators(totalStakeArr)
}
