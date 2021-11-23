package token

import (
	"fmt"

	"github.com/go-co-op/gocron"
	"github.com/rs/zerolog/log"

	"github.com/forbole/flowJuno/client"
	"github.com/forbole/flowJuno/modules/utils"

	database "github.com/forbole/flowJuno/db/postgresql"
)

func RegisterPeriodicOps(scheduler *gocron.Scheduler, db *database.Db, flowClient client.Proxy) error {
	log.Debug().Str("module", "staking").Msg("setting up periodic tasks")

	if _, err := scheduler.Every(1).Week().Tuesday().At("15:00").StartImmediately().Do(func() {
		utils.WatchMethod(func() error { return getCurrentSupply(db, flowClient) })
	}); err != nil {
		return err
	}

	return getCurrentSupply(db, flowClient)
}

func getCurrentSupply(db *database.Db, flowClient client.Proxy) error {
	height, err := flowClient.LatestHeight()
	if err != nil {
		return fmt.Errorf("Cannot get latest height: %s", err)
	}
	script := fmt.Sprintf(`
	import FlowToken from %s
	pub fun main(): UFix64 {
		let supply = FlowToken.totalSupply
		return supply
	}`, flowClient.Contract().FlowToken)

	value, err := flowClient.Client().ExecuteScriptAtLatestBlock(flowClient.Ctx(), []byte(script), nil)
	if err != nil {
		return fmt.Errorf("Error on getting token supply:%s", err)
	}

	supply, err := utils.CadenceConvertUint64(value)
	if err != nil {
		return fmt.Errorf("Error on getting token supply:%s", err)
	}

	return db.SaveSupply(supply, uint64(height))

}
