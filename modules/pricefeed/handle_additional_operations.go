package pricefeed

import (
	"fmt"
	"time"

	"github.com/forbole/flowJuno/types"

	"github.com/rs/zerolog/log"
	db "github.com/forbole/flowJuno/db/postgresql"

)

// checkConfig checks if the module config is valid
func checkConfig(cfg *Config) error {
	if cfg == nil {
		return fmt.Errorf("pricefeed config is not set but module is enabled")
	}

	return nil
}

// storeTokens stores the tokens defined inside the given configuration into the database
func storeTokens(cfg *Config,db *db.Db) error {
	log.Debug().Str("module", "pricefeed").Msg("storing tokens")

		// Save the coin as a token with its units
		coin:=types.NewToken("Flow",[]types.TokenUnit{
			types.NewTokenUnit("Flow",0,nil,"FLOW"),
		})
		err := db.SaveToken(coin)
		if err != nil {
			return fmt.Errorf("error while saving token: %s", err)
		}

		prices := []types.TokenPrice{
			types.NewTokenPrice("Flow", 0, 0, time.Time{}),
		}

	err = db.SaveTokensPrices(prices)
	if err != nil {
		return fmt.Errorf("error while storing token prices: %s", err)
	}

	return updatePricesHistory(prices,db)
}
