package postgresql_test

import (
	time "time"

	dbtypes "github.com/forbole/flowJuno/db/types"
	"github.com/forbole/flowJuno/types"
)

func (suite *DbTestSuite) TestSaveConsensus_GetBlockHeightTimeMinuteAgo() {
	timeAgo, err := time.Parse(time.RFC3339, "2020-01-01T15:00:00Z")
	suite.Require().NoError(err)

	var height int64 = 1000

	_, err = suite.database.Sql.Exec(`INSERT INTO block(height, id, parent_id, collection_guarantees, timestamp)
	VALUES ($1, 'b3889bc09945045abf0f8d7d0d9109467c7e45c5a8e98dd355394f4fc47b80f7', '2f708745fff4f66db88fac8f2f41d496edd341a2837d3e990e87679266e9bdb8', '[]', $2)`, height, timeAgo)
	suite.Require().NoError(err)

	timeNow := timeAgo.Add(time.Minute)
	result, err := suite.database.GetBlockHeightTimeMinuteAgo(timeNow)
	suite.Require().NoError(err)

	suite.Require().True(result.Timestamp.Equal(timeAgo))
	suite.Require().Equal(height, result.Height)
}

func (suite *DbTestSuite) TestSaveConsensus_GetBlockHeightTimeHourAgo() {
	timeAgo, err := time.Parse(time.RFC3339, "2020-01-01T15:00:00Z")
	suite.Require().NoError(err)

	var height int64 = 1000

	_, err = suite.database.Sql.Exec(`INSERT INTO block(height, id, parent_id, collection_guarantees, timestamp)
	VALUES ($1, 'b3889bc09945045abf0f8d7d0d9109467c7e45c5a8e98dd355394f4fc47b80f7', '2f708745fff4f66db88fac8f2f41d496edd341a2837d3e990e87679266e9bdb8', '[]', $2)`, height, timeAgo)
	suite.Require().NoError(err)

	timeNow := timeAgo.Add(time.Hour)
	result, err := suite.database.GetBlockHeightTimeHourAgo(timeNow)
	suite.Require().NoError(err)

	suite.Require().True(result.Timestamp.Equal(timeAgo))
	suite.Require().Equal(height, result.Height)
}

func (suite *DbTestSuite) TestSaveConsensus_GetBlockHeightTimeDayAgo() {
	timeAgo, err := time.Parse(time.RFC3339, "2020-01-01T15:00:00Z")
	suite.Require().NoError(err)

	var height int64 = 1000

	_, err = suite.database.Sql.Exec(`INSERT INTO block(height, id, parent_id, collection_guarantees, timestamp)
	VALUES ($1, 'b3889bc09945045abf0f8d7d0d9109467c7e45c5a8e98dd355394f4fc47b80f7', '2f708745fff4f66db88fac8f2f41d496edd341a2837d3e990e87679266e9bdb8', '[]', $2)`, height, timeAgo)
	suite.Require().NoError(err)

	timeNow := timeAgo.Add(time.Hour * 24)
	result, err := suite.database.GetBlockHeightTimeDayAgo(timeNow)
	suite.Require().NoError(err)

	suite.Require().True(result.Timestamp.Equal(timeAgo))
	suite.Require().Equal(height, result.Height)
}

func (suite *DbTestSuite) TestSaveConsensus_SaveAverageBlockTimePerMin() {
	// Save the data
	err := suite.database.SaveAverageBlockTimePerMin(types.NewBlockTime(10, 5.05))
	suite.Require().NoError(err)

	original := dbtypes.NewAverageTimeRow(5.05, 10)

	// Verify the data
	var rows []dbtypes.AverageTimeRow
	err = suite.database.Sqlx.Select(&rows, "SELECT * FROM average_block_time_per_minute")
	suite.Require().NoError(err)
	suite.Require().Len(rows, 1)
	suite.Require().True(rows[0].Equal(original))

	// ----------------------------------------------------------------------------------------------------------------

	// Try updating with a lower height
	err = suite.database.SaveAverageBlockTimePerMin(types.NewBlockTime(9, 6))
	suite.Require().NoError(err)

	// Verify the data
	rows = []dbtypes.AverageTimeRow{}
	err = suite.database.Sqlx.Select(&rows, "SELECT * FROM average_block_time_per_minute")
	suite.Require().NoError(err)
	suite.Require().Len(rows, 1)
	suite.Require().True(rows[0].Equal(original), "updating with a lower height should not change the data")

	// ----------------------------------------------------------------------------------------------------------------

	// Try updating with the same height
	err = suite.database.SaveAverageBlockTimePerMin(types.NewBlockTime(10, 10))
	suite.Require().NoError(err)

	// Verify the data
	expected := dbtypes.NewAverageTimeRow(10, 10)

	rows = []dbtypes.AverageTimeRow{}
	err = suite.database.Sqlx.Select(&rows, "SELECT * FROM average_block_time_per_minute")
	suite.Require().NoError(err)
	suite.Require().Len(rows, 1)
	suite.Require().True(rows[0].Equal(expected), "updating with same height should change the data")

	// ----------------------------------------------------------------------------------------------------------------

	// Try updating with a higher height
	err = suite.database.SaveAverageBlockTimePerMin(types.NewBlockTime(15, 20))
	suite.Require().NoError(err)

	// Verify the data
	expected = dbtypes.NewAverageTimeRow(20, 15)

	rows = []dbtypes.AverageTimeRow{}
	err = suite.database.Sqlx.Select(&rows, "SELECT * FROM average_block_time_per_minute")
	suite.Require().NoError(err)
	suite.Require().Len(rows, 1)
	suite.Require().True(rows[0].Equal(expected), "updating with higher height should change the data")
}

