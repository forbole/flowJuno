package stakingutils

import (
	"fmt"

	"github.com/rs/zerolog/log"

	"github.com/forbole/flowJuno/types"

	"github.com/forbole/flowJuno/client"
	"github.com/forbole/flowJuno/modules/utils"

	db "github.com/forbole/flowJuno/db/postgresql"
	"github.com/onflow/cadence"
)

func GetDataWithNoArgs( db *db.Db, height int64, flowClient client.Proxy) error {
	payout,err := getWeeklyPayout(height, flowClient)
	if err != nil {
		return fmt.Errorf("Fail to get data with no arg:%s",err)
	}

	err=db.SaveWeeklyPayout(*payout)
	if err!=nil{
		return fmt.Errorf("Fail to get data with no arg:%s",err)
	}

	totalstake,err := getTotalStake(height, flowClient)
	if err != nil {
		return err
	}

	err=db.SaveTotalStake(*totalstake)
	if err!=nil{
		return fmt.Errorf("Fail to get data with no arg:%s",err)
	}

	totalStakeByType,err := getTotalStakeByType(height, flowClient)
	if err!=nil{
		return fmt.Errorf("Fail to get data with no arg:%s",err)
	}

	err=db.SaveTotalStakeByType(totalStakeByType)
	if err!=nil{
		return fmt.Errorf("Fail to get data with no arg:%s",err)
	}

	table, err := getCurrentTable(height, flowClient)
	if err!=nil{
		return fmt.Errorf("Fail to get data with no arg:%s",err)
	}

	err=db.SaveCurrentTable(*table)
	if err!=nil{
		return fmt.Errorf("Fail to get data with no arg:%s",err)
	}

	err = db.SaveCurrentTable(*table)
	if err!=nil{
		return fmt.Errorf("Fail to get data with no arg:%s",err)
	}


	stakeRequirement,err := getStakeRequirements(height, flowClient)
	if err!=nil{
		return fmt.Errorf("Fail to get data with no arg:%s",err)
	}

	err=db.SaveStakeRequirements(stakeRequirement)
	if err!=nil{
		return fmt.Errorf("Fail to get data with no arg:%s",err)
	}

	proposedTable,err := getProposedTable(height, flowClient)
	if err!=nil{
		return fmt.Errorf("Fail to get data with no arg:%s",err)
	}

	err=db.SaveProposedTable(*proposedTable)
	if err!=nil{
		return fmt.Errorf("Fail to get data with no arg:%s",err)
	}

	cutPercentage,err := getCutPercentage(height, flowClient)
	if err != nil {
		return err
	}

	err=db.SaveCutPercentage(*cutPercentage)
	if err!=nil{
		return fmt.Errorf("Fail to get data with no arg:%s",err)
	}

	return nil
}

// getCurrentTable get FlowIDTableStaking.getStakedNodeIDs() at the latest height
func getCurrentTable(height int64,flowClient client.Proxy) (*types.CurrentTable, error) {
	log.Trace().Str("module", "staking").Int64("height", int64(height)).
		Msg("Getting staked node ids")

	script := fmt.Sprintf(`
	import FlowIDTableStaking from %s
	pub fun main(): [String] {
		return FlowIDTableStaking.getStakedNodeIDs()
	}`, flowClient.Contract().StakingTable)

	value, err := flowClient.Client().ExecuteScriptAtLatestBlock(flowClient.Ctx(), []byte(script), nil)
	if err != nil {
		return nil, err
	}

	ids, err := utils.CadenceConvertStringArray(value)
	if err != nil {
		return nil, err
	}

	table:=types.NewCurrentTable(int64(height), ids)

	return &table, nil
}

// getWeeklyPayout get weekly payout for each epoch
func getWeeklyPayout(height int64,flowClient client.Proxy) (*types.WeeklyPayout,error) {
	log.Trace().Str("module", "staking").Int64("height", int64(height)).
		Msg("updating get weekly epoch payout")

	script := fmt.Sprintf(`
	import FlowIDTableStaking from %s
	pub fun main(): UFix64 {
	  return FlowIDTableStaking.getEpochTokenPayout()
  }`, flowClient.Contract().StakingTable)

	value, err := flowClient.Client().ExecuteScriptAtLatestBlock(flowClient.Ctx(), []byte(script), nil)
	if err != nil {
		return nil,err
	}

	payout, err := utils.CadenceConvertUint64(value)
	if err != nil {
		return nil,err
	}

	p:=types.NewWeeklyPayout(int64(height), payout)

	return &p,err
}

func getTotalStake(height int64,flowClient client.Proxy) (*types.TotalStake,error) {
	log.Trace().Str("module", "staking").Int64("height", int64(height)).
		Msg("updating get total stake by type")

	script := fmt.Sprintf(`
	import FlowIDTableStaking from %s
	pub fun main(): UFix64 {
	  let stakedTokens = FlowIDTableStaking.getTotalTokensStakedByNodeType()
  
	  // calculate the total number of tokens staked
	  var totalStaked: UFix64 = 0.0
	  for nodeType in stakedTokens.keys {
		  // Do not count access nodes
		  if nodeType != UInt8(5) {
			  totalStaked = totalStaked + stakedTokens[nodeType]!
		  }
	  }
  
	  return totalStaked
  }`, flowClient.Contract().StakingTable)

	value, err := flowClient.Client().ExecuteScriptAtLatestBlock(flowClient.Ctx(), []byte(script), nil)
	if err != nil {
		return nil,err
	}

	totalStake, err := utils.CadenceConvertUint64(value)
	if err != nil {
		return nil,err
	}

	t:=types.NewTotalStake(int64(height), totalStake)

	return &t,nil
}

