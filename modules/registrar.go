package modules

import (
	"fmt"
	"reflect"

	"github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/forbole/flowJuno/client"
	"github.com/forbole/flowJuno/db"
	"github.com/forbole/flowJuno/db/postgresql"

	"github.com/forbole/flowJuno/modules/messages"
	"github.com/forbole/flowJuno/modules/modules"
	"github.com/forbole/flowJuno/modules/registrar"
	juno "github.com/forbole/flowJuno/types"

	"github.com/forbole/flowJuno/modules/auth"
	"github.com/forbole/flowJuno/modules/consensus"
	"github.com/forbole/flowJuno/modules/staking"
	"github.com/forbole/flowJuno/modules/token"
	"github.com/forbole/flowJuno/modules/telemetry"

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
	cfg juno.Config, encodingConfig *params.EncodingConfig, database db.Database, cp *client.Proxy, cfg types.Config
) modules.Modules {

	bigDipperBd := postgresql.Cast(database)
	fmt.Println("BuildModules")

	fmt.Println(reflect.TypeOf(r.parser))
	return []modules.Module{
		messages.NewModule(r.parser, encodingConfig.Marshaler, database),
		auth.NewModule(r.parser, *cp, encodingConfig, bigDipperBd),
		consensus.NewModule(r.parser, *cp, encodingConfig, bigDipperBd),
		staking.NewModule(r.parser, *cp, encodingConfig, bigDipperBd),
		token.NewModule(r.parser, *cp, encodingConfig, bigDipperBd),
		telemetry.NewModule(r.parser, *cp, encodingConfig, cfg),

	}
}