func (suite *DbTestSuite) TestSaveConsensus_SaveAverageBlockTimePerHour() {
	// Save the data
	err := suite.database.SaveAverageBlockTimePerHour(types.NewBlockTime(10, 5.05))
	suite.Require().NoError(err)

	original := dbtypes.NewAverageTimeRow(5.05, 10)

	// Verify the data
	var rows []dbtypes.AverageTimeRow
	err = suite.database.Sqlx.Select(&rows, "SELECT * FROM average_block_time_per_hour")
	suite.Require().NoError(err)
	suite.Require().Len(rows, 1)
	suite.Require().True(rows[0].Equal(original))

	// ----------------------------------------------------------------------------------------------------------------

	// Try updating with a lower height
	err = suite.database.SaveAverageBlockTimePerHour(types.NewBlockTime(9, 6))
	suite.Require().NoError(err)

	// Verify the data
	rows = []dbtypes.AverageTimeRow{}
	err = suite.database.Sqlx.Select(&rows, "SELECT * FROM average_block_time_per_hour")
	suite.Require().NoError(err)
	suite.Require().Len(rows, 1)
	suite.Require().True(rows[0].Equal(original), "updating with a lower height should not change the data")

	// ----------------------------------------------------------------------------------------------------------------

	// Try updating with the same height
	err = suite.database.SaveAverageBlockTimePerHour(types.NewBlockTime(10, 10))
	suite.Require().NoError(err)

	// Verify the data
	expected := dbtypes.NewAverageTimeRow(10, 10)

	rows = []dbtypes.AverageTimeRow{}
	err = suite.database.Sqlx.Select(&rows, "SELECT * FROM average_block_time_per_hour")
	suite.Require().NoError(err)
	suite.Require().Len(rows, 1)
	suite.Require().True(rows[0].Equal(expected), "updating with same height should change the data")

	// ----------------------------------------------------------------------------------------------------------------

	// Try updating with a higher height
	err = suite.database.SaveAverageBlockTimePerHour(types.NewBlockTime(15, 20))
	suite.Require().NoError(err)

	// Verify the data
	expected = dbtypes.NewAverageTimeRow(20, 15)

	rows = []dbtypes.AverageTimeRow{}
	err = suite.database.Sqlx.Select(&rows, "SELECT * FROM average_block_time_per_hour")
	suite.Require().NoError(err)
	suite.Require().Len(rows, 1)
	suite.Require().True(rows[0].Equal(expected), "updating with higher height should change the data")
}

func (suite *DbTestSuite) TestSaveConsensus_SaveAverageBlockTimePerDay() {
	// Save the data
	err := suite.database.SaveAverageBlockTimePerDay(types.NewBlockTime(10, 5.05))
	suite.Require().NoError(err)

	original := dbtypes.NewAverageTimeRow(5.05, 10)

	// Verify the data
	var rows []dbtypes.AverageTimeRow
	err = suite.database.Sqlx.Select(&rows, "SELECT * FROM average_block_time_per_day")
	suite.Require().NoError(err)
	suite.Require().Len(rows, 1)
	suite.Require().True(rows[0].Equal(original))

	// ----------------------------------------------------------------------------------------------------------------

	// Try updating with a lower height
	err = suite.database.SaveAverageBlockTimePerDay(types.NewBlockTime(9, 6))
	suite.Require().NoError(err)

	// Verify the data
	rows = []dbtypes.AverageTimeRow{}
	err = suite.database.Sqlx.Select(&rows, "SELECT * FROM average_block_time_per_day")
	suite.Require().NoError(err)
	suite.Require().Len(rows, 1)
	suite.Require().True(rows[0].Equal(original), "updating with a lower height should not change the data")

	// ----------------------------------------------------------------------------------------------------------------

	// Try updating with the same height
	err = suite.database.SaveAverageBlockTimePerDay(types.NewBlockTime(10, 10))
	suite.Require().NoError(err)

	// Verify the data
	expected := dbtypes.NewAverageTimeRow(10, 10)

	rows = []dbtypes.AverageTimeRow{}
	err = suite.database.Sqlx.Select(&rows, "SELECT * FROM average_block_time_per_day")
	suite.Require().NoError(err)
	suite.Require().Len(rows, 1)
	suite.Require().True(rows[0].Equal(expected), "updating with same height should change the data")

	// ----------------------------------------------------------------------------------------------------------------

	// Try updating with a higher height
	err = suite.database.SaveAverageBlockTimePerDay(types.NewBlockTime(15, 20))
	suite.Require().NoError(err)

	// Verify the data
	expected = dbtypes.NewAverageTimeRow(20, 15)

	rows = []dbtypes.AverageTimeRow{}
	err = suite.database.Sqlx.Select(&rows, "SELECT * FROM average_block_time_per_day")
	suite.Require().NoError(err)
	suite.Require().Len(rows, 1)
	suite.Require().True(rows[0].Equal(expected), "updating with higher height should change the data")
}

