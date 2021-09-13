module github.com/forbole/flowJuno

go 1.13

require (
	github.com/cosmos/cosmos-sdk v0.42.9
	github.com/desmos-labs/juno v0.0.0-20210820090829-4142e0029177
	github.com/go-co-op/gocron v0.3.3
	github.com/jmoiron/sqlx v1.2.1-0.20200324155115-ee514944af4b
	github.com/lib/pq v1.9.0
	github.com/onflow/cadence v0.18.0
	github.com/onflow/flow-go-sdk v0.21.0
	github.com/pelletier/go-toml v1.8.1
	github.com/proullon/ramsql v0.0.0-20181213202341-817cee58a244
	github.com/rs/zerolog v1.21.0
	github.com/spf13/cobra v1.1.3
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/tendermint v0.34.11
	github.com/ziutek/mymysql v1.5.4 // indirect
	google.golang.org/grpc v1.37.0
	gopkg.in/yaml.v2 v2.4.0
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776 // indirect
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
