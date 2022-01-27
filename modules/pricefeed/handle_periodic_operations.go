package pricefeed

import (
	"fmt"
	"strings"


	"github.com/go-co-op/gocron"
	"github.com/rs/zerolog/log"

	"github.com/forbole/flowJuno/client"
	"github.com/forbole/flowJuno/modules/utils"
	"github.com/forbole/flowJuno/types"

	"github.com/forbole/flowJuno/modules/pricefeed/coingecko"


	database "github.com/forbole/flowJuno/db/postgresql"
)

// RegisterPeriodicOperations implements modules.PeriodicOperationsModule
func RegisterPeriodicOps(scheduler *gocron.Scheduler, db *database.Db, flowClient client.Proxy) error {
	log.Debug().Str("module", "pricefeed").Msg("setting up periodic tasks")

	// Fetch total supply of token in 30 seconds each
	if _, err := scheduler.Every(30).Second().StartImmediately().Do(func() {
		utils.WatchMethod(func ()error{return updatePrice(db)})
	}); err != nil {
		return fmt.Errorf("error while setting up pricefeed period operations: %s", err)
	}

	return nil
}

// updatePrice fetch total amount of coins in the system from RPC and store it into database
func updatePrice(db *database.Db ) error {
	log.Debug().
		Str("module", "pricefeed").
		Str("operation", "pricefeed").
		Msg("getting token price and market cap")

	// Since this is specialise for Flow, it does not need to have any variable at
	coins :=[]coingecko.Token{
		coingecko.NewToken("FLOW","FLOW","Flow"),
	}
	// Get the list of token units
	units, err := db.GetTokenUnits()
	if err != nil {
		return fmt.Errorf("error while getting token units: %s", err)
	}

	// Find the id of the coins
	var ids []string
	for _, tradedToken := range units {
		// Skip the token if the price id is empty
		if tradedToken.PriceID == "" {
			continue
		}

		for _, coin := range coins {
			if strings.EqualFold(coin.ID, tradedToken.PriceID) {
				ids = append(ids, coin.ID)
				break
			}
		}
	}

	if len(ids) == 0 {
		log.Debug().Str("module", "pricefeed").Msg("no traded tokens found")
		return nil
	}

	// Get the tokens prices
	prices, err := coingecko.GetTokensPrices(ids)
	if err != nil {
		return fmt.Errorf("error while getting tokens prices: %s", err)
	}

	// Save the token prices
	err = db.SaveTokensPrices(prices)
	if err != nil {
		return fmt.Errorf("error while saving token prices: %s", err)
	}

	return updatePricesHistory(prices,db)
}

func updatePricesHistory(prices []types.TokenPrice,db *database.Db) error {

	return db.SaveTokenPricesHistory(prices)
}
