package postgresql

import (
	"fmt"

	"github.com/forbole/flowJuno/types"
)

// SaveStakeRequirements save the stake requirement from cadence call
func (db *Db) SaveStakeRequirements(stakeRequirements []types.StakeRequirements) error {
	stmt := `INSERT INTO stake_requirements(height,role,requirements,timestamp) VALUES `

	var params []interface{}

	for i, rows := range stakeRequirements {
		ai := i * 4
		stmt += fmt.Sprintf("($%d,$%d,$%d,$%d)", ai+1, ai+2, ai+3, ai+4)

		params = append(params, rows.Height, rows.Role, rows.Requirement, rows.Timestamp)

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
	stmt := `INSERT INTO total_stake_by_type(height,role,total_stake,timestamp) VALUES `

	var params []interface{}

	for i, rows := range totalStake {
		ai := i * 4
		stmt += fmt.Sprintf("($%d,$%d,$%d,$%d)", ai+1, ai+2, ai+3, ai+4)

		params = append(params, rows.Height, rows.Role, rows.TotalStake, rows.Timestamp)

	}
	stmt = stmt[:len(stmt)-1]

	stmt += ` ON CONFLICT DO NOTHING`

	_, err := db.Sqlx.Exec(stmt, params...)
	if err != nil {
		return err
	}

	return nil
}
