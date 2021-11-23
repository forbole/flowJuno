package postgresql

import (
	"encoding/json"
	"fmt"

	"github.com/lib/pq"

	"github.com/forbole/flowJuno/types"
)

// SaveStakeRequirements save the stake requirement from cadence call
func (db *Db) SaveStakeRequirements(stakeRequirements []types.StakeRequirements) error {
	stmt := `INSERT INTO stake_requirements(height,role,requirements) VALUES `

	var params []interface{}

	for i, rows := range stakeRequirements {
		ai := i * 3
		stmt += fmt.Sprintf("($%d,$%d,$%d),", ai+1, ai+2, ai+3)

		params = append(params, rows.Height, rows.Role, rows.Requirement)

	}
	stmt = stmt[:len(stmt)-1]

	stmt += ` ON CONFLICT DO NOTHING`

	_, err := db.Sqlx.Exec(stmt, params...)
	if err != nil {
		return err
	}

	return nil
}

func (db *Db) SaveTotalStakeByType(totalStake []types.TotalStakeByType) error {
	stmt := `INSERT INTO total_stake_by_type(height,role,total_stake) VALUES `

	var params []interface{}

	for i, rows := range totalStake {
		ai := i * 3
		stmt += fmt.Sprintf("($%d,$%d,$%d),", ai+1, ai+2, ai+3)

		params = append(params, rows.Height, rows.Role, rows.TotalStake)

	}
	stmt = stmt[:len(stmt)-1]

	stmt += ` ON CONFLICT DO NOTHING`

	_, err := db.Sqlx.Exec(stmt, params...)
	if err != nil {
		return err
	}

	return nil
}

func (db *Db) SaveWeeklyPayout(weeklyPayout types.WeeklyPayout) error {
	stmt := `INSERT INTO weekly_payout(height,payout) VALUES ($1,$2) ON CONFLICT DO NOTHING`
	_, err := db.Sql.Exec(stmt, weeklyPayout.Height,
		weeklyPayout.Payout)
	return err
}

func (db *Db) SaveTotalStake(totalStake types.TotalStake) error {
	stmt := `INSERT INTO total_stake(height,total_stake) VALUES ($1,$2) ON CONFLICT DO NOTHING`
	_, err := db.Sql.Exec(stmt, totalStake.Height,
		totalStake.TotalStake)
	return err
}

func (db *Db) SaveStakingTable(stakingTable types.StakingTable) error {
	stmt := `INSERT INTO staking_table(node_id) VALUES `

	var params []interface{}

	for i, rows := range stakingTable.StakingTable {
		ai := i * 1
		stmt += fmt.Sprintf("($%d),", ai+1)

		params = append(params, rows)

	}
	stmt = stmt[:len(stmt)-1]
	stmt += ` ON CONFLICT DO NOTHING`

	_, err := db.Sqlx.Exec(stmt, params...)
	if err != nil {
		return err
	}

	return nil
}

func (db *Db) SaveProposedTable(proposedTable types.ProposedTable) error {
	stmt := `INSERT INTO proposed_table(height,proposed_table) VALUES`

	var params []interface{}

	for i, rows := range proposedTable.ProposedTable {
		ai := i * 2
		stmt += fmt.Sprintf("($%d,$%d),", ai+1, ai+2)
		fmt.Println(rows)

		params = append(params, proposedTable.Height, rows)
	}
	stmt = stmt[:len(stmt)-1]
	stmt += ` ON CONFLICT DO NOTHING`

	_, err := db.Sqlx.Exec(stmt, params...)
	if err != nil {
		return fmt.Errorf("fail to insert into proposed table: %s",err)
	}

	return err
}

func (db *Db) SaveCurrentTable(currentTable types.CurrentTable) error {
	stmt := `INSERT INTO current_table(height,current_table) VALUES `

	var params []interface{}

	for i, rows := range currentTable.Table {
		ai := i * 2
		stmt += fmt.Sprintf("($%d,$%d),", ai+1, ai+2)
		fmt.Println(rows)

		params = append(params, currentTable.Height, rows)
	}
	stmt = stmt[:len(stmt)-1]
	stmt += ` ON CONFLICT DO NOTHING`

	_, err := db.Sqlx.Exec(stmt, params...)
	if err != nil {
		return fmt.Errorf("fail to insert into current table: %s",err)
	}

	return nil
}

func (db *Db) SaveNodeTotalCommitment(nodeTotalCommitment []types.NodeTotalCommitment) error {
	stmt := `INSERT INTO node_total_commitment(node_id,total_commitment,height) VALUES `

	var params []interface{}

	for i, rows := range nodeTotalCommitment {
		ai := i * 3
		stmt += fmt.Sprintf("($%d,$%d,$%d),", ai+1, ai+2, ai+3)

		params = append(params, rows.NodeId, rows.TotalCommitment, rows.Height)

	}
	stmt = stmt[:len(stmt)-1]
	stmt += ` ON CONFLICT DO NOTHING`

	_, err := db.Sqlx.Exec(stmt, params...)
	if err != nil {
		return err
	}

	return nil
}

func (db *Db) SaveNodeTotalCommitmentWithoutDelegators(nodeTotalCommitmentWithoutDelegators []types.NodeTotalCommitmentWithoutDelegators) error {
	stmt := `INSERT INTO node_total_commitment_without_delegators(node_id,total_commitment_without_delegators,height) VALUES `

	var params []interface{}

	for i, rows := range nodeTotalCommitmentWithoutDelegators {
		ai := i * 3
		stmt += fmt.Sprintf("($%d,$%d,$%d),", ai+1, ai+2, ai+3)

		params = append(params, rows.NodeId, rows.TotalCommitmentWithoutDelegators, rows.Height)

	}
	stmt = stmt[:len(stmt)-1]
	stmt += ` ON CONFLICT DO NOTHING`

	_, err := db.Sqlx.Exec(stmt, params...)
	if err != nil {
		return err
	}

	return nil
}

