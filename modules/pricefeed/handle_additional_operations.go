package pricefeed

import (
	"fmt"
	"time"

	"github.com/forbole/flowJuno/types"

	db "github.com/forbole/flowJuno/db/postgresql"
	"github.com/rs/zerolog/log"
)

// checkConfig checks if the module config is valid
func checkConfig(cfg *Config) error {
	if cfg == nil {
		return fmt.Errorf("pricefeed config is not set but module is enabled")
	}

	return nil
}

// storeTokens stores the tokens defined inside the given configuration into the database
func storeTokens(cfg *Config, db *db.Db) error {
	log.Debug().Str("module", "pricefeed").Msg("storing tokens")

	// Save the coin as a token with its units
	coin := types.NewToken("flow", []types.TokenUnit{
		types.NewTokenUnit("flow", 0, nil, "flow"),
	})
	err := db.SaveToken(coin)
	if err != nil {
		return fmt.Errorf("error while saving token: %s", err)
	}

	prices := []types.TokenPrice{
		types.NewTokenPrice("flow", 0, 0, time.Time{}),
	}

	err = db.SaveTokensPrices(prices)
	if err != nil {
		return fmt.Errorf("error while storing token prices: %s", err)
	}

	return updatePricesHistory(prices, db)
}
