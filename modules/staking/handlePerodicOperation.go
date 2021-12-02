package staking

import (
	"fmt"

	"github.com/go-co-op/gocron"
	"github.com/rs/zerolog/log"

	"github.com/forbole/flowJuno/modules/staking/stakingutils"

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
	height := int64(block.Height)
	if err != nil {
		return fmt.Errorf("fail to handle staking:%s", err)
	}

	table, err := stakingutils.GetTable(int64(height), flowClient)
	if err != nil {
		return fmt.Errorf("fail to handle staking:%s", err)
	}

	err = db.SaveStakingTable(*table)
	if err != nil {
		return fmt.Errorf("fail to handle staking:%s", err)
	}

	nodeInfo, err := stakingutils.GetNodeInfosFromTable(height, flowClient)
	if err != nil {
		return fmt.Errorf("fail to handle staking:%s", err)
	}

	err = db.SaveNodeInfosFromTable(nodeInfo, block.Height)
	if err != nil {
		return fmt.Errorf("fail to handle staking:%s", err)
	}

	err = stakingutils.GetDataWithNoArgs(db, height, flowClient)
	if err != nil {
		return err
	}

	err = stakingutils.GetDataFromNodeID(nodeInfo, height, db, flowClient)
	if err != nil {
		return err
	}

	err = stakingutils.GetDataFromNodeDelegatorID(nodeInfo, height, db, flowClient)
	if err != nil {
		return err
	}

	return nil
}
