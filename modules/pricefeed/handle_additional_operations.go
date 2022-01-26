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

	var prices []types.TokenPrice
	for _, coin := range cfg.Tokens {
		// Save the coin as a token with its units
		err := db.SaveToken(coin)
		if err != nil {
			return fmt.Errorf("error while saving token: %s", err)
		}

		// Create the price entry
		for _, unit := range coin.Units {
			// Skip units with empty price ids
			if unit.PriceID == "" {
				continue
			}

			prices = append(prices, types.NewTokenPrice(unit.Denom, 0, 0, time.Time{}))
		}
	}

	err := db.SaveTokensPrices(prices)
	if err != nil {
		return fmt.Errorf("error while storing token prices: %s", err)
	}

	return updatePricesHistory(prices,db)
}
