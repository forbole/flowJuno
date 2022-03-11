package postgresql

import (
	"fmt"

	juno "github.com/forbole/flowJuno/types"

	"github.com/cosmos/cosmos-sdk/simapp/params"

	"github.com/forbole/flowJuno/db"
	database "github.com/forbole/flowJuno/db/db"

	"github.com/jmoiron/sqlx"
)

var _ db.Database = &Db{}

// Db represents a PostgreSQL database with expanded features.
// so that it can properly store custom BigDipper-related data.
type Db struct {
	*database.Database
	Sqlx                *sqlx.DB
	storeHistoricalData bool
	PartitionSize int
}

// Builder allows to create a new Db instance implementing the db.Builder type
func Builder(cfg juno.Config, codec *params.EncodingConfig) (db.Database, error) {
	localdb, err := database.Builder(cfg.GetDatabaseConfig(), codec)
	if err != nil {
		return nil, err
	}

	psqlDb, ok := (localdb).(*database.Database)
	if !ok {
		return nil, fmt.Errorf("invalid configuration database, must be PostgreSQL")
	}
	/*
		dbCfg, ok := cfg.GetDatabaseConfig().(*config.DatabaseConfig)
		if !ok {
			return nil, fmt.Errorf("invalid database configuration type")
		}
	*/
	return &Db{
		Database:            psqlDb,
		Sqlx:                sqlx.NewDb(psqlDb.Sql, "postgresql"),
		storeHistoricalData: true,
		PartitionSize: cfg.GetDatabaseConfig().GetPartitionSize(),
	}, nil
}

// IsStoreHistoricDataEnabled tells whether or not the historical data should be stored inside the database
func (db *Db) IsStoreHistoricDataEnabled() bool {
	return db.storeHistoricalData
}

// Cast allows to cast the given db to a Db instance
func Cast(db db.Database) *Db {
	bdDatabase, ok := db.(*Db)
	if !ok {
		panic(fmt.Errorf("given database instance is not a Db"))
	}
	return bdDatabase
}
