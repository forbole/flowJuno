package staking

import (
	"fmt"

	"github.com/rs/zerolog/log"

	"github.com/forbole/flowJuno/modules/messages"
	"github.com/forbole/flowJuno/types"
	"github.com/forbole/flowJuno/modules/staking/stakedNodes"

	"github.com/onflow/flow-go-sdk"

	"github.com/forbole/flowJuno/client"
	"github.com/forbole/flowJuno/modules/utils"

	database "github.com/forbole/flowJuno/db/postgresql"
	db "github.com/forbole/flowJuno/db/postgresql"
	"github.com/onflow/cadence"
)

func HandleAdditionalOperation(block *flow.Block, _ messages.MessageAddressesParser, db *db.Db, height int64, flowClient client.Proxy) error {
	
	accounts,err:=db.GetAccounts()
	if err!=nil{
		return err
	}
	nodeIds,err:=getTable(block, db, flowClient)
	if err!=nil{
		return err
	}

	NodeInfos,err:=getNodeInfoFromNodeID(nodeIds,block, db, flowClient)
	if err!=nil{
		return err
	}

	addresses :=make([]string,len(accounts))
	for i,account:=range accounts{
		addresses[i]= account.Address
	}
	
	err=staking.GetDataFromAddresses(addresses,block, db, flowClient)
	if err!=nil{
		return err
	}
	
	err = staking.GetDataWithNoArgs(block, db, int64(block.Height), flowClient)
	if err!=nil{
		return err
	}

	err=staking.GetDataFromNodeID(block, db, flowClient)

	
	err=staking.GetDataFromNodeDelegatorID(block, db, flowClient)


	return nil
}

func getTable(block *flow.Block, db *database.Db, flowClient client.Proxy) ([]string,error) {
	log.Trace().Str("module", "staking").Int64("height", int64(block.Height)).
		Msg("updating get Staked Node id per block")

	script := fmt.Sprintf(`
	import FlowIDTableStaking from %s
	pub fun main(): [String] {
	  return FlowIDTableStaking.getNodeIDs()
  }`, flowClient.Contract().StakingTable)

	value, err := flowClient.Client().ExecuteScriptAtLatestBlock(flowClient.Ctx(), []byte(script), nil)
	if err != nil {
		return nil,err
	}

	table, err := utils.CadenceConvertStringArray(value)
	if err != nil {
		return nil,err
	}

	return table,db.SaveStakingTable(types.NewStakingTable(int64(block.Height), table))

}


func getNodeInfoFromNodeID(nodeIds []string, block *flow.Block, db *database.Db, flowClient client.Proxy) ([]types.NodeInfoFromNodeID,error) {
	log.Trace().Str("module", "staking").Int64("height", int64(block.Height)).
		Msg("updating get node networking address")
	script := fmt.Sprintf(`
	import FlowIDTableStaking from %s
	pub fun main(nodeID: String): FlowIDTableStaking.NodeInfo {
	  return FlowIDTableStaking.NodeInfo(nodeID: nodeID)
  }`, flowClient.Contract().StakingTable)

	totalStakeArr := make([]types.NodeInfoFromNodeID, len(nodeIds))
	for i, id := range nodeIds {
		nodeId := []cadence.Value{cadence.NewString(id)}
		value, err := flowClient.Client().ExecuteScriptAtLatestBlock(flowClient.Ctx(), []byte(script), nodeId)
		if err != nil {
			return nil,err
		}

		stakingKey, err := types.NewStakerNodeInfoFromCadence(value)
		if err != nil {
			return nil,err
		}

		totalStakeArr[i] = types.NewNodeInfoFromNodeID(nodeIds[i], stakingKey, int64(block.Height))
	}

	return totalStakeArr,db.SaveNodeInfoFromNodeIDs(totalStakeArr)
}