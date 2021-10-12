package stakingutils

import (
	"fmt"

	"github.com/rs/zerolog/log"

	"github.com/forbole/flowJuno/types"

	"github.com/onflow/flow-go-sdk"

	"github.com/forbole/flowJuno/client"
	"github.com/forbole/flowJuno/modules/utils"

	database "github.com/forbole/flowJuno/db/postgresql"
	db "github.com/forbole/flowJuno/db/postgresql"
	"github.com/onflow/cadence"
)
func GetDataWithNoArgs(block *flow.Block, db *db.Db, height int64, flowClient client.Proxy) error {
	err := getWeeklyPayout(block, db, flowClient)
	if err != nil {
		return err
	}

	err = getTotalStake(block, db, flowClient)
	if err != nil {
		return err
	}

	err = getTotalStakeByType(block, db, flowClient)
	if err != nil {
		return err
	}

	_,err = getCurrentTable(block, db, flowClient)
	if err != nil {
		return err
	}

	err = getStakeRequirements(block, db, flowClient)
	if err != nil {
		return err
	}

	err = getProposedTable(block, db, flowClient)
	if err != nil {
		return err
	}

	err=getCutPercentage(block, db, flowClient)
	if err!=nil{
		return err
	}

	return nil
}

func getCurrentTable(block *flow.Block, db *database.Db, flowClient client.Proxy) ([]string, error) {
	log.Trace().Str("module", "staking").Int64("height", int64(block.Height)).
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

	err = db.SaveCurrentTable(types.NewCurrentTable(int64(block.Height), ids))
	if err != nil {
		return nil, err
	}

	return ids, nil
}

// getStakedNodeId get staked node id from cadence script
func getWeeklyPayout(block *flow.Block, db *database.Db, flowClient client.Proxy) error {
	log.Trace().Str("module", "staking").Int64("height", int64(block.Height)).
		Msg("updating get weekly epoch payout")

	script := fmt.Sprintf(`
	import FlowIDTableStaking from %s
	pub fun main(): UFix64 {
	  return FlowIDTableStaking.getEpochTokenPayout()
  }`, flowClient.Contract().StakingTable)

	value, err := flowClient.Client().ExecuteScriptAtLatestBlock(flowClient.Ctx(), []byte(script), nil)
	if err != nil {
		return err
	}

	payout, err := utils.CadenceConvertUint64(value)
	if err != nil {
		return err
	}

	return db.SaveWeeklyPayout(types.NewWeeklyPayout(int64(block.Height), payout))
}

func getTotalStake(block *flow.Block, db *database.Db, flowClient client.Proxy) error {
	log.Trace().Str("module", "staking").Int64("height", int64(block.Height)).
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
		return err
	}

	totalStake, err := utils.CadenceConvertUint64(value)
	if err != nil {
		return err
	}

	return db.SaveTotalStake(types.NewTotalStake(int64(block.Height), totalStake))
}

func getTotalStakeByType(block *flow.Block, db *database.Db, flowClient client.Proxy) error {
	log.Trace().Str("module", "staking").Int64("height", int64(block.Height)).
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
			return err
		}

		totalStake, err := utils.CadenceConvertUint64(value)
		if err != nil {
			return err
		}

		totalStakeArr[role-1] = types.NewTotalStakeByType(int64(block.Height), int8(role), totalStake)
	}

	return db.SaveTotalStakeByType(totalStakeArr)
}



func getStakeRequirements(block *flow.Block, db *database.Db, flowClient client.Proxy) error {
	log.Trace().Str("module", "staking").Int64("height", int64(block.Height)).
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
			return err
		}

		totalStake, err := utils.CadenceConvertUint64(value)
		if err != nil {
			return err
		}
		stakeRequirements[role-1] = types.NewStakeRequirements(int64(block.Height), uint8(role), totalStake)
	}

	return db.SaveStakeRequirements(stakeRequirements)
}

func getProposedTable(block *flow.Block, db *database.Db, flowClient client.Proxy) error {
	log.Trace().Str("module", "staking").Int64("height", int64(block.Height)).
		Msg("updating get ProposedTable")

	script := fmt.Sprintf(`
	import FlowIDTableStaking from %s
          pub fun main(): [String] {
            return FlowIDTableStaking.getProposedNodeIDs()
        }`, flowClient.Contract().StakingTable)

	value, err := flowClient.Client().ExecuteScriptAtLatestBlock(flowClient.Ctx(), []byte(script), nil)
	if err != nil {
		return err
	}

	table, err := utils.CadenceConvertStringArray(value)
	if err != nil {
		return err
	}

	return db.SaveProposedTable(types.NewProposedTable(int64(block.Height), table))
}

func getCutPercentage(block *flow.Block, db *database.Db, flowClient client.Proxy) error {
	log.Trace().Str("module", "staking").Int64("height", int64(block.Height)).
		Msg("updating cut percentage")
	script := fmt.Sprintf(`
	import FlowIDTableStaking from %s
	pub fun main(): UFix64 {
	  return FlowIDTableStaking.getRewardCutPercentage()
  }`, flowClient.Contract().StakingTable)

	value, err := flowClient.Client().ExecuteScriptAtLatestBlock(flowClient.Ctx(), []byte(script), nil)
	if err != nil {
		return err
	}

	table, err := utils.CadenceConvertUint64(value)
	if err != nil {
		return err
	}

	return db.SaveCutPercentage(types.NewCutPercentage(table, int64(block.Height)))
}
