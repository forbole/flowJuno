package postgresql_test

import (
	dbtypes "github.com/forbole/flowJuno/db/types"
)

func (suite *DbTestSuite) TestSaveConsensus_SaveSupply() {
	// Save the data
	err := suite.database.SaveSupply(20, 10)
	suite.Require().NoError(err)

	original := dbtypes.NewSupplyRow(10, 20)

	// Verify the data
	var rows []dbtypes.SupplyRow
	err = suite.database.Sqlx.Select(&rows, "SELECT * FROM supply")
	suite.Require().NoError(err)
	suite.Require().Len(rows, 1)
	suite.Require().True(rows[0].Equal(original))

	// ----------------------------------------------------------------------------------------------------------------

	// Try updating with a lower height
	err = suite.database.SaveSupply(6, 9)
	suite.Require().NoError(err)

	// Verify the data
	rows = []dbtypes.SupplyRow{}
	err = suite.database.Sqlx.Select(&rows, "SELECT * FROM supply")
	suite.Require().NoError(err)
	suite.Require().Len(rows, 1)
	suite.Require().True(rows[0].Equal(original), "updating with a lower height should not change the data")

	// ----------------------------------------------------------------------------------------------------------------

	// Try updating with the same height
	err = suite.database.SaveSupply(10, 10)
	suite.Require().NoError(err)

	// Verify the data
	expected := dbtypes.NewSupplyRow(10, 10)

	rows = []dbtypes.SupplyRow{}
	err = suite.database.Sqlx.Select(&rows, "SELECT * FROM supply")
	suite.Require().NoError(err)
	suite.Require().Len(rows, 1)
	suite.Require().True(rows[0].Equal(expected), "updating with same height should change the data")

	// ----------------------------------------------------------------------------------------------------------------

	// Try updating with a higher height
	err = suite.database.SaveSupply(20, 15)
	suite.Require().NoError(err)

	// Verify the data
	expected = dbtypes.NewSupplyRow(15, 20)

	rows = []dbtypes.SupplyRow{}
	err = suite.database.Sqlx.Select(&rows, "SELECT * FROM supply")
	suite.Require().NoError(err)
	suite.Require().Len(rows, 1)
	suite.Require().True(rows[0].Equal(expected), "updating with higher height should change the data")
}
