package parse

import (
	"github.com/spf13/cobra"

	"github.com/forbole/flowJuno/logging"

	"github.com/forbole/flowJuno/types"
)

// setupLogging setups the logging for the entire project
func setupLogging(_ *cobra.Command, _ []string) error {
	cfg := types.Cfg.GetLoggingConfig()

	err := logging.SetLogLevel(cfg.GetLogLevel())
	if err != nil {
		return err
	}

	err = logging.SetLogFormat(cfg.GetLogFormat())
	if err != nil {
		return err
	}

	return err
}