func (db *Db) SaveNodeInfoFromAddresses(NodeInfoFromAddress []types.NodeInfoFromAddress) error {
	stmt := `INSERT INTO node_info_from_address(address,node_info,height) VALUES `

	var params []interface{}

	for i, rows := range NodeInfoFromAddress {
		ai := i * 3
		stmt += fmt.Sprintf("($%d,$%d,$%d),", ai+1, ai+2, ai+3)

		nodeInfo, err := json.Marshal(rows.NodeInfo)
		if err != nil {
			return err
		}
		params = append(params, rows.Address, nodeInfo, rows.Height)

	}
	stmt = stmt[:len(stmt)-1]
	stmt += ` ON CONFLICT DO NOTHING`

	_, err := db.Sqlx.Exec(stmt, params...)
	if err != nil {
		return err
	}

	return nil
}

func (db *Db) SaveNodeInfosFromTable(nodeInfosFromTable []types.StakerNodeInfo, height uint64) error {
	stmt := `INSERT INTO node_infos_from_table(id,role,networking_address,networking_key,staking_key,tokens_staked,tokens_committed,tokens_unstaking,tokens_unstaked,tokens_rewarded,delegators,delegator_i_d_counter,tokens_requested_to_unstake,initial_weight,height) VALUES `

	var params []interface{}

	for i, rows := range nodeInfosFromTable {
		ai := i * 15
		stmt += fmt.Sprintf("($%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d),", ai+1, ai+2, ai+3, ai+4, ai+5, ai+6, ai+7, ai+8, ai+9, ai+10, ai+11, ai+12, ai+13, ai+14, ai+15)

		params = append(params, rows.Id, rows.Role, rows.NetworkingAddress, rows.NetworkingKey, rows.StakingKey, rows.TokensStaked, rows.TokensCommitted, rows.TokensUnstaking, rows.TokensUnstaked, rows.TokensRewarded, pq.Array(rows.Delegators), rows.DelegatorIDCounter, rows.TokensRequestedToUnstake, rows.InitialWeight, height)

	}
	stmt = stmt[:len(stmt)-1]
	stmt += ` ON CONFLICT DO NOTHING`

	_, err := db.Sqlx.Exec(stmt, params...)
	if err != nil {
		return err
	}

	return nil
}

func (db *Db) SaveCutPercentage(cutPercentage types.CutPercentage) error {
	stmt := `INSERT INTO cut_percentage(cut_percentage,height) VALUES ($1,$2) ON CONFLICT DO NOTHING`

	_, err := db.Sql.Exec(stmt, cutPercentage.CutPercentage, cutPercentage.Height)
	return err
}

func (db *Db) SaveDelegatorInfo(delegatorInfo []types.DelegatorNodeInfo, height uint64) error {
	stmt := `INSERT INTO delegator_info(id,node_id,tokens_committed,tokens_staked,tokens_unstaking,tokens_rewarded,tokens_unstaked,tokens_requested_to_unstake,height) VALUES `

	var params []interface{}

	for i, rows := range delegatorInfo {
		ai := i * 9
		stmt += fmt.Sprintf("($%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d),", ai+1, ai+2, ai+3, ai+4, ai+5, ai+6, ai+7, ai+8, ai+9)

		params = append(params, rows.Id, rows.NodeID, rows.TokensCommitted, rows.TokensStaked, rows.TokensUnstaking, rows.TokensRewarded, rows.TokensUnstaked, rows.TokensRequestedToUnstake, height)

	}
	stmt = stmt[:len(stmt)-1]
	stmt += ` ON CONFLICT DO NOTHING`
	fmt.Println(stmt)
	_, err := db.Sqlx.Exec(stmt, params...)
	if err != nil {
		return err
	}

	return nil
}

func (db *Db) SaveDelegatorInfoFromAddress(delegatorInfoFromAddress []types.DelegatorInfoFromAddress) error {
	stmt := `INSERT INTO delegator_info_from_address(delegator_info,height,address) VALUES `

	var params []interface{}

	for i, rows := range delegatorInfoFromAddress {
		ai := i * 3
		stmt += fmt.Sprintf("($%d,$%d,$%d),", ai+1, ai+2, ai+3)

		info, err := json.Marshal(rows.DelegatorInfo)
		if err != nil {
			return err
		}

		params = append(params, info, rows.Height, rows.Address)

	}
	stmt = stmt[:len(stmt)-1]
	stmt += ` ON CONFLICT DO NOTHING`

	_, err := db.Sqlx.Exec(stmt, params...)
	if err != nil {
		return err
	}

	return nil
}

func (db *Db) SaveNodeUnstakingTokens(nodeUnstakingTokens []types.NodeUnstakingTokens) error {
	stmt := `INSERT INTO node_unstaking_tokens(node_id,token_unstaking,height) VALUES `

	var params []interface{}

	for i, rows := range nodeUnstakingTokens {
		ai := i * 3
		stmt += fmt.Sprintf("($%d,$%d,$%d),", ai+1, ai+2, ai+3)

		params = append(params, rows.NodeId, rows.TokenUnstaking, rows.Height)

	}
	stmt = stmt[:len(stmt)-1]
	stmt += ` ON CONFLICT DO NOTHING`

	_, err := db.Sqlx.Exec(stmt, params...)
	if err != nil {
		return err
	}

	return nil
}
