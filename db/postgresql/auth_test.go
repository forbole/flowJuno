package postgresql_test

import (
	"encoding/json"

	"github.com/onflow/flow-go-sdk"

	dbtypes "github.com/forbole/flowJuno/db/types"
	"github.com/forbole/flowJuno/types"
)

func (suite *DbTestSuite) AddAccount(){

}

func (suite *DbTestSuite) TestSaveAccount() {
	emptyContracts := make(map[string][]byte)
	accountKey :=
		[]*flow.AccountKey{
			flow.NewAccountKey().SetWeight(1000).SetSigAlgo(2).SetHashAlgo(1),
		}
	address := flow.HexToAddress("0x1")
	balance := uint64(10)

	accounts := []flow.Account{
		{
			Address:   address,
			Balance:   balance,
			Code:      nil,
			Keys:      accountKey,
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

	expectedAccountKey, err := json.Marshal(accountKey)
	suite.Require().NoError(err)
	expectedEmptyContracts, err := json.Marshal(emptyContracts)
	suite.Require().NoError(err)
	expectedAccountRow := dbtypes.NewAccountRow(address.String(), float64(balance), "",
		string(expectedAccountKey), string(expectedEmptyContracts))

	suite.Require().True(expectedAccountRow.Equal(accountRows[0]))
}

func (suite *DbTestSuite) SaveAccount() []flow.Account {
	emptyContracts := make(map[string][]byte)
	accountKey :=
		[]*flow.AccountKey{
			flow.NewAccountKey().SetWeight(1000).SetSigAlgo(2).SetHashAlgo(1),
		}
	address := flow.HexToAddress("0x1")
	balance := uint64(10)

	accounts := []flow.Account{
		{
			Address:   address,
			Balance:   balance,
			Code:      nil,
			Keys:      accountKey,
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
	baseAccount := suite.SaveAccount()
	address := baseAccount[0].Address
	lockedAddress := flow.HexToAddress("0x2")
	balance := uint64(10)
	unlockLimit := uint64(20)

	accounts := []types.LockedAccountBalance{
		types.NewLockedAccountBalance(lockedAddress.String(), balance, unlockLimit,10),
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
	expectedAccountRow := dbtypes.NewLockedAccountBalanceRow(lockedAddress.String(), int(balance), int(unlockLimit),10)

	var accountRows []dbtypes.LockedAccountBalanceRow
	err = suite.database.Sqlx.Select(&accountRows, `SELECT * FROM locked_account`)
	suite.Require().NoError(err)
	suite.Require().Len(accountRows, 1, "account table should contain only one row")

	suite.Require().True(expectedAccountRow.Equal(accountRows[0]))
}

func (suite *DbTestSuite) TestSaveDelegatorAccount() {
	baseAccount := suite.SaveAccount()
	address := baseAccount[0].Address
	delegatorId := int64(8411)
	delegatorNodeId := "2cfab7e9163475282f67186b06ce6eea7fa0687d25dd9c7a84532f2016bc2e5e"
	nodeInfo := types.NewDelegatorNodeInfo(uint32(delegatorId), delegatorNodeId, 0, 0, 0, 0, 0, 0)

	expectedDelegatorNodeId := `{"Id":8411,"NodeID":"2cfab7e9163475282f67186b06ce6eea7fa0687d25dd9c7a84532f2016bc2e5e","TokensCommitted":0,"TokensStaked":0,"TokensUnstaking":0,"TokensRewarded":0,"TokensUnstaked":0,"TokensRequestedToUnstake":0}`

	accounts := []types.DelegatorAccount{
		types.NewDelegatorAccount(address.String(), delegatorId, delegatorNodeId, nodeInfo),
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
	expectedAccountRow := dbtypes.NewDelegatorAccountRow(address.String(), delegatorId, delegatorNodeId, expectedDelegatorNodeId)

	var accountRows []dbtypes.DelegatorAccountRow
	err = suite.database.Sqlx.Select(&accountRows, `SELECT * FROM delegator_account`)

	suite.Require().NoError(err)
	suite.Require().Len(accountRows, 1, "account table should contain only one row")

	suite.Require().True(expectedAccountRow.Equal(accountRows[0]))

}


func (suite *DbTestSuite) TestBigDipperDb_StakerNodeId(){

	// ------------------------------
	  // --- Prepare the data
	  // ------------------------------
  
	suite.Require().NoError(err)
  
	/*  TODO: Prepare parameter    */
  
	input:=[]types.StakerNodeId{
		  types.NewStakerNodeId(/*    TODO: Add parameter      */),
	  }
  
  
	  // ------------------------------
	  // --- Save the data
	  // ------------------------------
  
	err := suite.database.SaveStakerNodeId(input)
	suite.Require().NoError(err)
  
	// ------------------------------
	  // --- Verify the data
	  // ------------------------------
	  expectedRow := dbtypes.NewStakerNodeIdRow(/*    TODO: Add parameter      */)
	var outputs []dbtypes.StakerNodeIdRow
	  err = suite.database.Sqlx.Select(&outputs, `SELECT * FROM staker_node_id`)
	  suite.Require().NoError(err)
	  suite.Require().Len(outputs, 1, "should contain only one row")
	suite.Require().True(expectedRow.Equal(outputs[0]))
  
  }