package postgresql

import (
	"fmt"

	dbtypes "github.com/forbole/flowJuno/db/types"
	dbutils "github.com/forbole/flowJuno/db/utils"
	"github.com/forbole/flowJuno/types"
)

// SaveAccounts saves the given accounts inside the database
func (db *Db) SaveAccounts(accounts []types.Account) error {
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

func (db *Db) saveAccounts(accounts []types.Account) error {
	if len(accounts) == 0 {
		return nil
	}
	stmt := `INSERT INTO account (address) VALUES `
	var params []interface{}

	for i, account := range accounts {
		ai := i
		stmt += fmt.Sprintf("($%d),", ai+1)

		params = append(params, account.Address)
	}
	stmt = stmt[:len(stmt)-1]
	stmt += " ON CONFLICT (address) DO NOTHING"
	_, err := db.Sqlx.Exec(stmt, params...)
	if err != nil {
		fmt.Println(stmt)
		return err
	}
	return nil
}

func (db *Db) SaveLockedTokenAccounts(accounts []types.LockedAccount) error {
	if len(accounts) == 0 {
		return nil
	}
	stmt := `INSERT INTO locked_account (account_address,locked_address,balance,unlock_limit) VALUES `
	var params []interface{}

	for i, account := range accounts {
		ai := i * 4
		stmt += fmt.Sprintf("($%d,$%d,$%d,$%d),", ai+1, ai+2, ai+3, ai+4)

		params = append(params, account.Address, account.LockedAddress, account.Balance, account.UnlockLimit)

	}

	stmt = stmt[:len(stmt)-1]
	stmt += " ON CONFLICT (account_address) DO NOTHING"
	_, err := db.Sqlx.Exec(stmt, params...)
	if err != nil {
		return err
	}
	return nil
}

func (db *Db) SaveDelegatorAccount(accounts []types.DelegatorAccount) error {
	if len(accounts) == 0 {
		return nil
	}
	stmt := `INSERT INTO delegator_account (account_address,delegator_id,delegator_node_id ,delegator_node_info) VALUES `
	var params []interface{}

	for i, account := range accounts {
		ai := i * 4
		stmt += fmt.Sprintf("($%d,$%d,$%d,$%d),", ai+1, ai+2, ai+3, ai+4)

		params = append(params, account.Address, account.DelegatorId, account.DelegatorNodeId, account.DelegatorNodeInfo)

	}

	stmt = stmt[:len(stmt)-1]
	stmt += " ON CONFLICT (account_address) DO NOTHING"
	_, err := db.Sqlx.Exec(stmt, params...)
	if err != nil {
		return err
	}
	return nil
}

// GetAccounts returns all the accounts that are currently stored inside the database.
func (db *Db) GetAccounts() ([]types.Account, error) {
	var rows []dbtypes.AccountRow
	err := db.Sqlx.Select(&rows, `SELECT address FROM account`)
	if err != nil {
		return nil, err
	}

	returnRows := make([]types.Account, len(rows))
	for i, row := range rows {
		returnRows[i] = types.NewAccount(row.Address)
	}
	return returnRows, nil
}