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
	stmt := `INSERT INTO staking_table(height,staking_table) VALUES ($1,$2) ON CONFLICT DO NOTHING`

	nodeInfoList, err := json.Marshal(stakingTable.StakingTable)
	if err != nil {
		return err
	}
	_, err = db.Sql.Exec(stmt, stakingTable.Height,
		string(nodeInfoList))
	return err
}

func (db *Db) SaveProposedTable(proposedTable types.ProposedTable) error {
	stmt := `INSERT INTO proposed_table(height,proposed_table) VALUES ($1,$2) ON CONFLICT DO NOTHING`

	table, err := json.Marshal(proposedTable.ProposedTable)
	if err != nil {
		return err
	}
	_, err = db.Sql.Exec(stmt, proposedTable.Height,
		table)
	return err
}

func (db *Db) SaveCurrentTable(currentTable types.CurrentTable) error {
	stmt := `INSERT INTO current_table(height,current_table) VALUES ($1,$2) ON CONFLICT DO NOTHING`

	table, err := json.Marshal(currentTable.Table)
	if err != nil {
		return err
	}

	_, err = db.Sql.Exec(stmt, currentTable.Height,
		table)
	return err
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

func (db *Db) SaveNodeStakingKey(nodeStakingKey []types.NodeStakingKey) error {
	stmt := `INSERT INTO node_staking_key(node_id,node_staking_key,height) VALUES `

	var params []interface{}

	for i, rows := range nodeStakingKey {
		ai := i * 3
		stmt += fmt.Sprintf("($%d,$%d,$%d),", ai+1, ai+2, ai+3)

		params = append(params, rows.NodeId, rows.NodeStakingKey, rows.Height)

	}
	stmt = stmt[:len(stmt)-1]
	stmt += ` ON CONFLICT DO NOTHING`

	_, err := db.Sqlx.Exec(stmt, params...)
	if err != nil {
		return err
	}

	return nil
}

func (db *Db) SaveNodeStakedTokens(nodeStakedTokens []types.NodeStakedTokens) error {
	stmt := `INSERT INTO node_staked_tokens(node_id,node_staked_tokens,height) VALUES `

	var params []interface{}

	for i, rows := range nodeStakedTokens {
		ai := i * 3
		stmt += fmt.Sprintf("($%d,$%d,$%d),", ai+1, ai+2, ai+3)

		params = append(params, rows.NodeId, rows.NodeStakedTokens, rows.Height)
	}
	stmt = stmt[:len(stmt)-1]
	stmt += ` ON CONFLICT DO NOTHING`

	_, err := db.Sqlx.Exec(stmt, params...)
	if err != nil {
		return err
	}

	return nil
}

func (db *Db) SaveNodeRole(nodeRole []types.NodeRole) error {
	stmt := `INSERT INTO node_role(node_id,role,height) VALUES `

	var params []interface{}

	for i, rows := range nodeRole {
		ai := i * 3
		stmt += fmt.Sprintf("($%d,$%d,$%d),", ai+1, ai+2, ai+3)

		params = append(params, rows.NodeId, rows.Role, rows.Height)

	}
	stmt = stmt[:len(stmt)-1]
	stmt += ` ON CONFLICT DO NOTHING`

	_, err := db.Sqlx.Exec(stmt, params...)
	if err != nil {
		return err
	}

	return nil
}

func (db *Db) SaveNodeRewardedTokens(nodeRewardedTokens []types.NodeRewardedTokens) error {
	stmt := `INSERT INTO node_rewarded_tokens(node_id,node_rewarded_tokens,height) VALUES `

	var params []interface{}

	for i, rows := range nodeRewardedTokens {
		ai := i * 3
		stmt += fmt.Sprintf("($%d,$%d,$%d),", ai+1, ai+2, ai+3)

		params = append(params, rows.NodeId, rows.NodeRewardedTokens, rows.Height)

	}
	stmt = stmt[:len(stmt)-1]
	stmt += ` ON CONFLICT DO NOTHING`

	_, err := db.Sqlx.Exec(stmt, params...)
	if err != nil {
		return err
	}

	return nil
}

func (db *Db) SaveNodeNetworkingKey(nodeNetworkingKey []types.NodeNetworkingKey) error {
	stmt := `INSERT INTO node_networking_key(node_id,networking_key,height) VALUES `

	var params []interface{}

	for i, rows := range nodeNetworkingKey {
		ai := i * 3
		stmt += fmt.Sprintf("($%d,$%d,$%d),", ai+1, ai+2, ai+3)

		params = append(params, rows.NodeId, rows.NetworkingKey, rows.Height)

	}
	stmt = stmt[:len(stmt)-1]
	stmt += ` ON CONFLICT DO NOTHING`

	_, err := db.Sqlx.Exec(stmt, params...)
	if err != nil {
		return err
	}

	return nil
}

func (db *Db) SaveNodeNetworkingAddress(nodeNetworkingAddress []types.NodeNetworkingAddress) error {
	stmt := `INSERT INTO node_networking_address(node_id,networking_address,height) VALUES `

	var params []interface{}

	for i, rows := range nodeNetworkingAddress {
		ai := i * 3
		stmt += fmt.Sprintf("($%d,$%d,$%d),", ai+1, ai+2, ai+3)

		params = append(params, rows.NodeId, rows.NetworkingAddress, rows.Height)

	}
	stmt = stmt[:len(stmt)-1]
	stmt += ` ON CONFLICT DO NOTHING`

	_, err := db.Sqlx.Exec(stmt, params...)
	if err != nil {
		return err
	}

	return nil
}

func (db *Db) SaveNodeInitialWeight(nodeInitialWeight []types.NodeInitialWeight) error {
	stmt := `INSERT INTO node_initial_weight(node_id,initial_weight,height) VALUES `

	var params []interface{}

	for i, rows := range nodeInitialWeight {
		ai := i * 3
		stmt += fmt.Sprintf("($%d,$%d,$%d),", ai+1, ai+2, ai+3)

		params = append(params, rows.NodeId, rows.InitialWeight, rows.Height)

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

func (db *Db) SaveNodeInfosFromTable(nodeInfosFromTable []types.StakerNodeInfo,height uint64) error {
    stmt:= `INSERT INTO node_infos_from_table(id,role,networking_address,networking_key,staking_key,tokens_staked,tokens_committed,tokens_unstaking,tokens_unstaked,tokens_rewarded,delegators,delegator_i_d_counter,tokens_requested_to_unstake,initial_weight,height) VALUES `

    var params []interface{}

	  for i, rows := range nodeInfosFromTable{
      ai := i * 15
      stmt += fmt.Sprintf("($%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d),", ai+1,ai+2,ai+3,ai+4,ai+5,ai+6,ai+7,ai+8,ai+9,ai+10,ai+11,ai+12,ai+13,ai+14,ai+15)
      
      params = append(params,rows.Id,rows.Role,rows.NetworkingAddress,rows.NetworkingKey,rows.StakingKey,rows.TokensStaked,rows.TokensCommitted,rows.TokensUnstaking,rows.TokensUnstaked,rows.TokensRewarded,pq.Array(rows.Delegators),rows.DelegatorIDCounter,rows.TokensRequestedToUnstake,rows.InitialWeight,height)

    }
	  stmt = stmt[:len(stmt)-1]
    stmt += ` ON CONFLICT DO NOTHING` 

    _, err := db.Sqlx.Exec(stmt, params...)
    if err != nil {
      return err
    }

    return nil 
    }

func (db *Db) SaveNodeCommittedTokens(nodeCommittedTokens []types.NodeCommittedTokens) error {
	stmt := `INSERT INTO node_committed_tokens(node_id,committed_tokens,height) VALUES `

	var params []interface{}

	for i, rows := range nodeCommittedTokens {
		ai := i * 3
		stmt += fmt.Sprintf("($%d,$%d,$%d),", ai+1, ai+2, ai+3)

		params = append(params, rows.NodeId, rows.CommittedTokens, rows.Height)

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

func (db *Db) SaveDelegatorCommitted(delegatorCommitted types.DelegatorCommitted) error {
	stmt := `INSERT INTO delegator_committed(committed,height,node_id,delegator_id) VALUES ($1,$2,$3,$4) ON CONFLICT DO NOTHING`
	_, err := db.Sql.Exec(stmt, delegatorCommitted.Committed,
		delegatorCommitted.Height,
		delegatorCommitted.NodeId,
		delegatorCommitted.DelegatorID)
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

func (db *Db) SaveDelegatorRequest(delegatorRequest types.DelegatorRequest) error {
	stmt := `INSERT INTO delegator_request(request_to_unstake,height,node_id,delegator_id) VALUES ($1,$2,$3,$4) ON CONFLICT DO NOTHING`
	_, err := db.Sql.Exec(stmt, delegatorRequest.RequestToUnstake,
		delegatorRequest.Height,
		delegatorRequest.NodeId,
		delegatorRequest.DelegatorId)
	return err
}

func (db *Db) SaveDelegatorRewarded(delegatorRewarded types.DelegatorRewarded) error {
	stmt := `INSERT INTO delegator_rewarded(rewarded,height,node_id,delegator_id) VALUES ($1,$2,$3,$4) ON CONFLICT DO NOTHING`
	_, err := db.Sql.Exec(stmt, delegatorRewarded.Rewarded,
		delegatorRewarded.Height,
		delegatorRewarded.NodeId,
		delegatorRewarded.DelegatorId)
	return err
}

func (db *Db) SaveDelegatorStaked(delegatorStaked types.DelegatorStaked) error {
	stmt := `INSERT INTO delegator_staked(staked,height,node_id,delegator_id) VALUES ($1,$2,$3,$4) ON CONFLICT DO NOTHING`
	_, err := db.Sql.Exec(stmt, delegatorStaked.Staked,
		delegatorStaked.Height,
		delegatorStaked.NodeId,
		delegatorStaked.DelegatorId)
	return err
}

func (db *Db) SaveDelegatorUnstaked(delegatorUnstaked types.DelegatorUnstaked) error {
	stmt := `INSERT INTO delegator_unstaked(unstaked,height,node_id,delegator_id) VALUES ($1,$2,$3,$4) ON CONFLICT DO NOTHING`
	_, err := db.Sql.Exec(stmt, delegatorUnstaked.Unstaked,
		delegatorUnstaked.Height,
		delegatorUnstaked.NodeId,
		delegatorUnstaked.DelegatorId)
	return err
}

func (db *Db) SaveDelegatorUnstaking(delegatorUnstaking types.DelegatorUnstaking) error {
	stmt := `INSERT INTO delegator_unstaking(unstaking,height,node_id,delegator_id) VALUES ($1,$2,$3,$4) ON CONFLICT DO NOTHING`
	_, err := db.Sql.Exec(stmt, delegatorUnstaking.Unstaking,
		delegatorUnstaking.Height,
		delegatorUnstaking.NodeId,
		delegatorUnstaking.DelegatorId)
	return err
}

func (db *Db) SaveDelegatorUnstakingRequest(delegatorUnstakingRequest types.DelegatorUnstakingRequest) error {
	stmt := `INSERT INTO delegator_unstaking_request(unstaking,height,node_id,delegator_id) VALUES ($1,$2,$3,$4) ON CONFLICT DO NOTHING`
	_, err := db.Sql.Exec(stmt, delegatorUnstakingRequest.Unstaking,
		delegatorUnstakingRequest.Height,
		delegatorUnstakingRequest.NodeId,
		delegatorUnstakingRequest.DelegatorId)
	return err
}

/* func (db *Db) GetDelegatorInfoFromNodeId(nodeId string, height int64) ([]types.NodeInfoFromNodeID, error) {
	var rows []dbtypes.NodeInfoFromNodeIDRow
	stmt := `SELECT * FROM node_info_from_node_id WHERE node_id=$1 AND height=$2;`
	err := db.Sqlx.Select(&rows, stmt, nodeId, height)
	if err != nil {
		return nil, err
	}
	returnRows := make([]types.NodeInfoFromNodeID, len(rows))
	for i, row := range rows {
		var nodeInfo types.StakerNodeInfo
		err = json.Unmarshal([]byte(row.NodeInfo), &nodeInfo)
		if err != nil {
			return nil, err
		}
		returnRows[i] = types.NewNodeInfoFromNodeID(nodeId, nodeInfo, height)
	}
	return returnRows, nil
} */
