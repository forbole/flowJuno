package parse

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/forbole/flowJuno/client"
	"github.com/forbole/flowJuno/modules/modules"
	modsregistrar "github.com/forbole/flowJuno/modules/registrar"
	"github.com/forbole/flowJuno/types"
)

// SetupParsing setups all the things that should be later passed to StartParsing in order
// to parse the chain data properly.
func SetupParsing(parseConfig *Config) (*ParserData, error) {
	// Get the global config
	cfg := types.Cfg

	// Build the codec
	encodingConfig := parseConfig.GetEncodingConfigBuilder()()

	// Setup the SDK configuration
	sdkConfig := sdk.GetConfig()
	parseConfig.GetSetupConfig()(cfg, sdkConfig)
	sdkConfig.Seal()

	// Get the database
	database, err := parseConfig.GetDBBuilder()(cfg, &encodingConfig)
	if err != nil {
		return nil, err
	}

	// Init the client
	cp, err := client.NewClientProxy(cfg, &encodingConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to start client: %s", err)
	}

	// Setup the logging
	err = parseConfig.GetLogger().SetLogFormat(cfg.GetLoggingConfig().GetLogFormat())
	if err != nil {
		return nil, fmt.Errorf("error while setting logging format: %s", err)
	}

	err = parseConfig.GetLogger().SetLogLevel(cfg.GetLoggingConfig().GetLogLevel())
	if err != nil {
		return nil, fmt.Errorf("error while setting logging level: %s", err)
	}

	// Get the modules
	mods := parseConfig.GetRegistrar().BuildModules(cfg, &encodingConfig, sdkConfig, database, cp)
	registeredModules := modsregistrar.GetModules(mods, cfg.GetCosmosConfig().GetModules())
	logger := parseConfig.GetLogger()

	// Run all the additional operations
	for _, module := range registeredModules {
		if module, ok := module.(modules.AdditionalOperationsModule); ok {
			err := module.RunAdditionalOperations()
			if err != nil {
				return nil, err
			}
		}
	}

	return NewParserData(&encodingConfig, cp, database, registeredModules, logger), nil
}
