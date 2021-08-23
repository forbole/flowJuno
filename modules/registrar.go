package modules

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/forbole/flowJuno/client"
	"github.com/forbole/flowJuno/db"
	"github.com/forbole/flowJuno/modules/messages"
	"github.com/forbole/flowJuno/modules/modules"
	"github.com/forbole/flowJuno/modules/registrar"
	juno "github.com/forbole/flowJuno/types"

	"github.com/forbole/flowJuno/modules/auth"
)

var (
	_ registrar.Registrar = &Registrar{}
)

// Registrar represents the modules.Registrar that allows to register all modules that are supported by BigDipper
type Registrar struct {
	parser messages.MessageAddressesParser
}

// NewRegistrar allows to build a new Registrar instance
func NewRegistrar(parser messages.MessageAddressesParser) *Registrar {
	return &Registrar{
		parser: parser,
	}
}

// BuildModules implements modules.Registrar
func (r *Registrar) BuildModules(
	cfg juno.Config, encodingConfig *params.EncodingConfig, _ *sdk.Config, database db.Database, cp *client.Proxy,
) modules.Modules{
	flowClient,err:= client.NewClientProxy(cfg,encodingConfig)

	if err!=nil{
		fmt.Errorf("Cannot connect to client Proxy")
		return nil
	}
	
	return []modules.Module{
		messages.NewModule(r.parser, encodingConfig.Marshaler, database),
		auth.NewModule(r.parser, *flowClient, encodingConfig, &database),
	}
}
