package postgresql

import (
	"encoding/json"
	"fmt"

	dbtypes "github.com/forbole/flowJuno/db/types"
	dbutils "github.com/forbole/flowJuno/db/utils"
	"github.com/forbole/flowJuno/types"
	"github.com/onflow/flow-go-sdk"
)

// SaveAccounts saves the given accounts inside the database
func (db *Db) SaveAccounts(accounts []flow.Account) error {
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

func (db *Db) saveAccounts(accounts []flow.Account) error {
	if len(accounts) == 0 {
		return nil
	}

	stmt := `INSERT INTO account(address) VALUES `

    var params []interface{}

	  for i, rows := range accounts{
      ai := i * 1
      stmt += fmt.Sprintf("($%d),", ai+1)
      
      params = append(params,rows.Address)

    }
	  stmt = stmt[:len(stmt)-1]
    stmt += ` ON CONFLICT DO NOTHING` 

    _, err := db.Sqlx.Exec(stmt, params...)
    if err != nil {
      return err
    }

	stmt = `INSERT INTO account_balance (address,balance,code,keys_list,contract_map) VALUES `
	var params2 []interface{}

	for i, account := range accounts {
		ai := i * 5
		stmt += fmt.Sprintf("($%d,$%d,$%d,$%d,$%d),", ai+1, ai+2, ai+3, ai+4, ai+5)

		keys, err := json.Marshal(account.Keys)
		if err != nil {
			return err
		}

		contracts, err := json.Marshal(account.Contracts)
		if err != nil {
			return err
		}
		params2 = append(params2, account.Address.String(), account.Balance, account.Code, keys, contracts)
	}
	stmt = stmt[:len(stmt)-1]
	stmt += " ON CONFLICT (address) DO UPDATE excluded. "
	_, err = db.Sqlx.Exec(stmt, params2...)
	if err != nil {
		fmt.Println(stmt)
		return err
	}
	return nil
}

func (db *Db) SaveLockedAccountBalance(accounts []types.LockedAccountBalance) error {
	if len(accounts) == 0 {
		return nil
	}

	stmt:= `INSERT INTO locked_account(address,locked_address) VALUES `

    var params []interface{}

	  for i, rows := range accounts{
      ai := i * 2
      stmt += fmt.Sprintf("($%d,$%d),", ai+1,ai+2)
      
      params = append(params,rows.Address,rows.LockedAddress)

    }
	  stmt = stmt[:len(stmt)-1]
    stmt += ` ON CONFLICT DO NOTHING` 

    _, err := db.Sqlx.Exec(stmt, params...)
    if err != nil {
      return err
    }

	stmt = `INSERT INTO locked_account_balance (locked_address,balance,unlock_limit,height) VALUES `
	var params2 []interface{}

	for i, account := range accounts {
		ai := i * 4
		stmt += fmt.Sprintf("($%d,$%d,$%d,$%d),", ai+1, ai+2, ai+3, ai+4)

		params2 = append(params2, account.LockedAddress, account.Balance, account.UnlockLimit,account.Height)

	}

	stmt = stmt[:len(stmt)-1]
	stmt += " ON CONFLICT (account_address) DO NOTHING "
	_, err = db.Sqlx.Exec(stmt, params2...)
	if err != nil {
		return err
	}
	return nil
}

func (db *Db) SaveDelegatorAccounts(accounts []types.DelegatorAccount) error {
	if len(accounts) == 0 {
		return nil
	}
	stmt := `INSERT INTO delegator_account (account_address,delegator_id,delegator_node_id ) VALUES `
	var params []interface{}

	for i, account := range accounts {
		ai := i * 4
		stmt += fmt.Sprintf("($%d,$%d,$%d),", ai+1, ai+2, ai+3)

		params = append(params, account.Address, account.DelegatorId, account.DelegatorNodeId)

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

func (db *Db) SaveLockedAccount(lockedAccount []types.LockedAccount) error {
    stmt:= `INSERT INTO locked_account(locked_address,node_id,delegator_id) VALUES `

    var params []interface{}

	  for i, rows := range lockedAccount{
      ai := i * 3
      stmt += fmt.Sprintf("($%d,$%d,$%d),", ai+1,ai+2,ai+3)
      
      params = append(params,rows.LockedAddress,rows.NodeId,rows.DelegatorId)

    }
	  stmt = stmt[:len(stmt)-1]
    stmt += ` ON CONFLICT DO NOTHING` 

    _, err := db.Sqlx.Exec(stmt, params...)
    if err != nil {
      return err
    }

    return nil 
    }

	func (db *Db) SaveStakerNodeId(stakerNodeId []types.StakerNodeId) error {
		stmt:= `INSERT INTO staker_node_id(address,node_id) VALUES `
	
		var params []interface{}
	
		  for i, rows := range stakerNodeId{
		  ai := i * 2
		  stmt += fmt.Sprintf("($%d,$%d),", ai+1,ai+2)
		  
		  params = append(params,rows.Address,rows.NodeId)
	
		}
		  stmt = stmt[:len(stmt)-1]
		stmt += ` ON CONFLICT DO NOTHING` 
	
		_, err := db.Sqlx.Exec(stmt, params...)
		if err != nil {
		  return err
		}
	
		return nil 
		}

