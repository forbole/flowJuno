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

		_, err := getDelegatorInfo(node, block, db, flowClient)
		if err != nil {
			return err
		}

	}

	return nil
}

func getDelegatorInfo(nodeInfo types.StakerNodeInfo, block *flow.Block, db *database.Db, flowClient client.Proxy) ([]types.DelegatorNodeInfo, error) {
	log.Trace().Str("module", "staking").Int64("height", int64(block.Height)).
		Msg("updating node unstaking tokens")

	if nodeInfo.DelegatorIDCounter == 0 {
		return nil, nil
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
	var delegatorInfoArray []types.DelegatorNodeInfo
	for i = 0; i <= delegatorNum; i = i + 4000 {
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

		committed, err := types.DelegatorNodeInfoArrayFromCadence(value)
		if err != nil {
			return nil, err
		}

		delegatorInfoArray = append(delegatorInfoArray, committed...)

	}

	splittedDelegatorInfos := utils.SplitDelegatorNodeInfo(delegatorInfoArray, 9)

	for _, arr := range splittedDelegatorInfos {
		fmt.Println(len(arr))
		err := db.SaveDelegatorInfo(arr, block.Height)
		if err != nil {
			return nil, err
		}
	}
	return delegatorInfoArray, nil
}
