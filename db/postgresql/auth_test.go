package postgresql_test

import (
	"github.com/onflow/flow-go-sdk"

	dbtypes "github.com/forbole/flowJuno/db/types"
)

func (suite *DbTestSuite) TestSaveAccount() {
	account := flow.Account{
		Address: flow.HexToAddress("0x1"),
		Balance: 10,
		Code:    nil,
		Keys: []*flow.AccountKey{
			&flow.AccountKey{
				Index:          0,
				PublicKey:      nil,
				SigAlgo:        2,
				HashAlgo:       1,
				Weight:         1000,
				SequenceNumber: 7935,
				Revoked:        false,
			},
		},
		Contracts: nil,
	}

	accounts := []flow.Account{
		account,
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

	expectedAccountRow := dbtypes.NewAccountRow("0x1", 10, "{}",
		`{"Index":0,"PublicKey":{},"SigAlgo":2,"HashAlgo":1,"Weight":1000,"SequenceNumber":7935,"Revoked":false}`,
		"{}")

	suite.Require().True(expectedAccountRow.Equal(accountRows[0]))
}
