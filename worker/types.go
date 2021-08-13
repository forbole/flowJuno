package worker

import (
	"github.com/cosmos/cosmos-sdk/simapp/params"

	"github.com/forbole/flowJuno/client"
	"github.com/forbole/flowJuno/db"
	"github.com/forbole/flowJuno/modules"
	"github.com/forbole/flowJuno/types"
)

type Config struct {
	EncodingConfig *params.EncodingConfig
	Queue          types.HeightQueue
	ClientProxy    *client.Proxy
	Database       db.Database
	Modules        []modules.Module
}

func NewConfig(
	queue types.HeightQueue,
	encodingConfig *params.EncodingConfig,
	clientProxy *client.Proxy,
	db db.Database,
	modules []modules.Module,
) *Config {
	return &Config{
		EncodingConfig: encodingConfig,
		Queue:          queue,
		ClientProxy:    clientProxy,
		Database:       db,
		Modules:        modules,
	}
}
