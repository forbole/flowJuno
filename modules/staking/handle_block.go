package staking

import (
	"fmt"

	"github.com/rs/zerolog/log"

	"github.com/forbole/flowJuno/modules/messages"
	"github.com/forbole/flowJuno/types"
	"github.com/onflow/flow-go-sdk"

	"github.com/forbole/flowJuno/client"
	database "github.com/forbole/flowJuno/db/postgresql"
	db "github.com/forbole/flowJuno/db/postgresql"
	"github.com/onflow/cadence"
)

func HandleBlock(block *flow.Block, _ messages.MessageAddressesParser, db *db.Db, height int64, flowClient client.Proxy) error {
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

	err = getTable(block, db, flowClient)
	if err != nil {
		return err
	}

	err = getStakeRequirements(block, db, flowClient)
	if err != nil {
		return err
	}

	return nil
}

// getStakedNodeId get staked node id from cadence script
func getWeeklyPayout(block *flow.Block, db *database.Db, flowClient client.Proxy) error {
	log.Trace().Str("module", "staking").Int64("height", int64(block.Height)).
		Msg("updating get Staked Node id per block")

	script := fmt.Sprintf(`
	import FlowIDTableStaking from %s
	pub fun main(): UFix64 {
	  return FlowIDTableStaking.getEpochTokenPayout()
  }`, flowClient.Contract().StakingTable)

	value, err := flowClient.Client().ExecuteScriptAtLatestBlock(flowClient.Ctx(), []byte(script), nil)
	if err != nil {
		return err
	}

	payout, ok := value.ToGoValue().(uint64)
	if !ok {
		return fmt.Errorf("Payout is not a uint64 value")
	}

	return db.SaveWeeklyPayout(types.NewWeeklyPayout(int64(block.Height), payout))
}

func getTotalStake(block *flow.Block, db *database.Db, flowClient client.Proxy) error {
	log.Trace().Str("module", "staking").Int64("height", int64(block.Height)).
		Msg("updating get Staked Node id per block")

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

	totalStake, ok := value.ToGoValue().(uint64)
	if !ok {
		return fmt.Errorf("totalStake is not a uint64 value")
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

		totalStake, ok := value.ToGoValue().(uint64)
		if !ok {
			return fmt.Errorf("totalStake is not a uint64 value")
		}

		totalStakeArr[role-1] = types.NewTotalStakeByType(int64(block.Height), int8(role), totalStake)
	}

	return db.SaveTotalStakeByType(totalStakeArr)
}

func getTable(block *flow.Block, db *database.Db, flowClient client.Proxy) error {
	log.Trace().Str("module", "staking").Int64("height", int64(block.Height)).
		Msg("updating get Staked Node id per block")

	script := fmt.Sprintf(`
	import FlowIDTableStaking from %s
	pub fun main(): [String] {
	  return FlowIDTableStaking.getNodeIDs()
  }`, flowClient.Contract().StakingTable)

	value, err := flowClient.Client().ExecuteScriptAtLatestBlock(flowClient.Ctx(), []byte(script), nil)
	if err != nil {
		return err
	}

	valueArray, ok := value.(cadence.Array)
	if !ok {
		return fmt.Errorf("table is not an array")
	}

	table := make([]string, len(valueArray.Values))
	for i, val := range valueArray.Values {
		table[i] = val.String()
	}

	return db.SaveStakingTable(types.NewStakingTable(int64(block.Height),table))

}

func getStakeRequirements(block *flow.Block, db *database.Db, flowClient client.Proxy) error {
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

		totalStake, ok := value.ToGoValue().(uint64)
		if !ok {
			return fmt.Errorf("totalStake is not a uint64 value")
		}
		stakeRequirements[role-1] = types.NewStakeRequirements(int64(block.Height), uint8(role), totalStake)
	}

	return db.SaveStakeRequirements(stakeRequirements)
}
