package postgresql_test

import (
	"encoding/json"

	"github.com/onflow/flow-go-sdk"

	dbtypes "github.com/forbole/flowJuno/db/types"
	"github.com/forbole/flowJuno/types"
)

func (suite *DbTestSuite) AddAccount(address string) {
	_, err := suite.database.Sqlx.Exec(
		`INSERT INTO account(address) VALUES ($1)`, address)
	suite.Require().NoError(err)
}

func (suite *DbTestSuite) AddLockedAccount(address, lockedAddress string) {
	_, err := suite.database.Sqlx.Exec(
		`INSERT INTO account(address) VALUES ($1)`, address)
	suite.Require().NoError(err)
	_, err = suite.database.Sqlx.Exec(
		`INSERT INTO locked_account(address,locked_address) VALUES ($1,$2)`, address, lockedAddress)
	suite.Require().NoError(err)
}

func (suite *DbTestSuite) TestSaveAccount() {
	emptyContracts := make(map[string][]byte)
	accountKey :=
		[]*flow.AccountKey{
			flow.NewAccountKey().SetWeight(1000).SetSigAlgo(2).SetHashAlgo(1).SetWeight(1),
			flow.NewAccountKey().SetWeight(1000).SetSigAlgo(2).SetHashAlgo(1).SetWeight(2),

		}
	address := flow.HexToAddress("0x1")
	balance := uint64(10)

	flowAccount :=flow.Account{
			Address:   address,
			Balance:   balance,
			Code:      nil,
			Keys:      accountKey,
			Contracts: emptyContracts,
	}
	acc,err:=types.NewAccount(flowAccount)
	suite.Require().NoError(err)

	accounts:=[]types.Account{
		acc,
	}

	// ------------------------------
	// --- Save the data
	// ------------------------------

	err = suite.database.SaveAccounts(accounts,1)
	suite.Require().NoError(err)

	err = suite.database.SaveAccounts(accounts,1)
	suite.Require().NoError(err, "double account insertion should not insert and returns no error")

	// ------------------------------
	// --- Verify the data
	// ------------------------------
	expectedEmptyContracts, err := json.Marshal(emptyContracts)
	suite.Require().NoError(err)
	// Get Accounts row
	var accountRows []dbtypes.AccountRow
	err = suite.database.Sqlx.Select(&accountRows, `SELECT * FROM account`)
	suite.Require().NoError(err)
	suite.Require().Len(accountRows, 1, "account table should contain only one row")
	
	expectedAccountRow := dbtypes.NewAccountRow(address.String())
	suite.Require().True(expectedAccountRow.Equal(accountRows[0]))

	// Get Account Balance Row 
	var accountBalanceRows []dbtypes.AccountBalanceRow
	err = suite.database.Sqlx.Select(&accountRows, `SELECT * FROM account_balance`)
	suite.Require().NoError(err)
	suite.Require().Len(accountRows, 1, "account table should contain only one row")

	// Account Balance Row
	expectedAccountBalanceRow := dbtypes.NewAccountBalanceRow(address.String(), float64(balance), "",
		string(expectedEmptyContracts))
	suite.Require().True(expectedAccountBalanceRow.Equal(accountBalanceRows[0]))

	var accountKeyList []dbtypes.AccountKeyListRow
	err = suite.database.Sqlx.Select(&accountKeyList, `SELECT * FROM account_key_list`)
	suite.Require().NoError(err)
	suite.Require().Len(accountRows, 1, "account table should contain only one row")
}


func (suite *DbTestSuite) TestSaveLockedAccountBalance() {
	suite.AddAccount("0x1")
	suite.AddAccount(("0x2"))
	balance := uint64(10)
	unlockLimit := uint64(20)

	accounts := []types.LockedAccountBalance{
		types.NewLockedAccountBalance("0x1", balance, unlockLimit, 10),
	}

	// ------------------------------
	// --- Save the data
	// ------------------------------

	err := suite.database.SaveLockedAccountBalance(accounts)
	suite.Require().NoError(err)

	err = suite.database.SaveLockedAccountBalance(accounts)
	suite.Require().NoError(err, "double account insertion should not insert and returns no error")

	// ------------------------------
	// --- Verify the data
	// ------------------------------
	expectedAccountRow := dbtypes.NewLockedAccountBalanceRow("0x1", int(balance), int(unlockLimit), 10)

	var accountRows []dbtypes.LockedAccountBalanceRow
	err = suite.database.Sqlx.Select(&accountRows, `SELECT * FROM locked_account`)
	suite.Require().NoError(err)
	suite.Require().Len(accountRows, 1, "account table should contain only one row")

	suite.Require().True(expectedAccountRow.Equal(accountRows[0]))
}

func (suite *DbTestSuite) TestSaveDelegatorAccount() {
	suite.AddAccount("0x1")
	address := "0x1"
	delegatorId := int64(8411)
	delegatorNodeId := "2cfab7e9163475282f67186b06ce6eea7fa0687d25dd9c7a84532f2016bc2e5e"
	//nodeInfo := types.NewDelegatorNodeInfo(uint32(delegatorId), delegatorNodeId, 0, 0, 0, 0, 0, 0)

	expectedDelegatorNodeId := `{"Id":8411,"NodeID":"2cfab7e9163475282f67186b06ce6eea7fa0687d25dd9c7a84532f2016bc2e5e","TokensCommitted":0,"TokensStaked":0,"TokensUnstaking":0,"TokensRewarded":0,"TokensUnstaked":0,"TokensRequestedToUnstake":0}`

	accounts := []types.DelegatorAccount{
		types.NewDelegatorAccount(address, delegatorId, delegatorNodeId),
	}

	// ------------------------------
	// --- Save the data
	// ------------------------------

	err := suite.database.SaveDelegatorAccounts(accounts)
	suite.Require().NoError(err)

	err = suite.database.SaveDelegatorAccounts(accounts)
	suite.Require().NoError(err, "double account insertion should not insert and returns no error")

	// ------------------------------
	// --- Verify the data
	// ------------------------------
	expectedDelegatorNodeId = `{"Id":8411,"NodeID":"2cfab7e9163475282f67186b06ce6eea7fa0687d25dd9c7a84532f2016bc2e5e","TokensCommitted":0,"TokensStaked":0,"TokensUnstaking":0,"TokensRewarded":0,"TokensUnstaked":0,"TokensRequestedToUnstake":0}`
	expectedAccountRow := dbtypes.NewDelegatorAccountRow("0x1", delegatorId, delegatorNodeId, expectedDelegatorNodeId)

	var accountRows []dbtypes.DelegatorAccountRow
	err = suite.database.Sqlx.Select(&accountRows, `SELECT * FROM delegator_account`)

	suite.Require().NoError(err)
	suite.Require().Len(accountRows, 1, "account table should contain only one row")

	suite.Require().True(expectedAccountRow.Equal(accountRows[0]))

}
