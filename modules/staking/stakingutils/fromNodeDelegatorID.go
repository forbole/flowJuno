package stakingutils

import (
	"fmt"

	"github.com/rs/zerolog/log"

	"github.com/forbole/flowJuno/modules/utils"
	"github.com/forbole/flowJuno/types"
	"github.com/onflow/cadence"
	"github.com/onflow/flow-go-sdk"

	"github.com/forbole/flowJuno/client"

	database "github.com/forbole/flowJuno/db/postgresql"
)

func GetDataFromNodeDelegatorID(nodeInfo []types.StakerNodeInfo, block *flow.Block, db *database.Db, flowClient client.Proxy) error {

	for _, node := range nodeInfo {
		/* err:=getDelegatorCommitted(node,block,db,flowClient)
		if err!=nil{
			return err
		}
		*/
		_, err := getDelegatorInfo(node, block, db, flowClient)
		if err != nil {
			return err
		}
		/*
			err=getDelegatorRequest(node,block,db,flowClient)
			if err!=nil{
				return err
			}


			err=getDelegatorRewarded(node,block,db,flowClient)
			if err!=nil{
				return err
			}

			err=getDelegatorStaked(node,block,db,flowClient)
			if err!=nil{
				return err
			}

			err=getDelegatorUnstaked(node,block,db,flowClient)
			if err!=nil{
				return err
			}

			err=getDelegatorUnstaking(node,block,db,flowClient)
			if err!=nil{
				return err
			}

			err=getDelegatorUnstakingRequest(node,block,db,flowClient)
			if err!=nil{
				return err
			}
		*/
	}

	return nil
}

/* func getDelegatorCommitted(nodeInfo types.NodeInfoFromNodeID, block *flow.Block, db *database.Db, flowClient client.Proxy) error {
	log.Trace().Str("module", "staking").Int64("height", int64(block.Height)).
		Msg("updating node unstaking tokens")
	script := fmt.Sprintf(`
	import FlowIDTableStaking from %s
	pub fun main(nodeID: String, delegatorID: UInt32): UFix64 {
	  let delInfo = FlowIDTableStaking.DelegatorInfo(nodeID: nodeID, delegatorID: delegatorID)
	  return delInfo.tokensCommitted
  }`, flowClient.Contract().StakingTable)

	args := []cadence.Value{cadence.NewString(nodeId), cadence.NewUInt32(delegatorID)}
	value, err := flowClient.Client().ExecuteScriptAtLatestBlock(flowClient.Ctx(), []byte(script), args)
	if err != nil {
		return err
	}

	committed, err := utils.CadenceConvertUint64(value)
	if err != nil {
		return err
	}

	return db.SaveDelegatorCommitted(types.NewDelegatorCommitted(committed, block.Height, nodeId, delegatorID))

	return nil
} */

func getDelegatorInfo(nodeInfo types.StakerNodeInfo, block *flow.Block, db *database.Db, flowClient client.Proxy) ([]types.DelegatorNodeInfo, error) {
	log.Trace().Str("module", "staking").Int64("height", int64(block.Height)).
		Msg("updating node unstaking tokens")

	if nodeInfo.DelegatorIDCounter==0{
		return nil,nil
	}
	
	script := fmt.Sprintf(`
	import FlowIDTableStaking from %s
	pub fun main(node:String,begin:UInt32,end:UInt32): [FlowIDTableStaking.DelegatorInfo] {
		let delegatorInfoArray: [FlowIDTableStaking.DelegatorInfo] = []
		
		var i:UInt32 = begin

		while i<end {
			i=i+1
			delegatorInfoArray.append(FlowIDTableStaking.DelegatorInfo(nodeID: node, delegatorID: i))
		}
			
		return delegatorInfoArray
	
  }`, flowClient.Contract().StakingTable)


	var i uint32
	delegatorNum := nodeInfo.DelegatorIDCounter - 1
	fmt.Println(delegatorNum)
	fmt.Println(nodeInfo.Delegators)
	fmt.Println(len(nodeInfo.Delegators))
	fmt.Println(nodeInfo.Id)
	delegatorInfoArray := make([]types.DelegatorNodeInfo, delegatorNum)
	for i = 0; i <= delegatorNum;i = i + 4000 {
		end := i + 4000
		if end > delegatorNum {
			end = delegatorNum
		}
		if end==0{
			end=1
		}
		fmt.Println(fmt.Sprintf("start %d end %d",i,end))
		args := []cadence.Value{cadence.NewString(nodeInfo.Id), cadence.NewUInt32(i), cadence.NewUInt32(end)}
		value, err := flowClient.Client().ExecuteScriptAtLatestBlock(flowClient.Ctx(), []byte(script), args)
		if err != nil {
			return nil, err
		}

		committed, err := types.DelegatorNodeInfoArrayFromCadence(value)
		if err != nil {
			return nil, err
		}
		fmt.Println(committed[0])

		delegatorInfoArray = append(delegatorInfoArray, committed...)
		fmt.Println(delegatorInfoArray[0])

	}
	//fmt.Println(delegatorInfoArray[0])

	
	splittedDelegatorInfos:=utils.SplitDelegatorNodeInfo(delegatorInfoArray,9)

	for _,arr:=range splittedDelegatorInfos{
		fmt.Println(len(arr))
		err:=db.SaveDelegatorInfo(arr, block.Height)
		if err!=nil{
			return nil,err
		}
	}
	return delegatorInfoArray, nil
}

