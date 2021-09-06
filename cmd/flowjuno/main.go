package main

import (
	"os"

	"github.com/forbole/flowJuno/cmd/parse"

	"github.com/forbole/flowJuno/modules/messages"

	"github.com/forbole/flowJuno/cmd"
	database "github.com/forbole/flowJuno/db/postgresql"
	"github.com/forbole/flowJuno/modules"
	"github.com/forbole/flowJuno/types/config"
)

func main() {
	// Config the runner
	config := cmd.NewConfig("flowjuno").
		WithParseConfig(parse.NewConfig().
			WithConfigParser(config.ParseConfig).
			WithDBBuilder(database.Builder).
			WithEncodingConfigBuilder(config.MakeEncodingConfig(nil)).
			WithRegistrar(modules.NewRegistrar(getAddressesParser())),
		)

	// Run the commands and panic on any error
	exec := cmd.BuildDefaultExecutor(config)
	err := exec.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func getAddressesParser() messages.MessageAddressesParser {
	return messages.DefaultMessagesParser
}
