package staking

import (
	"fmt"

	"github.com/go-co-op/gocron"
	"github.com/rs/zerolog/log"

	"github.com/forbole/flowJuno/modules/staking/stakingutils"

	"github.com/forbole/flowJuno/types"

	"github.com/onflow/flow-go-sdk"

	"github.com/forbole/flowJuno/client"
	"github.com/forbole/flowJuno/modules/utils"

	database "github.com/forbole/flowJuno/db/postgresql"
	db "github.com/forbole/flowJuno/db/postgresql"
)

func RegisterPeriodicOps(scheduler *gocron.Scheduler, db *database.Db, flowClient client.Proxy) error {
	log.Debug().Str("module", "staking").Msg("setting up periodic tasks")

	if _, err := scheduler.Every(1).Week().Tuesday().At("15:00").StartImmediately().Do(func() {
		utils.WatchMethod(func() error { return HandleStaking(db, flowClient) })
	}); err != nil {
		return err
	}

	return HandleStaking(db, flowClient)
}

func HandleStaking(db *db.Db, flowClient client.Proxy) error {
	block, err := flowClient.Client().GetLatestBlock(flowClient.Ctx(), false)
	if err != nil {
		return err
	}

	addresses, err := db.GetAddresses()
	if err != nil {
		return err
	}
	_, err = getTable(block, db, flowClient)
	if err != nil {
		return err
	}

	nodeInfo, err := getNodeInfosFromTable(block, db, flowClient)
	if err != nil {
		return err
	}

	if len(addresses) != 0 {
		err = stakingutils.GetDataFromAddresses(addresses, block, db, flowClient)
		if err != nil {
			return err
		}
	}

	err = stakingutils.GetDataWithNoArgs(block, db, int64(block.Height), flowClient)
	if err != nil {
		return err
	}

	err = stakingutils.GetDataFromNodeID(nodeInfo, block, db, flowClient)
	if err != nil {
		return err
	}

	err = stakingutils.GetDataFromNodeDelegatorID(nodeInfo, block, db, flowClient)
	if err != nil {
		return err
	}

	return nil
}

func getTable(block *flow.Block, db *database.Db, flowClient client.Proxy) ([]string, error) {
	log.Trace().Str("module", "staking").Int64("height", int64(block.Height)).
		Msg("updating get Staked Node id per block")

	script := fmt.Sprintf(`
	import FlowIDTableStaking from %s
	pub fun main(): [String] {
	  return FlowIDTableStaking.getNodeIDs()
  }`, flowClient.Contract().StakingTable)

	value, err := flowClient.Client().ExecuteScriptAtLatestBlock(flowClient.Ctx(), []byte(script), nil)
	if err != nil {
		return nil, err
	}

	table, err := utils.CadenceConvertStringArray(value)
	if err != nil {
		return nil, err
	}

	if len(table)==0{
		return nil,nil
	}

	return table, db.SaveStakingTable(types.NewStakingTable(int64(block.Height), table))

}

func getNodeInfosFromTable(block *flow.Block, db *database.Db, flowClient client.Proxy) ([]types.StakerNodeInfo, error) {

	script := fmt.Sprintf(`
	import FlowIDTableStaking from %s
	pub fun main(): [FlowIDTableStaking.NodeInfo] {
		var nodeIds=FlowIDTableStaking.getNodeIDs()
		let nodeInfoArray: [FlowIDTableStaking.NodeInfo] = []
		for node in nodeIds{
			nodeInfoArray.append(FlowIDTableStaking.NodeInfo(nodeID: node))
		}
		return nodeInfoArray
    }`, flowClient.Contract().StakingTable)

	value, err := flowClient.Client().ExecuteScriptAtLatestBlock(flowClient.Ctx(), []byte(script), nil)
	if err != nil {
		return nil, err
	}

	stakingKeys, err := types.NewStakerNodeInfoArrayFromCadence(value)
	if err != nil {
		return nil, err
	}

	if len(stakingKeys)==0{
		return nil,nil
	}

	return stakingKeys, db.SaveNodeInfosFromTable(stakingKeys, block.Height)
}

func getAddressesFromAccounts(accounts []types.Account) []string {
	addresses := make([]string, len(accounts))
	for i, account := range accounts {
		addresses[i] = account.Address
	}
	return addresses
}
