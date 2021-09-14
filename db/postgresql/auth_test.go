package postgresql_test

import (
	"encoding/json"

	"github.com/onflow/flow-go-sdk"

	dbtypes "github.com/forbole/flowJuno/db/types"
	"github.com/forbole/flowJuno/types"
)

func (suite *DbTestSuite) TestSaveAccount() {
	emptyContracts :=make(map[string][]byte)
	accountKey:=
	[]*flow.AccountKey{
		flow.NewAccountKey().SetWeight(1000).SetSigAlgo(2).SetHashAlgo(1),
	}
	address:=flow.HexToAddress("0x1")
	balance:=uint64(10)

	accounts := []flow.Account{
		{
			Address: address,
		Balance: balance,
		Code:    nil,
		Keys: accountKey,
		Contracts: emptyContracts,
		},
	}

	// ------------------------------
	// --- Save the data
	// ------------------------------

	err := suite.database.SaveAccounts(accounts)
	suite.Require().NoError(err)

	err = suite.database.SaveAccounts(accounts)
	suite.Require().NoError(err, "double account insertion should not insert and returns no error")

	// ------------------------------
	// --- Verify the data
	// ------------------------------

	// Accounts row
	var accountRows []dbtypes.AccountRow
	err = suite.database.Sqlx.Select(&accountRows, `SELECT * FROM account`)
	suite.Require().NoError(err)
	suite.Require().Len(accountRows, 1, "account table should contain only one row")

	expectedAccountKey,err:=json.Marshal(accountKey)
	suite.Require().NoError(err)
	expectedEmptyContracts,err:=json.Marshal(emptyContracts)
	suite.Require().NoError(err)
	expectedAccountRow := dbtypes.NewAccountRow(address.String(), float64(balance), "",
	string(expectedAccountKey),string(expectedEmptyContracts))
	
	suite.Require().True(expectedAccountRow.Equal(accountRows[0]))
}

func (suite *DbTestSuite) SaveAccount() []flow.Account{
	emptyContracts :=make(map[string][]byte)
	accountKey:=
	[]*flow.AccountKey{
		flow.NewAccountKey().SetWeight(1000).SetSigAlgo(2).SetHashAlgo(1),
	}
	address:=flow.HexToAddress("0x1")
	balance:=uint64(10)

	accounts := []flow.Account{
		{
			Address: address,
		Balance: balance,
		Code:    nil,
		Keys: accountKey,
		Contracts: emptyContracts,
		},
	}

	// ------------------------------
	// --- Save the data
	// ------------------------------

	err := suite.database.SaveAccounts(accounts)
	suite.Require().NoError(err)

	return accounts
}

func (suite *DbTestSuite) TestSaveLockedAccount() {
	baseAccount:=suite.SaveAccount()
	address:=baseAccount[0].Address
	lockedAddress:=flow.HexToAddress("0x2")
	balance:=uint64(10)
	unlockLimit:=uint64(20)

	accounts:=[]types.LockedAccount{
		types.NewLockedAccount(address.String(),lockedAddress.String(),balance,unlockLimit),
	}
	

	// ------------------------------
	// --- Save the data
	// ------------------------------

	err := suite.database.SaveLockedTokenAccounts(accounts)
	suite.Require().NoError(err)

	err = suite.database.SaveLockedTokenAccounts(accounts)
	suite.Require().NoError(err, "double account insertion should not insert and returns no error")

	// ------------------------------
	// --- Verify the data
	// ------------------------------
	expectedAccountRow:=dbtypes.NewLockedAccountRow(address.String(),lockedAddress.String(),int(balance),int(unlockLimit))

	var accountRows []dbtypes.LockedAccountRow
	err = suite.database.Sqlx.Select(&accountRows, `SELECT * FROM locked_account`)
	suite.Require().NoError(err)
	suite.Require().Len(accountRows, 1, "account table should contain only one row")

	suite.Require().True(expectedAccountRow.Equal(accountRows[0]))
}


func (suite *DbTestSuite) TestSaveLockedAccount() {
	baseAccount:=suite.SaveAccount()
	address:=baseAccount[0].Address
	lockedAddress:=flow.HexToAddress("0x2")
	balance:=uint64(10)
	unlockLimit:=uint64(20)

	accounts:=[]types.LockedAccount{
		types.NewLockedAccount(address.String(),lockedAddress.String(),balance,unlockLimit),
	}
	

	// ------------------------------
	// --- Save the data
	// ------------------------------

	err := suite.database.SaveLockedTokenAccounts(accounts)
	suite.Require().NoError(err)

	err = suite.database.SaveLockedTokenAccounts(accounts)
	suite.Require().NoError(err, "double account insertion should not insert and returns no error")

	// ------------------------------
	// --- Verify the data
	// ------------------------------
	expectedAccountRow:=dbtypes.NewLockedAccountRow(address.String(),lockedAddress.String(),int(balance),int(unlockLimit))

	var accountRows []dbtypes.LockedAccountRow
	err = suite.database.Sqlx.Select(&accountRows, `SELECT * FROM locked_account`)
	suite.Require().NoError(err)
	suite.Require().Len(accountRows, 1, "account table should contain only one row")

	suite.Require().True(expectedAccountRow.Equal(accountRows[0]))
	
}
