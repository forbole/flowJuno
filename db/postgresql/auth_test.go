package postgresql_test

import (
	"encoding/json"
	"fmt"

	"github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/crypto"

	dbtypes "github.com/forbole/flowJuno/db/types"
	"github.com/forbole/flowJuno/types"
)

func (suite *DbTestSuite) AddAccount(address string) error {
	_, err := suite.database.Sqlx.Exec(
		`INSERT INTO account(address) VALUES ($1)`, address)
	return err
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
	pubkeyString := "d0d45a9f40dc5e7440c71fcbc1a7836e7b38cc7874ac2875c5475fe550f582e005a5ed864c779d413fe49f58d3451c24ccf2cc12b9495c2baeeb0752538a0bcb"
	pubkey, err := crypto.DecodePublicKeyHex(crypto.ECDSA_P256, pubkeyString)
	suite.Require().NoError(err)

	emptyContracts := make(map[string][]byte)
	key := flow.NewAccountKey().SetWeight(1000).SetSigAlgo(crypto.ECDSA_P256).SetHashAlgo(crypto.SHA2_256).SetPublicKey(pubkey)

	accountKey :=
		[]*flow.AccountKey{
			key,
		}
	address := flow.HexToAddress("0x1")
	balance := uint64(10)

	flowAccount := flow.Account{
		Address:   address,
		Balance:   balance,
		Code:      nil,
		Keys:      accountKey,
		Contracts: emptyContracts,
	}
	acc, err := types.NewAccount(flowAccount)
	suite.Require().NoError(err)

	accounts := []types.Account{
		acc,
	}

	// ------------------------------
	// --- Save the data
	// ------------------------------

	err = suite.database.SaveAccounts(accounts, 1)
	suite.Require().NoError(err)

	err = suite.database.SaveAccounts(accounts, 1)
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
	err = suite.database.Sqlx.Select(&accountBalanceRows, `SELECT * FROM account_balance`)
	suite.Require().NoError(err)
	suite.Require().Len(accountBalanceRows, 1, "account table should contain only one row")

	// Test Account Balance Row
	expectedAccountBalanceRow := dbtypes.NewAccountBalanceRow(address.String(), float64(balance), "",
		string(expectedEmptyContracts), 1)
	suite.Require().True(expectedAccountBalanceRow.Equal(accountBalanceRows[0]))

	// Test Account Key List Row
	expectedAccountKeyList := dbtypes.NewAccountKeyListRow(address.String(), 0, 1000, false, "ECDSA_P256", "SHA2_256", "0x"+pubkeyString, key.SequenceNumber)
	var accountKeyList []dbtypes.AccountKeyListRow
	err = suite.database.Sqlx.Select(&accountKeyList, `SELECT * FROM account_key_list`)
	suite.Require().NoError(err)
	suite.Require().Len(accountKeyList, 1, "account table should contain only one row")
	fmt.Println(accountKeyList)
	suite.Require().True(expectedAccountKeyList.Equal(accountKeyList[0]))

}

func (suite *DbTestSuite) TestBigDipperDb_LockedAccount() {

	// ------------------------------
	// --- Prepare the data
	// ------------------------------

	suite.AddAccount("0x1")

	input := []types.LockedAccount{
		types.NewLockedAccount("0x1", "0x2"),
	}

	// ------------------------------
	// --- Save the data
	// ------------------------------

	err := suite.database.SaveLockedAccount(input)
	suite.Require().NoError(err)

	// ------------------------------
	// --- Verify the data
	// ------------------------------
	expectedRow := dbtypes.NewLockedAccountRow("0x1", "0x2")
	var outputs []dbtypes.LockedAccountRow
	err = suite.database.Sqlx.Select(&outputs, `SELECT * FROM locked_account`)
	suite.Require().NoError(err)
	suite.Require().Len(outputs, 1, "should contain only one row")
	suite.Require().True(expectedRow.Equal(outputs[0]))

}