func getTotalStakeByType(height int64,flowClient client.Proxy) ([]types.TotalStakeByType,error) {
	log.Trace().Str("module", "staking").Int64("height", int64(height)).
		Msg("updating getTotalStakeByType")

		// https://github.com/onflow/flow-core-contracts/blob/301206b27090fa9116c1aba35cd8bade8a26857c/contracts/FlowIDTableStaking.cdc#L101
		//
		// 1 = collection
		// 2 = consensus
		// 3 = execution
		// 4 = verification
		// 5 = access
	script := fmt.Sprintf(`
	import FlowIDTableStaking from %s
	pub fun main(role: UInt8): UFix64 {
	  let staked = FlowIDTableStaking.getTotalTokensStakedByNodeType()
 
	  return staked[role]!
	}`, flowClient.Contract().StakingTable)

	totalStakeArr := make([]types.TotalStakeByType, 5)
	for role := 1; role <= 5; role++ {
		arg := []cadence.Value{cadence.NewUInt8(uint8(role))}
		value, err := flowClient.Client().ExecuteScriptAtLatestBlock(flowClient.Ctx(), []byte(script), arg)
		if err != nil {
			return nil,err
		}

		totalStake, err := utils.CadenceConvertUint64(value)
		if err != nil {
			return nil,err
		}

		totalStakeArr[role-1] = types.NewTotalStakeByType(int64(height), int8(role), totalStake)
	}

	return totalStakeArr,nil
}

func getStakeRequirements(height int64,flowClient client.Proxy) ([]types.StakeRequirements,error) {
	log.Trace().Str("module", "staking").Int64("height", int64(height)).
		Msg("updating get stake requirement")

		// https://github.com/onflow/flow-core-contracts/blob/301206b27090fa9116c1aba35cd8bade8a26857c/contracts/FlowIDTableStaking.cdc#L101
		//
		// 1 = collection
		// 2 = consensus
		// 3 = execution
		// 4 = verification
		// 5 = access
	script := fmt.Sprintf(`
	import FlowIDTableStaking from %s
	pub fun main(role: UInt8): UFix64 {
	  let req = FlowIDTableStaking.getMinimumStakeRequirements()
  
	  return req[role]!
  }`, flowClient.Contract().StakingTable)

	stakeRequirements := make([]types.StakeRequirements, 5)
	for role := 1; role <= 5; role++ {
		arg := []cadence.Value{cadence.NewUInt8(uint8(role))}
		value, err := flowClient.Client().ExecuteScriptAtLatestBlock(flowClient.Ctx(), []byte(script), arg)
		if err != nil {
			return nil,err
		}

		totalStake, err := utils.CadenceConvertUint64(value)
		if err != nil {
			return nil,err
		}
		stakeRequirements[role-1] = types.NewStakeRequirements(int64(height), uint8(role), totalStake)
	}

	return stakeRequirements,nil
}

func getProposedTable(height int64,flowClient client.Proxy) (*types.ProposedTable,error) {
	log.Trace().Str("module", "staking").Int64("height", int64(height)).
		Msg("updating get ProposedTable")

	script := fmt.Sprintf(`
	import FlowIDTableStaking from %s
          pub fun main(): [String] {
            return FlowIDTableStaking.getProposedNodeIDs()
        }`, flowClient.Contract().StakingTable)

	value, err := flowClient.Client().ExecuteScriptAtLatestBlock(flowClient.Ctx(), []byte(script), nil)
	if err != nil {
		return nil,err
	}

	table, err := utils.CadenceConvertStringArray(value)
	if err != nil {
		return nil,err
	}

	t:=types.NewProposedTable(int64(height), table)

	return &t,nil
}

func getCutPercentage(height int64,flowClient client.Proxy) (*types.CutPercentage,error) {
	log.Trace().Str("module", "staking").Int64("height", int64(height)).
		Msg("updating cut percentage")
	script := fmt.Sprintf(`
	import FlowIDTableStaking from %s
	pub fun main(): UFix64 {
	  return FlowIDTableStaking.getRewardCutPercentage()
  }`, flowClient.Contract().StakingTable)

	value, err := flowClient.Client().ExecuteScriptAtLatestBlock(flowClient.Ctx(), []byte(script), nil)
	if err != nil {
		return nil,err
	}

	table, err := utils.CadenceConvertUint64(value)
	if err != nil {
		return nil,err
	}
	t:=types.NewCutPercentage(table, int64(height))

	return &t,nil
}



func GetTable(height int64,flowClient client.Proxy) (*types.StakingTable, error) {
	log.Trace().Str("module", "staking").Int64("height", int64(height)).
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

	if len(table) == 0 {
		return nil, nil
	}

	t:=types.NewStakingTable(height,table)

	return &t,nil

}

func GetNodeInfosFromTable(height int64,flowClient client.Proxy) ([]types.StakerNodeInfo, error) {

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

	if len(stakingKeys) == 0 {
		return nil, nil
	}

	return stakingKeys, nil
}

func getAddressesFromAccounts(accounts []types.Account) []string {
	addresses := make([]string, len(accounts))
	for i, account := range accounts {
		addresses[i] = account.Address
	}
	return addresses
}