/*
func getDelegatorRequest(nodeInfo types.NodeInfoFromNodeID, block *flow.Block, db *database.Db, flowClient client.Proxy) error {
	log.Trace().Str("module", "staking").Int64("height", int64(block.Height)).
		Msg("updating node unstaking tokens")
	script := fmt.Sprintf(`
	import FlowIDTableStaking from %s
	pub fun main(nodeID: String, delegatorID: UInt32): UFix64 {
	  let delInfo = FlowIDTableStaking.DelegatorInfo(nodeID: nodeID, delegatorID: delegatorID)
	  return delInfo.tokensRequestedToUnstake
  }`, flowClient.Contract().StakingTable)

	args := []cadence.Value{cadence.NewString(nodeId), cadence.NewUInt32(delegatorID)}
	value, err := flowClient.Client().ExecuteScriptAtLatestBlock(flowClient.Ctx(), []byte(script), args)
	if err != nil {
		return err
	}

	committed, err := utils.CadenceConvertUint64(value)
	if err != nil {
		return err
	}

	return db.SaveDelegatorRequest(types.NewDelegatorRequest(committed, int64(block.Height), nodeId, delegatorID))
}

func getDelegatorRewarded(nodeInfo types.NodeInfoFromNodeID, block *flow.Block, db *database.Db, flowClient client.Proxy) error {
	log.Trace().Str("module", "staking").Int64("height", int64(block.Height)).
		Msg("updating node unstaking tokens")
	script := fmt.Sprintf(`
	import FlowIDTableStaking from %s
	pub fun main(node:String,begin:UInt32,end:UInt32): [FlowIDTableStaking.DelegatorInfo] {
		let delegatorInfoArray: [FlowIDTableStaking.DelegatorInfo] = []

		var i:UInt32 = begin

		while i<end {
			i=i+1
			delegatorInfoArray.append(FlowIDTableStaking.DelegatorInfo(nodeID: node, delegatorID: i))
		}

		return delegatorInfoArray

  }`, flowClient.Contract().StakingTable)

	args := []cadence.Value{cadence.NewString(nodeId), cadence.NewUInt32(delegatorID)}
	value, err := flowClient.Client().ExecuteScriptAtLatestBlock(flowClient.Ctx(), []byte(script), args)
	if err != nil {
		return err
	}

	committed, err := utils.CadenceConvertUint64(value)
	if err != nil {
		return err
	}

	return db.SaveDelegatorRewarded(types.NewDelegatorRewarded(committed, int64(block.Height), nodeId, delegatorID))
}

func getDelegatorStaked(nodeInfo types.NodeInfoFromNodeID, block *flow.Block, db *database.Db, flowClient client.Proxy) error {
	log.Trace().Str("module", "staking").Int64("height", int64(block.Height)).
		Msg("updating node unstaking tokens")
	script := fmt.Sprintf(`
	import FlowIDTableStaking from %s
	pub fun main(nodeID: String, delegatorID: UInt32): UFix64 {
	  let delInfo = FlowIDTableStaking.DelegatorInfo(nodeID: nodeID, delegatorID: delegatorID)
	  return delInfo.tokensStaked
  }`, flowClient.Contract().StakingTable)

	args := []cadence.Value{cadence.NewString(nodeId), cadence.NewUInt32(delegatorID)}
	value, err := flowClient.Client().ExecuteScriptAtLatestBlock(flowClient.Ctx(), []byte(script), args)
	if err != nil {
		return err
	}

	staked, err := utils.CadenceConvertUint64(value)
	if err != nil {
		return err
	}

	return db.SaveDelegatorStaked(types.NewDelegatorStaked(staked, int64(block.Height), nodeId, delegatorID))
}

func getDelegatorUnstaked(nodeInfo types.NodeInfoFromNodeID, block *flow.Block, db *database.Db, flowClient client.Proxy) error {
	log.Trace().Str("module", "staking").Int64("height", int64(block.Height)).
		Msg("updating node unstaking tokens")
	script := fmt.Sprintf(`
	import FlowIDTableStaking from %s
	pub fun main(nodeID: String, delegatorID: UInt32): UFix64 {
	  let delInfo = FlowIDTableStaking.DelegatorInfo(nodeID: nodeID, delegatorID: delegatorID)
	  return delInfo.tokensUnstaked
  }`, flowClient.Contract().StakingTable)

	args := []cadence.Value{cadence.NewString(nodeId), cadence.NewUInt32(delegatorID)}
	value, err := flowClient.Client().ExecuteScriptAtLatestBlock(flowClient.Ctx(), []byte(script), args)
	if err != nil {
		return err
	}

	staked, err := utils.CadenceConvertUint64(value)
	if err != nil {
		return err
	}

	return db.SaveDelegatorUnstaked(types.NewDelegatorUnstaked(staked, int64(block.Height), nodeId, delegatorID))
}

func getDelegatorUnstaking(nodeInfo types.NodeInfoFromNodeID, block *flow.Block, db *database.Db, flowClient client.Proxy) error {
	log.Trace().Str("module", "staking").Int64("height", int64(block.Height)).
		Msg("updating node unstaking tokens")
	script := fmt.Sprintf(`
	import FlowIDTableStaking from %s
	pub fun main(nodeID: String, delegatorID: UInt32): UFix64 {
	  let delInfo = FlowIDTableStaking.DelegatorInfo(nodeID: nodeID, delegatorID: delegatorID)
	  return delInfo.tokensUnstaking
  }`, flowClient.Contract().StakingTable)

	args := []cadence.Value{cadence.NewString(nodeId), cadence.NewUInt32(delegatorID)}
	value, err := flowClient.Client().ExecuteScriptAtLatestBlock(flowClient.Ctx(), []byte(script), args)
	if err != nil {
		return err
	}

	staking, err := utils.CadenceConvertUint64(value)
	if err != nil {
		return err
	}

	return db.SaveDelegatorUnstaking(types.NewDelegatorUnstaking(staking, int64(block.Height), nodeId, delegatorID))
}

func getDelegatorUnstakingRequest(nodeInfo types.NodeInfoFromNodeID, block *flow.Block, db *database.Db, flowClient client.Proxy) error {
	log.Trace().Str("module", "staking").Int64("height", int64(block.Height)).
		Msg("updating node unstaking tokens")
	script := fmt.Sprintf(`
	import FlowIDTableStaking from %s
	pub fun main(nodeID: String, delegatorID: UInt32): UFix64 {
	  let delInfo = FlowIDTableStaking.DelegatorInfo(nodeID: nodeID, delegatorID: delegatorID)
	  return delInfo.tokensRequestedToUnstake
  }`, flowClient.Contract().StakingTable)

	args := []cadence.Value{cadence.NewString(nodeId), cadence.NewUInt32(delegatorID)}
	value, err := flowClient.Client().ExecuteScriptAtLatestBlock(flowClient.Ctx(), []byte(script), args)
	if err != nil {
		return err
	}

	staking, err := utils.CadenceConvertUint64(value)
	if err != nil {
		return err
	}

	return db.SaveDelegatorUnstakingRequest(types.NewDelegatorUnstakingRequest(staking, int64(block.Height), nodeId, delegatorID))
}
*/
