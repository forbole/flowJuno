package postgresql

import (
	"fmt"

	dbutils "github.com/forbole/flowJuno/db/utils"
	"github.com/forbole/flowJuno/types"
	dbtypes "github.com/forbole/flowJuno/db/types"

)

// SaveAccounts saves the given accounts inside the database
func (db *Database) SaveAccounts(accounts []types.Account) error {
	paramsNumber := 1
	slices := dbutils.SplitAccounts(accounts, paramsNumber)

	for _, accounts := range slices {
		if len(accounts) == 0 {
			continue
		}

		// Store up-to-date data
		err := db.saveAccounts(accounts)
		if err != nil {
			return fmt.Errorf("error while storing accounts: %s", err)
		}
	}

	return nil
}

func (db *Database) saveAccounts(accounts []types.Account) error {
	if len(accounts) == 0 {
		return nil
	}
	stmt := `INSERT INTO account (address) VALUES `
	var params []interface{}

	for i, account := range accounts {
		ai := i
		stmt += fmt.Sprintf("($%d),", ai+1)

		params = append(params, account.Address)
		stmt = stmt[:len(stmt)-1]
		stmt += " ON CONFLICT (address) DO NOTHING"
		_, err := db.Sql.Exec(stmt, params...)
		if err != nil {
			return err
		}
	}
	return nil
}

// GetAccounts returns all the accounts that are currently stored inside the database.
func (db *Database) GetAccounts() ([]types.Account, error) {
	var rows []dbtypes.AccountRow
	err := db.Select(&rows, `SELECT address FROM account`)
	if err != nil {
		return nil, err
	}

	returnRows := make([]types.Account, len(rows))
	for i, row := range rows {
		b := []byte(row.Details)

		if len(b) == 0 {
			returnRows[i] = types.NewAccount(row.Address)
		} else {
			//var inter interface{}
			
			returnRows[i] = types.NewAccount(row.Address)
		}
	}
	return returnRows, nil
}
