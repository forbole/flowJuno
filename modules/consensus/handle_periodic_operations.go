package consensus

import (
	"fmt"

	"github.com/go-co-op/gocron"
	"github.com/rs/zerolog/log"

	database "github.com/forbole/flowJuno/db/postgresql"
	consutils "github.com/forbole/flowJuno/modules/consensus/utils"
	"github.com/forbole/flowJuno/modules/utils"
)

// Register registers the utils that should be run periodically
func Register(scheduler *gocron.Scheduler, db *database.Db) error {
	log.Debug().Str("module", "consensus").Msg("setting up periodic tasks")

	if _, err := scheduler.Every(1).Minute().Do(func() {
		utils.WatchMethod(func() error { return updateBlockTimeInMinute(db) })
	}); err != nil {
		return err
	}

	if _, err := scheduler.Every(1).Hour().StartImmediately().Do(func() {
		utils.WatchMethod(func() error { return updateBlockTimeInHour(db) })
	}); err != nil {
		return err
	}

	if _, err := scheduler.Every(1).Day().StartImmediately().Do(func() {
		utils.WatchMethod(func() error { return updateBlockTimeInDay(db) })
	}); err != nil {
		return err
	}

	return nil
}

// updateBlockTimeInMinute insert average block time in the latest minute
func updateBlockTimeInMinute(db *database.Db) error {
	log.Trace().Str("module", "consensus").Str("operation", "block time").
		Msg("updating block time in minutes")

	newBlockTime, err := consutils.GetBlockTimeInMinute(db)
	if err != nil {
		return fmt.Errorf("Fail to update block time:%s", err)
	}

	return db.SaveAverageBlockTimePerMin(*newBlockTime)
}

// updateBlockTimeInHour insert average block time in the latest hour
func updateBlockTimeInHour(db *database.Db) error {
	log.Trace().Str("module", "consensus").Str("operation", "block time").
		Msg("updating block time in hours")

	blocktime, err := consutils.GetBlockTimeInHour(db)
	if err != nil {
		return err
	}

	return db.SaveAverageBlockTimePerHour(*blocktime)
}

// updateBlockTimeInDay insert average block time in the latest minute
func updateBlockTimeInDay(db *database.Db) error {
	log.Trace().Str("module", "consensus").Str("operation", "block time").
		Msg("updating block time in days")

	blocktime, err := consutils.GetBlockTimeInDay(db)
	if err != nil {
		return err
	}
	return db.SaveAverageBlockTimePerDay(*blocktime)
}
