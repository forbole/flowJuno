package token

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/go-co-op/gocron"


	"github.com/forbole/flowJuno/client"
	"github.com/forbole/flowJuno/modules/utils"

	database "github.com/forbole/flowJuno/db/postgresql"
)

func RegisterPeriodicOps(scheduler *gocron.Scheduler, db *database.Db, flowClient client.Proxy) error {
	log.Debug().Str("module", "token").Msg("Get token info from cadence")

	if _, err := scheduler.Every(1).Week().Tuesday().At("15:00").StartImmediately().Do(func() {
		utils.WatchMethod(func() error { return getCurrentSupply(db,  flowClient)
		})
	}); err != nil {
		return err
	}
	return getCurrentSupply(db,  flowClient)
}

func getCurrentSupply(db *database.Db, flowClient client.Proxy) error {
	log.Trace().Str("module", "token").
		Msg("updating get ProposedTable")

	block,err:=flowClient.Client().GetLatestBlockHeader(flowClient.Ctx(),true)

	height:=block.Height
		script := fmt.Sprintf(`
	import FlowToken from %s
	pub fun main(): UFix64 {
		let supply = FlowToken.totalSupply
		return supply
	}`, flowClient.Contract().FlowToken)

	value, err := flowClient.Client().ExecuteScriptAtLatestBlock(flowClient.Ctx(), []byte(script), nil)
	if err != nil {
		return err
	}

	supply, err := utils.CadenceConvertUint64(value)
	if err != nil {
		return err
	}

	return db.SaveSupply(supply, height)

}
