package builder

import (
	"github.com/cosmos/cosmos-sdk/simapp/params"

	"github.com/forbole/flowJuno/types"

	"github.com/forbole/flowJuno/db"

	"github.com/forbole/flowJuno/db/postgresql"
)

// Builder represents a generic Builder implementation that build the proper database
// instance based on the configuration the user has specified
func Builder(cfg types.Config, encodingConfig *params.EncodingConfig) (db.Database, error) {
	return postgresql.Builder(cfg.GetDatabaseConfig(), encodingConfig)
}
