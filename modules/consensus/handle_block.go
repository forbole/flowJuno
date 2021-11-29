package consensus

import (
	"github.com/rs/zerolog/log"

	"github.com/forbole/flowJuno/modules/messages"
	"github.com/onflow/flow-go-sdk"

	"github.com/forbole/flowJuno/client"
	database "github.com/forbole/flowJuno/db/postgresql"
	db "github.com/forbole/flowJuno/db/postgresql"
	consutils "github.com/forbole/flowJuno/modules/consensus/utils"

)

func HandleBlock(block *flow.Block, _ messages.MessageAddressesParser, db *db.Db, height int64, flowClient client.Proxy) error {
	err := updateBlockTimeFromGenesis(block, db)
	if err != nil {
		log.Error().Str("module", "consensus").Int64("height", int64(block.Height)).
			Err(err).Msg("error while updating block time from genesis")
	}

	return nil
}

// updateBlockTimeFromGenesis insert average block time from genesis
func updateBlockTimeFromGenesis(block *flow.Block, db *database.Db) error {
	log.Trace().Str("module", "consensus").Int64("height", int64(block.Height)).
		Msg("updating block time from genesis")

	blocktime,err:=consutils.GetGenesisBlockTime(db,*block)
	if err!=nil{
		return err
	}

	return db.SaveAverageBlockTimeGenesis(*blocktime)
}