func (suite *DbTestSuite) TestSaveConsensus_SaveAverageBlockTimeGenesis() {
	// Save the data
	err := suite.database.SaveAverageBlockTimeGenesis(types.NewBlockTime(10, 5.05))
	suite.Require().NoError(err)

	original := dbtypes.NewAverageTimeRow(5.05, 10)

	// Verify the data
	var rows []dbtypes.AverageTimeRow
	err = suite.database.Sqlx.Select(&rows, "SELECT * FROM average_block_time_from_genesis")
	suite.Require().NoError(err)
	suite.Require().Len(rows, 1)
	suite.Require().True(rows[0].Equal(original))

	// ----------------------------------------------------------------------------------------------------------------

	// Try updating with a lower height
	err = suite.database.SaveAverageBlockTimeGenesis(types.NewBlockTime(9, 6))
	suite.Require().NoError(err)

	// Verify the data
	rows = []dbtypes.AverageTimeRow{}
	err = suite.database.Sqlx.Select(&rows, "SELECT * FROM average_block_time_from_genesis")
	suite.Require().NoError(err)
	suite.Require().Len(rows, 1)
	suite.Require().True(rows[0].Equal(original), "updating with a lower height should not change the data")

	// ----------------------------------------------------------------------------------------------------------------

	// Try updating with the same height
	err = suite.database.SaveAverageBlockTimeGenesis(types.NewBlockTime(10, 10))
	suite.Require().NoError(err)

	// Verify the data
	expected := dbtypes.NewAverageTimeRow(10, 10)

	rows = []dbtypes.AverageTimeRow{}
	err = suite.database.Sqlx.Select(&rows, "SELECT * FROM average_block_time_from_genesis")
	suite.Require().NoError(err)
	suite.Require().Len(rows, 1)
	suite.Require().True(rows[0].Equal(expected), "updating with same height should change the data")

	// ----------------------------------------------------------------------------------------------------------------

	// Try updating with a higher height
	err = suite.database.SaveAverageBlockTimeGenesis(types.NewBlockTime(15, 20))
	suite.Require().NoError(err)

	// Verify the data
	expected = dbtypes.NewAverageTimeRow(20, 15)

	rows = []dbtypes.AverageTimeRow{}
	err = suite.database.Sqlx.Select(&rows, "SELECT * FROM average_block_time_from_genesis")
	suite.Require().NoError(err)
	suite.Require().Len(rows, 1)
	suite.Require().True(rows[0].Equal(expected), "updating with higher height should change the data")
}

func (suite *DbTestSuite) TestSaveConsensus_SaveGenesisData() {
	err := suite.database.SaveGenesis(types.NewGenesis(
		time.Date(2020, 1, 02, 15, 00, 00, 000, time.UTC),
		0, "testnet-2",
	))
	suite.Require().NoError(err)

	// Should have only one row
	err = suite.database.SaveGenesis(types.NewGenesis(
		time.Date(2020, 1, 1, 15, 00, 00, 000, time.UTC),
		0,
		"testnet-2",
	))

	var rows []*dbtypes.GenesisRow
	err = suite.database.Sqlx.Select(&rows, "SELECT * FROM genesis")
	suite.Require().Len(rows, 1)
	suite.Require().True(rows[0].Equal(dbtypes.NewGenesisRow(
		time.Date(2020, 1, 1, 15, 00, 00, 000, time.UTC),
		0,
		"testnet-2",
	)))
}

func (suite *DbTestSuite) TestSaveConsensus_GetGenesis() {
	_, err := suite.database.Sqlx.Exec(
		`INSERT INTO genesis(chain_id, time, initial_height) VALUES ($1, $2, $3)`,
		"testnet-1",
		time.Date(2020, 1, 1, 15, 00, 00, 000, time.UTC),
		0,
	)

	genesis, err := suite.database.GetGenesis()
	suite.Require().NoError(err)
	suite.Require().True(genesis.Equal(types.NewGenesis(
		time.Date(2020, 1, 1, 15, 00, 00, 000, time.UTC),
		0,
		"testnet-1",
	)))
}
