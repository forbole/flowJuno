package postgresql_test

import (
	"io/ioutil"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/forbole/flowJuno/types/config"
	"github.com/onflow/flow-go-sdk"

	database "github.com/forbole/flowJuno/db/postgresql"
	"github.com/forbole/flowJuno/types"

	juno "github.com/desmos-labs/juno/types"

	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/stretchr/testify/suite"

	_ "github.com/proullon/ramsql/driver"
)

func TestDatabaseTestSuite(t *testing.T) {
	suite.Run(t, new(DbTestSuite))
}

type DbTestSuite struct {
	suite.Suite

	database *database.Db
}

func (suite *DbTestSuite) SetupTest() {
	// Create the codec
	codec := simapp.MakeTestEncodingConfig()

	// Build the database
	cfg := types.NewConfig(
		nil, nil, nil,
		config.NewDatabaseConfig(
			juno.NewDatabaseConfig(
				"bdjuno",
				"localhost",
				5433,
				"bdjuno",
				"password",
				"",
				"public",
				-1,
				-1,
			),
			true,
		),
		nil, nil, nil,
	)

	db, err := database.Builder(cfg, &codec)
	suite.Require().NoError(err)

	bigDipperDb, ok := (db).(*database.Db)
	suite.Require().True(ok)

	// Delete the public schema
	_, err = bigDipperDb.Sql.Exec(`DROP SCHEMA public CASCADE;`)
	suite.Require().NoError(err)

	// Re-create the schema
	_, err = bigDipperDb.Sql.Exec(`CREATE SCHEMA public;`)
	suite.Require().NoError(err)

	dirPath := path.Join(".", "schema")
	dir, err := ioutil.ReadDir(dirPath)
	suite.Require().NoError(err)

	for _, fileInfo := range dir {
		file, err := ioutil.ReadFile(filepath.Join(dirPath, fileInfo.Name()))
		suite.Require().NoError(err)

		commentsRegExp := regexp.MustCompile(`/\*.*\*/`)
		requests := strings.Split(string(file), ";")
		for _, request := range requests {
			_, err := bigDipperDb.Sql.Exec(commentsRegExp.ReplaceAllString(request, ""))
			suite.Require().NoError(err)
		}
	}

	suite.database = bigDipperDb
}

// getBlock builds, stores and returns a block for the provided height
func (suite *DbTestSuite) getBlock(height int64) *flow.Block {
	block := flow.Block{
		flow.BlockHeader{
			ID:        flow.HexToID("0x1"),
			ParentID:  flow.HexToID("0x2"),
			Height:    10,
			Timestamp: time.Date(2021, 1, 1, 0, 0, 0, 0, time.Now().UTC().Location()),
		},
		flow.BlockPayload{
			[]*flow.CollectionGuarantee{
				&flow.CollectionGuarantee{
					flow.HexToID("0x3"),
				},
			},
			[]*flow.BlockSeal{
				&flow.BlockSeal{
					BlockID:                    flow.HexToID("0x4"),
					ExecutionReceiptID:         flow.HexToID("0x5"),
					ExecutionReceiptSignatures: nil,
				},
			},
		},
	}
	err := suite.database.SaveBlock(&block)
	suite.Require().NoError(err)
	return &block
}
/* 
// getValidator stores inside the database a validator having the given
// consensus address, validator address and validator public key
func (suite *DbTestSuite) getValidator(consAddr, valAddr, pubkey string) types.Validator {
	selfDelegation := suite.getAccount("cosmos1z4hfrxvlgl4s8u4n5ngjcw8kdqrcv43599amxs")

	maxRate := sdk.NewDec(10)
	maxChangeRate := sdk.NewDec(20)

	validator := types.NewValidator(
		consAddr,
		valAddr,
		pubkey,
		selfDelegation.String(),
		&maxChangeRate,
		&maxRate,
		1,
	)
	err := suite.database.SaveValidatorData(validator)
	suite.Require().NoError(err)

	return validator
}

// getAccount saves inside the database an account having the given address
func (suite *DbTestSuite) getAccount(addr string) sdk.AccAddress {
	delegator, err := sdk.AccAddressFromBech32(addr)
	suite.Require().NoError(err)

	_, err = suite.database.Sql.Exec(`INSERT INTO account (address) VALUES ($1) ON CONFLICT DO NOTHING`, delegator.String())
	suite.Require().NoError(err)

	return delegator
}
 */