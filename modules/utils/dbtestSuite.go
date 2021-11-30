package utils

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

func TestDbProxyTestSuite(t *testing.T) {
	suite.Run(t, new(DbProxyTestSuite))
}

type DbProxyTestSuite struct {
	ProxyTestSuite
	Database *database.Db
}

func (suite *DbProxyTestSuite) SetupTest() {
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

 //Delete the public schema
_, err = bigDipperDb.Sql.Exec(`DROP SCHEMA public CASCADE;`)
suite.Require().NoError(err)
// Re-create the schema
_, err = bigDipperDb.Sql.Exec(`CREATE SCHEMA public;`)
suite.Require().NoError(err)
dirPath := path.Join(".","..","..","..","db","postgresql","schema")
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
	

	suite.Database = bigDipperDb
}

// getBlock builds, stores and returns a block for the provided height
func (suite *DbProxyTestSuite) getBlock(height int64) *flow.Block {
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
	err := suite.Database.SaveBlock(&block)
	suite.Require().NoError(err)
	return &block
}
