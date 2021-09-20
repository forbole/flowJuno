package postgresql_test

import (
	"encoding/json"

	"github.com/onflow/flow-go-sdk"

	dbtypes "github.com/forbole/flowJuno/db/types"
	"github.com/forbole/flowJuno/types"
)

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

	accounts := []types.LockedAccount{
		types.NewLockedAccount(address.String(), lockedAddress.String(), balance, unlockLimit),
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
	expectedAccountRow := dbtypes.NewLockedAccountRow(address.String(), lockedAddress.String(), int(balance), int(unlockLimit))

	var accountRows []dbtypes.LockedAccountRow
	err = suite.database.Sqlx.Select(&accountRows, `SELECT * FROM locked_account`)
	suite.Require().NoError(err)
	suite.Require().Len(accountRows, 1, "account table should contain only one row")

	suite.Require().True(expectedAccountRow.Equal(accountRows[0]))
}

func (suite *DbTestSuite) TestSaveDelegatorAccount() {
	baseAccount := suite.SaveAccount()
	address := baseAccount[0].Address
	delegatorId:=int64(8411)
	delegatorNodeId:="2cfab7e9163475282f67186b06ce6eea7fa0687d25dd9c7a84532f2016bc2e5e" 
	nodeInfo:=types.NewDelegatorNodeInfo(uint32(delegatorId),delegatorNodeId,0,0,0,0,0,0)
	
	expectedDelegatorNodeId:=`{"Id":8411,"NodeID":"2cfab7e9163475282f67186b06ce6eea7fa0687d25dd9c7a84532f2016bc2e5e","TokensCommitted":0,"TokensStaked":0,"TokensUnstaking":0,"TokensRewarded":0,"TokensUnstaked":0,"TokensRequestedToUnstake":0}`

	accounts := []types.DelegatorAccount{
		types.NewDelegatorAccount(address.String(),delegatorId, delegatorNodeId, nodeInfo),
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
	expectedDelegatorNodeId=`{"Id":8411,"NodeID":"2cfab7e9163475282f67186b06ce6eea7fa0687d25dd9c7a84532f2016bc2e5e","TokensCommitted":0,"TokensStaked":0,"TokensUnstaking":0,"TokensRewarded":0,"TokensUnstaked":0,"TokensRequestedToUnstake":0}`
	expectedAccountRow := dbtypes.NewDelegatorAccountRow(address.String(), delegatorId, delegatorNodeId, expectedDelegatorNodeId)

	var accountRows []dbtypes.DelegatorAccountRow
	err = suite.database.Sqlx.Select(&accountRows, `SELECT * FROM delegator_account`)
	
	suite.Require().NoError(err)
	suite.Require().Len(accountRows, 1, "account table should contain only one row")

	suite.Require().True(expectedAccountRow.Equal(accountRows[0]))

}

func (suite *DbTestSuite) TestSaveStakerAccount() {
	baseAccount := suite.SaveAccount()
	address := baseAccount[0].Address
	id:="e7df1454826425251716a703e907981672a43208ef3eabfc95d593673da778f6"
	stakerNodeInfo:=types.NewStakerNodeInfo(id,5,"34.211.45.12:3569",
	"3fa19db960a86a1722a2d8ffa9563bd1e7d905c91536860c64f9e808ef88862639112a1e872d7eef93dd91207d07dd7c043e2a1e80077b109290682250429f1f",
      "87d827de3e1b3541c394dcbbb6d76d98e3a7710d6740d28122c468b83f41002625e7a5788cabfcd6ce76b188f7f60de614364d4ab2932dfe0ed6f2d602bd551606ea31045ca2ccde9658a175ccd73da859ab17e56ad81ca4f6ef982c5968a7cb",
	0,20000000,0,0,0,[]uint32{},0,0,0)
	expectedStakerNodeInfo:=`{"Id":"e7df1454826425251716a703e907981672a43208ef3eabfc95d593673da778f6","Role":5,"NetworkingAddress":"34.211.45.12:3569","NetworkingKey":"3fa19db960a86a1722a2d8ffa9563bd1e7d905c91536860c64f9e808ef88862639112a1e872d7eef93dd91207d07dd7c043e2a1e80077b109290682250429f1f","StakingKey":"87d827de3e1b3541c394dcbbb6d76d98e3a7710d6740d28122c468b83f41002625e7a5788cabfcd6ce76b188f7f60de614364d4ab2932dfe0ed6f2d602bd551606ea31045ca2ccde9658a175ccd73da859ab17e56ad81ca4f6ef982c5968a7cb","TokensStaked":0,"TokensCommitted":20000000,"TokensUnstaking":0,"TokensUnstaked":0,"TokensRewarded":0,"Delegators":[],"DelegatorIDCounter":0,"TokensRequestedToUnstake":0,"InitialWeight":0}`

	stakerAccount:=types.NewStakerAccount(address.String(),"e7df1454826425251716a703e907981672a43208ef3eabfc95d593673da778f6" ,stakerNodeInfo)
	err := suite.database.SaveStakerAccounts([]types.StakerAccount{stakerAccount})
	suite.Require().NoError(err)

	expectedAccountRow:=dbtypes.NewStakerAccountRow(address.String(),"e7df1454826425251716a703e907981672a43208ef3eabfc95d593673da778f6",expectedStakerNodeInfo)

	var accountRows []dbtypes.StakerAccountRow
	err = suite.database.Sqlx.Select(&accountRows, `SELECT * FROM staker_account`)


	suite.Require().NoError(err)
	suite.Require().Len(accountRows, 1, "account table should contain only one row")

	suite.Require().True(expectedAccountRow.Equal(accountRows[0]))

}
