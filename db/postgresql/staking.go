package postgresql

import (
	"encoding/json"
	"fmt"

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
	stmt:= `INSERT INTO proposed_table(height,proposed_table) VALUES ($1,$2) ON CONFLICT DO NOTHING`
	
	table, err := json.Marshal(proposedTable.ProposedTable)
	if err != nil {
		return err
	}
	_, err = db.Sql.Exec(stmt,proposedTable.Height,
		table)
		return err 
	 }
	