package staking

import (
	"fmt"

	"github.com/go-co-op/gocron"
	"github.com/rs/zerolog/log"

	"github.com/forbole/flowJuno/modules/staking/stakingutils"

	"github.com/forbole/flowJuno/types"

	"github.com/onflow/flow-go-sdk"

	"github.com/forbole/flowJuno/client"
	"github.com/forbole/flowJuno/modules/utils"

	database "github.com/forbole/flowJuno/db/postgresql"
	db "github.com/forbole/flowJuno/db/postgresql"
)

func RegisterPeriodicOps(scheduler *gocron.Scheduler, db *database.Db, flowClient client.Proxy) error {
	log.Debug().Str("module", "staking").Msg("setting up periodic tasks")

	if _, err := scheduler.Every(1).Week().Tuesday().At("15:00").StartImmediately().Do(func() {
		utils.WatchMethod(func() error { return HandleStaking(db, flowClient) })
	}); err != nil {
		return err
	}

	return HandleStaking(db, flowClient)
}

func HandleStaking(db *db.Db, flowClient client.Proxy) error {
	block, err := flowClient.Client().GetLatestBlock(flowClient.Ctx(), false)
	if err != nil {
		return err
	}

	addresses, err := db.GetAddresses()
	if err != nil {
		return err
	}
	_, err = stakingutils.GetTable(block, db, flowClient)
	if err != nil {
		return err
	}

	nodeInfo, err := stakingutils.GetNodeInfosFromTable(block, db, flowClient)
	if err != nil {
		return err
	}

	if len(addresses) != 0 {
		err = stakingutils.GetDataFromAddresses(addresses, block, db, flowClient)
		if err != nil {
			return err
		}
	}

	err = stakingutils.GetDataWithNoArgs(block, db, int64(block.Height), flowClient)
	if err != nil {
		return err
	}

	err = stakingutils.GetDataFromNodeID(nodeInfo, block, db, flowClient)
	if err != nil {
		return err
	}

	err = stakingutils.GetDataFromNodeDelegatorID(nodeInfo, block, db, flowClient)
	if err != nil {
		return err
	}

	return nil
}
