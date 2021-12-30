package parse

import (
	"fmt"

	"github.com/forbole/flowJuno/client"
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
	logger := parseConfig.GetLogger()
	mods := parseConfig.GetRegistrar().BuildModules(cfg, &encodingConfig, database, cp)
	registeredModules := modsregistrar.GetModules(mods, cfg.GetCosmosConfig().GetModules(), logger)


	return NewParserData(&encodingConfig, cp, database, registeredModules, logger), nil
}
