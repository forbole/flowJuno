package postgresql

import (
	"fmt"

	"github.com/forbole/flowJuno/types"
)


func (db Db) SaveStakeRequirements(stakeRequirements []types.StakeRequirements) error {
    stmt:= `INSERT INTO stake_requirements(height,role,requirements,timestamp) VALUES `

    var params []interface{}

	  for i, rows := range stakeRequirements{
      ai := i * 4
      stmt += fmt.Sprintf("($%d,$%d,$%d,$%d)", ai+1,ai+2,ai+3,ai+4)
      
      params = append(params,rows.Height,rows.Role,rows.Requirement,rows.Timestamp)

    }
	  stmt = stmt[:len(stmt)-1]

    stmt += ` ON CONFLICT DO NOTHING` 

    _, err := db.Sqlx.Exec(stmt, params...)
    if err != nil {
      return err
    }

    return nil 
}
    