func (suite *DbTestSuite) TestSaveLockedAccountBalance() {
	balance := uint64(10)
	unlockLimit := uint64(20)

	accounts := []types.LockedAccountBalance{
		types.NewLockedAccountBalance("0x1", balance, unlockLimit, 10),
	}

	suite.AddLockedAccount("0x1", "0x1")

	account2 := []types.LockedAccountBalance{
		types.NewLockedAccountBalance("0x1", balance, unlockLimit, 20),
	}
	// ------------------------------
	// --- Save the data
	// ------------------------------

	err := suite.database.SaveLockedAccountBalance(accounts)
	suite.Require().NoError(err)

	err = suite.database.SaveLockedAccountBalance(accounts)
	suite.Require().NoError(err, "double account insertion should not insert and returns no error")

	err = suite.database.SaveLockedAccountBalance(account2)
	suite.Require().NoError(err, "double account insertion should not insert and returns no error")

	// ------------------------------
	// --- Verify the data
	// ------------------------------
	expectedAccountRows := []dbtypes.LockedAccountBalanceRow{
		dbtypes.NewLockedAccountBalanceRow("0x1", int(balance), int(unlockLimit), 10),
		dbtypes.NewLockedAccountBalanceRow("0x1", int(balance), int(unlockLimit), 20),
	}

	var accountRows []dbtypes.LockedAccountBalanceRow
	err = suite.database.Sqlx.Select(&accountRows, `SELECT * FROM locked_account_balance`)
	suite.Require().NoError(err)
	suite.Require().Len(accountRows, 2, "account table should contain only two row")

	for i, row := range expectedAccountRows {
		suite.Require().True(row.Equal(accountRows[i]))

	}
}

func (suite *DbTestSuite) TestSaveDelegatorAccount() {
	suite.AddAccount("0x1")
	address := "0x1"
	delegatorId := int64(8411)
	delegatorNodeId := "2cfab7e9163475282f67186b06ce6eea7fa0687d25dd9c7a84532f2016bc2e5e"
	//nodeInfo := types.NewDelegatorNodeInfo(uint32(delegatorId), delegatorNodeId, 0, 0, 0, 0, 0, 0)

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
	expectedAccountRow := dbtypes.NewDelegatorAccountRow("0x1", delegatorId, delegatorNodeId)

	var accountRows []dbtypes.DelegatorAccountRow
	err = suite.database.Sqlx.Select(&accountRows, `SELECT * FROM delegator_account`)

	suite.Require().NoError(err)
	suite.Require().Len(accountRows, 1, "account table should contain only one row")

	suite.Require().True(expectedAccountRow.Equal(accountRows[0]))

}

func (suite *DbTestSuite) TestSaveStakerNodeId() {
	address := "0x1"
	delegatorNodeId := "2cfab7e9163475282f67186b06ce6eea7fa0687d25dd9c7a84532f2016bc2e5e"
	//nodeInfo := types.NewDelegatorNodeInfo(uint32(delegatorId), delegatorNodeId, 0, 0, 0, 0, 0, 0)
	err := suite.AddAccount(address)
	suite.Require().NoError(err)

	err = suite.InsertIntoStakingTable(1, delegatorNodeId)
	suite.Require().NoError(err)

	accounts := []types.StakerNodeId{
		types.NewStakerNodeId(address, delegatorNodeId),
	}

	// ------------------------------
	// --- Save the data
	// ------------------------------

	err = suite.database.SaveStakerNodeId(accounts)
	suite.Require().NoError(err)

	err = suite.database.SaveStakerNodeId(accounts)
	suite.Require().NoError(err, "double account insertion should not insert and returns no error")

	// ------------------------------
	// --- Verify the data
	// ------------------------------
	expectedAccountRow := dbtypes.NewStakerNodeIdRow("0x1", delegatorNodeId)

	var accountRows []dbtypes.StakerNodeIdRow
	err = suite.database.Sqlx.Select(&accountRows, `SELECT * FROM staker_node_id`)

	suite.Require().NoError(err)
	suite.Require().Len(accountRows, 1, "account table should contain only one row")

	suite.Require().True(expectedAccountRow.Equal(accountRows[0]))

}
