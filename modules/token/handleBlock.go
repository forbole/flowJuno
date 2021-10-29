package token

import (
	"fmt"

	"github.com/rs/zerolog/log"

	"github.com/forbole/flowJuno/client"
	"github.com/forbole/flowJuno/modules/utils"

	database "github.com/forbole/flowJuno/db/postgresql"
)

func HandleBlock(db *database.Db, height uint64, flowClient client.Proxy) error {
	log.Debug().Str("module", "token").Msg("Get token info from cadence")

	err := getCurrentSupply(db, height, flowClient)
	if err != nil {
		return err
	}
	return nil
}

func getCurrentSupply(db *database.Db, height uint64, flowClient client.Proxy) error {
	log.Trace().Str("module", "staking").Int64("height", int64(height)).
		Msg("updating get ProposedTable")

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
