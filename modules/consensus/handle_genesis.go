package consensus

import (
	"fmt"

	"github.com/forbole/flowJuno/types"

	"github.com/forbole/flowJuno/client"
	db "github.com/forbole/flowJuno/db/postgresql"
	"github.com/rs/zerolog/log"
)

func HandleGenesis(height int32, db *db.Db, flowClient client.Proxy) error {
	log.Debug().Str("module", "consensus").Msg("parsing genesis")

	//get genesis block
	block,err:=flowClient.Client().GetBlockByHeight(flowClient.Ctx(),uint64(height))
	
	// Save the genesis time
	err = db.SaveGenesis(types.NewGenesis(block.Timestamp, int64(block.Height)))
	
	if err != nil {
		return fmt.Errorf("error while storing genesis time: %s", err)
	}

	return nil
}



