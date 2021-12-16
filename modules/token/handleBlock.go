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

	if _, err := scheduler.Every(1).Minute().StartImmediately().Do(func() {
		utils.WatchMethod(func() error {
			startEpoch, err := utils.CheckStartEpochEvent(flowClient)
			if err != nil {
				return err
			}
			if !startEpoch {
				return nil
			}

			return HandlePerodicOperation(db, flowClient)
		})
	}); err != nil {
		return err
	}

	return HandlePerodicOperation(db, flowClient)
}

func HandlePerodicOperation(db *database.Db, flowClient client.Proxy) error {
	height, err := flowClient.LatestHeight()
	if err != nil {
		return fmt.Errorf("Cannot get latest height: %s", err)
	}

	supply, err := getCurrentSupply(flowClient)
	if err != nil {
		return fmt.Errorf("Fail to handle perodic operation for Token:%s", err)
	}

	return db.SaveSupply(supply, uint64(height))
}

func getCurrentSupply(flowClient client.Proxy) (uint64, error) {

	script := fmt.Sprintf(`
	import FlowToken from %s
	pub fun main(): UFix64 {
		let supply = FlowToken.totalSupply
		return supply
	}`, flowClient.Contract().FlowToken)

	value, err := flowClient.Client().ExecuteScriptAtLatestBlock(flowClient.Ctx(), []byte(script), nil)
	if err != nil {
		return 0, fmt.Errorf("Error on getting token supply:%s", err)
	}

	supply, err := utils.CadenceConvertUint64(value)
	if err != nil {
		return 0, fmt.Errorf("Error on getting token supply:%s", err)
	}

	return supply, err

}
