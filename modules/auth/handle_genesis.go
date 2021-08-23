package auth

import (
	"encoding/json"

	"github.com/rs/zerolog/log"

	"github.com/forbole/flowJuno/database"

	"github.com/cosmos/cosmos-sdk/codec"
)

// Handler handles the genesis state of the x/auth module in order to store the initial values
// of the different accounts.
func Handler(appState map[string]json.RawMessage, cdc codec.Marshaler, db *database.Db) error {
	log.Debug().Str("module", "auth").Msg("parsing genesis")

	return nil
}
