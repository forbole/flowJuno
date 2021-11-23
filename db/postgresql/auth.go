package postgresql

import (
	"fmt"

	dbtypes "github.com/forbole/flowJuno/db/types"
	"github.com/forbole/flowJuno/db/utils"
	dbutils "github.com/forbole/flowJuno/db/utils"
	"github.com/forbole/flowJuno/types"
)

// SaveAccounts saves the given accounts inside the database
func (db *Db) SaveAccounts(accounts []types.Account, height uint64) error {
	paramsNumber := 1
	slices := dbutils.SplitAccounts(accounts, paramsNumber)
	for _, accounts := range slices {
		if len(accounts) == 0 {
			continue
		}

		// Store up-to-date data
		err := db.saveAccounts(accounts, height)
		if err != nil {
			return fmt.Errorf("error while storing accounts: %s", err)
		}
	}

	return nil
}

func (db *Db) saveAccounts(accounts []types.Account, height uint64) error {
	if len(accounts) == 0 {
		return nil
	}

	stmt := `INSERT INTO account(address) VALUES `

	var params []interface{}

	for i, rows := range accounts {
		ai := i * 1
		stmt += fmt.Sprintf("($%d),", ai+1)

		params = append(params, rows.Address)

	}
	stmt = stmt[:len(stmt)-1]
	stmt += ` ON CONFLICT DO NOTHING`

	_, err := db.Sqlx.Exec(stmt, params...)
	if err != nil {
		return fmt.Errorf("fail to insert into account: %s",err)
	}

	stmt = `INSERT INTO account_balance (address,balance,code,contract_map,height) VALUES `
	var params2 []interface{}

	for i, account := range accounts {
		ai := i * 5
		stmt += fmt.Sprintf("($%d,$%d,$%d,$%d,$%d),", ai+1, ai+2, ai+3, ai+4, ai+5)

		params2 = append(params2, account.Address, account.Balance, account.Code, account.Contracts,height)
	}
	stmt = stmt[:len(stmt)-1]
	stmt += " ON CONFLICT (address) DO NOTHING "
	_, err = db.Sqlx.Exec(stmt, params2...)
	if err != nil {
		return fmt.Errorf("fail to insert into account_balance: %s",err)
	}

	var params3 []interface{}

	for _, rows := range accounts {
		splitedAccountKeyList:=utils.SplitAccountKeyList(rows.Keys,8)
		for _,keyList:=range splitedAccountKeyList{
			stmt = `INSERT INTO account_key_list(address,index,weight,revoked,sig_algo,hash_algo,public_key,sequence_number) VALUES `
			i := 0
			for _, accountKey := range keyList {
				ai := i * 8
				i++
				stmt += fmt.Sprintf("($%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d),", ai+1, ai+2, ai+3, ai+4, ai+5, ai+6, ai+7, ai+8)
				params3 = append(params3, rows.Address, accountKey.Index, accountKey.Weight, accountKey.Revoked, accountKey.SigAlgo, accountKey.HashAlgo, accountKey.PublicKey, accountKey.SequenceNumber)
			}

			stmt = stmt[:len(stmt)-1]
			stmt += ` ON CONFLICT DO NOTHING`
		
			_, err = db.Sqlx.Exec(stmt, params3...)
			if err != nil {
				return fmt.Errorf("fail to insert into account_key_list: %s",err)
			}
			params3=make([]interface{},0)
		}
	}
	
	return nil
}

func (db *Db) SaveLockedAccount(accounts []types.LockedAccount) error {
	stmt := `INSERT INTO locked_account(address,locked_address) VALUES `

	var params []interface{}

	for i, rows := range accounts {
		ai := i * 2
		stmt += fmt.Sprintf("($%d,$%d),", ai+1, ai+2)

		params = append(params, rows.Address, rows.LockedAddress)

	}
	stmt = stmt[:len(stmt)-1]
	stmt += ` ON CONFLICT DO NOTHING`

	_, err := db.Sqlx.Exec(stmt, params...)
	if err != nil {
		return err
	}
	return nil
}

func (db *Db) SaveLockedAccountBalance(accounts []types.LockedAccountBalance) error {
	if len(accounts) == 0 {
		return nil
	}

	stmt := `INSERT INTO locked_account_balance (locked_address,balance,unlock_limit,height) VALUES `
	var params2 []interface{}

	for i, account := range accounts {
		ai := i * 4
		stmt += fmt.Sprintf("($%d,$%d,$%d,$%d),", ai+1, ai+2, ai+3, ai+4)

		params2 = append(params2, account.LockedAddress, account.Balance, account.UnlockLimit, account.Height)

	}

	stmt = stmt[:len(stmt)-1]
	stmt += " ON CONFLICT DO NOTHING "
	_, err := db.Sqlx.Exec(stmt, params2...)
	if err != nil {
		return fmt.Errorf("psql error on locked_account_balance: %s",err)
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
	stmt += " ON CONFLICT DO NOTHING"
	_, err := db.Sqlx.Exec(stmt, params...)
	if err != nil {
		return fmt.Errorf("fail to save into psql table delegator_account: %s",err)
	}
	return nil
}

// GetAccounts returns all the addresses that are currently stored inside the database.
func (db *Db) GetAddresses() ([]string, error) {
	var rows []dbtypes.AccountRow
	err := db.Sqlx.Select(&rows, `SELECT address FROM account`)
	if err != nil {
		return nil, err
	}

	addresses := make([]string, len(rows))
	for i, rows := range rows {
		addresses[i] = rows.Address
	}

	return addresses, nil
}

func (db *Db) SaveStakerNodeId(stakerNodeId []types.StakerNodeId) error {
	stmt := `INSERT INTO staker_node_id(address,node_id) VALUES `

	var params []interface{}

	for i, rows := range stakerNodeId {
		ai := i * 2
		stmt += fmt.Sprintf("($%d,$%d),", ai+1, ai+2)

		params = append(params, rows.Address, rows.NodeId)

	}
	stmt = stmt[:len(stmt)-1]
	stmt += ` ON CONFLICT DO NOTHING`

	_, err := db.Sqlx.Exec(stmt, params...)
	if err != nil {
		return err
	}

	return nil
}
