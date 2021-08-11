package logging

import (
	"fmt"
	"os"

	"github.com/onflow/flow-go-sdk"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/desmos-labs/juno/modules"
	"github.com/desmos-labs/juno/types"
)

var _ Logger = &defaultLogger{}

// defaultLogger represents the default logger for any kind of error
type defaultLogger struct{}

// SetLogLevel implements Logger
func (d *defaultLogger) SetLogLevel(level string) error {
	logLvl, err := zerolog.ParseLevel(level)
	if err != nil {
		return err
	}

	zerolog.SetGlobalLevel(logLvl)
	return nil
}

// SetLogFormat implements Logger
func (d *defaultLogger) SetLogFormat(format string) error {
	switch format {
	case "json":
		// JSON is the default logging format
		break

	case "text":
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
		break

	default:
		return fmt.Errorf("invalid logging format: %s", format)
	}

	return nil
}

// LogGenesisError implements Logger
func (d *defaultLogger) LogGenesisError(module modules.Module, err error) {
	log.Error().Err(err).Str(LogKeyModule, module.Name()).
		Msg("error while handling genesis")
}

// LogBLockError implements Logger
func (d *defaultLogger) LogBLockError(module modules.Module, block *flow.Block, err error) {
	log.Error().Err(err).Str(LogKeyModule, module.Name()).Int64(LogKeyHeight, int64(block.Height)).
		Msg("error while handling block")
}

// LogTxError implements Logger
func (d *defaultLogger) LogTxError(module modules.Module, tx types.Txs, err error) {
	log.Error().Err(err).Str(LogKeyModule, module.Name()).Int64("height", int64(tx[0].Height)).
		Str(LogKeyTxHash,tx[0].TransactionID).Msg("error while handling transaction")
}
