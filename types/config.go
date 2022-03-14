package types

import "github.com/pelletier/go-toml"

var (
	// Cfg represents the configuration to be used during the execution
	Cfg Config
)

// ConfigParser represents a function that allows to parse a file contents as a Config object
type ConfigParser = func(fileContents []byte) (Config, error)

type configToml struct {
	RPC       *rpcConfig       `toml:"rpc"`
	Grpc      *grpcConfig      `toml:"grpc"`
	Cosmos    *cosmosConfig    `toml:"cosmos"`
	Database  *databaseConfig  `toml:"database"`
	Logging   *loggingConfig   `toml:"logging"`
	Parsing   *parsingConfig   `toml:"parsing"`
	Pruning   *pruningConfig   `toml:"pruning"`
	Telemetry *telemetryConfig `toml:"telemetry"`
}

// DefaultConfigParser attempts to read and parse a flowjuno config from the given string bytes.
// An error reading or parsing the config results in a panic.
func DefaultConfigParser(configData []byte) (Config, error) {
	var cfg configToml
	err := toml.Unmarshal(configData, &cfg)
	return NewConfig(
		cfg.RPC,
		cfg.Grpc,
		cfg.Cosmos,
		cfg.Database,
		cfg.Logging,
		cfg.Parsing,
		cfg.Pruning,
		cfg.Telemetry,
	), err
}

// ---------------------------------------------------------------------------------------------------------------------

// Config represents the configuration to run flowjuno
type Config interface {
	GetRPCConfig() RPCConfig
	GetGrpcConfig() GrpcConfig
	GetCosmosConfig() CosmosConfig
	GetDatabaseConfig() DatabaseConfig
	GetLoggingConfig() LoggingConfig
	GetParsingConfig() ParsingConfig
	GetPruningConfig() PruningConfig
	GetTelemetryConfig() TelemetryConfig
}

var _ Config = &config{}

// Config defines all necessary flowjuno configuration parameters.
type config struct {
	RPC       RPCConfig       `toml:"rpc"`
	Grpc      GrpcConfig      `toml:"grpc"`
	Cosmos    CosmosConfig    `toml:"cosmos"`
	Database  DatabaseConfig  `toml:"database"`
	Logging   LoggingConfig   `toml:"logging"`
	Parsing   ParsingConfig   `toml:"parsing"`
	Pruning   PruningConfig   `toml:"pruning"`
	Telemetry TelemetryConfig `toml:"telemetry"`
}

// NewConfig builds a new Config instance
func NewConfig(
	rpcConfig RPCConfig, grpConfig GrpcConfig,
	cosmosConfig CosmosConfig, dbConfig DatabaseConfig,
	loggingConfig LoggingConfig, parsingConfig ParsingConfig,
	pruningConfig PruningConfig, telemetryConfig TelemetryConfig,
) Config {
	return &config{
		RPC:       rpcConfig,
		Grpc:      grpConfig,
		Cosmos:    cosmosConfig,
		Database:  dbConfig,
		Logging:   loggingConfig,
		Parsing:   parsingConfig,
		Pruning:   pruningConfig,
		Telemetry: telemetryConfig,
	}
}

// GetRPCConfig implements Config
func (c *config) GetRPCConfig() RPCConfig {
	return c.RPC
}

// GetGrpcConfig implements Config
func (c *config) GetGrpcConfig() GrpcConfig {
	return c.Grpc
}

// GetCosmosConfig implements Config
func (c *config) GetCosmosConfig() CosmosConfig {
	return c.Cosmos
}

// GetDatabaseConfig implements Config
func (c *config) GetDatabaseConfig() DatabaseConfig {
	return c.Database
}

// GetLoggingConfig implements Config
func (c *config) GetLoggingConfig() LoggingConfig {
	return c.Logging
}

// GetParsingConfig implements Config
func (c *config) GetParsingConfig() ParsingConfig {
	return c.Parsing
}

// GetPruningConfig implements Config
func (c *config) GetPruningConfig() PruningConfig {
	return c.Pruning
}

// GetTelemetryConfig implements Config
func (c *config) GetTelemetryConfig() TelemetryConfig {
	return c.Telemetry
}

// ---------------------------------------------------------------------------------------------------------------------

// GrpcConfig contains the configuration of the gRPC endpoint
type GrpcConfig interface {
	GetAddress() string
	IsInsecure() bool
}

var _ GrpcConfig = &grpcConfig{}

type grpcConfig struct {
	Address  string `toml:"address"`
	Insecure bool   `toml:"insecure"`
}

// NewGrpcConfig allows to build a new GrpcConfig instance
func NewGrpcConfig(address string, insecure bool) GrpcConfig {
	return &grpcConfig{
		Address:  address,
		Insecure: insecure,
	}
}

// GetAddress implements GrpcConfig
func (g *grpcConfig) GetAddress() string {
	return g.Address
}

// IsInsecure implements GrpcConfig
func (g *grpcConfig) IsInsecure() bool {
	return g.Insecure
}

// ---------------------------------------------------------------------------------------------------------------------

// RPCConfig contains the configuration of the RPC endpoint
type RPCConfig interface {
	GetClientName() string
	GetAddress() string
	GetContracts() string
}

var _ RPCConfig = &rpcConfig{}

type rpcConfig struct {
	ClientName string `toml:"client_name"`
	Address    string `toml:"address"`
	Contracts  string `toml:"contracts"`
}

// NewRPCConfig allows to build a new RPCConfig instance
func NewRPCConfig(clientName, address, contracts string) RPCConfig {
	return &rpcConfig{
		ClientName: clientName,
		Address:    address,
		Contracts:  contracts,
	}
}

// GetClientName implements RPCConfig
func (r *rpcConfig) GetClientName() string {
	return r.ClientName
}

// GetAddress implements RPCConfig
func (r *rpcConfig) GetAddress() string {
	return r.Address
}

// GetContract implenments RPCConfig
func (r *rpcConfig) GetContracts() string {
	return r.Contracts
}

// ---------------------------------------------------------------------------------------------------------------------

// CosmosConfig contains the data to configure the CosmosConfig SDK
type CosmosConfig interface {
	GetPrefix() string
	GetModules() []string
	GetGenesisHeight() uint64
}

var _ CosmosConfig = &cosmosConfig{}

type cosmosConfig struct {
	Prefix        string   `toml:"prefix"`
	Modules       []string `toml:"modules"`
	GenesisHeight uint64   `toml:"genesis_height"`
}

// NewCosmosConfig returns a new CosmosConfig instance
func NewCosmosConfig(prefix string, modules []string, genesisHeight uint64) CosmosConfig {
	return &cosmosConfig{
		Prefix:        prefix,
		Modules:       modules,
		GenesisHeight: genesisHeight,
	}
}

// GetPrefix implements CosmosConfig
func (c *cosmosConfig) GetPrefix() string {
	return c.Prefix
}

// GetModules implements CosmosConfig
func (c *cosmosConfig) GetModules() []string {
	return c.Modules
}

func (c *cosmosConfig) GetGenesisHeight() uint64 {
	return c.GenesisHeight
}

// ---------------------------------------------------------------------------------------------------------------------

// DatabaseConfig represents a generic database configuration
type DatabaseConfig interface {
	GetName() string
	GetHost() string
	GetPort() int64
	GetUser() string
	GetPassword() string
	GetSSLMode() string
	GetSchema() string
	GetMaxOpenConnections() int
	GetMaxIdleConnections() int
	GetPartitionSize() int
}

var _ DatabaseConfig = &databaseConfig{}

type databaseConfig struct {
	Name               string `toml:"name"`
	Host               string `toml:"host"`
	Port               int64  `toml:"port"`
	User               string `toml:"user"`
	Password           string `toml:"password"`
	SSLMode            string `toml:"ssl_mode"`
	Schema             string `toml:"schema"`
	MaxOpenConnections int    `toml:"max_open_connections"`
	MaxIdleConnections int    `toml:"max_idle_connections"`
	PartitionSize int `toml:"partition_size"`
}

func NewDatabaseConfig(
	name, host string, port int64, user string, password string,
	sslMode string, schema string,
	maxOpenConnections int, maxIdleConnections int, partitionSize int,
) DatabaseConfig {
	return &databaseConfig{
		Name:               name,
		Host:               host,
		Port:               port,
		User:               user,
		Password:           password,
		SSLMode:            sslMode,
		Schema:             schema,
		MaxOpenConnections: maxOpenConnections,
		MaxIdleConnections: maxIdleConnections,
		PartitionSize: partitionSize,
	}
}

// GetName implements DatabaseConfig
func (d *databaseConfig) GetName() string {
	return d.Name
}

// GetHost implements DatabaseConfig
func (d *databaseConfig) GetHost() string {
	return d.Host
}

// GetPort implements DatabaseConfig
func (d *databaseConfig) GetPort() int64 {
	return d.Port
}

// GetUser implements DatabaseConfig
func (d *databaseConfig) GetUser() string {
	return d.User
}

// GetPassword implements DatabaseConfig
func (d *databaseConfig) GetPassword() string {
	return d.Password
}

// GetSSLMode implements DatabaseConfig
func (d *databaseConfig) GetSSLMode() string {
	return d.SSLMode
}

// GetSchema implements DatabaseConfig
func (d *databaseConfig) GetSchema() string {
	return d.Schema
}

// GetMaxOpenConnections implements DatabaseConfig
func (d *databaseConfig) GetMaxOpenConnections() int {
	return d.MaxOpenConnections
}

// GetMaxIdleConnections implements DatabaseConfig
func (d *databaseConfig) GetMaxIdleConnections() int {
	return d.MaxIdleConnections
}

func (d *databaseConfig) GetPartitionSize() int {
	return d.PartitionSize
}
// ---------------------------------------------------------------------------------------------------------------------

// LoggingConfig represents the configuration for the logging part
type LoggingConfig interface {
	GetLogLevel() string
	GetLogFormat() string
}

var _ LoggingConfig = &loggingConfig{}

type loggingConfig struct {
	LogLevel  string `toml:"level"`
	LogFormat string `toml:"format"`
}

// NewLoggingConfig returns a new LoggingConfig instance
func NewLoggingConfig(level, format string) LoggingConfig {
	return &loggingConfig{
		LogLevel:  level,
		LogFormat: format,
	}
}

// GetLogLevel implements LoggingConfig
func (l *loggingConfig) GetLogLevel() string {
	return l.LogLevel
}

// GetLogFormat implements LoggingConfig
func (l *loggingConfig) GetLogFormat() string {
	return l.LogFormat
}

// ---------------------------------------------------------------------------------------------------------------------

// ParsingConfig represents the configuration of the parsing
type ParsingConfig interface {
	GetWorkers() int64
	ShouldParseNewBlocks() bool
	ShouldParseOldBlocks() bool
	ShouldParseGenesis() bool
	GetGenesisFilePath() string
	GetStartHeight() int64
	UseFastSync() bool
}

var _ ParsingConfig = &parsingConfig{}

type parsingConfig struct {
	Workers         int64  `toml:"workers"`
	ParseNewBlocks  bool   `toml:"listen_new_blocks"`
	ParseOldBlocks  bool   `toml:"parse_old_blocks"`
	GenesisFilePath string `toml:"genesis_file_path"`
	ParseGenesis    bool   `toml:"parse_genesis"`
	StartHeight     int64  `toml:"start_height"`
	FastSync        bool   `toml:"fast_sync"`
}

func NewParsingConfig(
	workers int64,
	parseNewBlocks, parseOldBlocks bool,
	parseGenesis bool, genesisFilePath string, startHeight int64, fastSync bool,
) ParsingConfig {
	return &parsingConfig{
		Workers:         workers,
		ParseOldBlocks:  parseOldBlocks,
		ParseNewBlocks:  parseNewBlocks,
		ParseGenesis:    parseGenesis,
		GenesisFilePath: genesisFilePath,
		StartHeight:     startHeight,
		FastSync:        fastSync,
	}
}

// GetWorkers implements ParsingConfig
func (p *parsingConfig) GetWorkers() int64 {
	return p.Workers
}

// ShouldParseNewBlocks implements ParsingConfig
func (p *parsingConfig) ShouldParseNewBlocks() bool {
	return p.ParseNewBlocks
}

// ShouldParseOldBlocks implements ParsingConfig
func (p *parsingConfig) ShouldParseOldBlocks() bool {
	return p.ParseOldBlocks
}

// ShouldParseGenesis implements ParsingConfig
func (p *parsingConfig) ShouldParseGenesis() bool {
	return p.ParseGenesis
}

func (p *parsingConfig) GetGenesisFilePath() string {
	return p.GenesisFilePath
}

// GetStartHeight implements ParsingConfig
func (p *parsingConfig) GetStartHeight() int64 {
	return p.StartHeight
}

// UseFastSync implements ParsingConfig
func (p *parsingConfig) UseFastSync() bool {
	return p.FastSync
}

// ---------------------------------------------------------------------------------------------------------------------

// PruningConfig contains the configuration of the pruning strategy
type PruningConfig interface {
	GetKeepRecent() int64
	GetKeepEvery() int64
	GetInterval() int64
}

var _ PruningConfig = &pruningConfig{}

type pruningConfig struct {
	KeepRecent int64 `toml:"keep_recent"`
	KeepEvery  int64 `toml:"keep_every"`
	Interval   int64 `toml:"interval"`
}

// NewPruningConfig returns a new PruningConfig
func NewPruningConfig(keepRecent, keepEvery, interval int64) PruningConfig {
	return &pruningConfig{
		KeepRecent: keepRecent,
		KeepEvery:  keepEvery,
		Interval:   interval,
	}
}

// GetKeepRecent implements PruningConfig
func (p *pruningConfig) GetKeepRecent() int64 {
	return p.KeepRecent
}

// GetKeepEvery implements PruningConfig
func (p *pruningConfig) GetKeepEvery() int64 {
	return p.KeepEvery
}

// GetInterval implements PruningConfig
func (p *pruningConfig) GetInterval() int64 {
	return p.Interval
}

// ---------------------------------------------------------------------------------------------------------------------

// PruningConfig contains the configuration of the pruning strategy
type TelemetryConfig interface {
	GetEnabled() bool
	GetPort() int64
}

var _ TelemetryConfig = &telemetryConfig{}

type telemetryConfig struct {
	Enabled bool  `toml:"enabled"`
	Port    int64 `toml:"port"`
}

func NewTelemetryConfig(enabled bool, port int64) TelemetryConfig {
	return &telemetryConfig{
		Enabled: enabled,
		Port:    port,
	}
}

func (t *telemetryConfig) GetEnabled() bool {
	return t.Enabled
}

func (t *telemetryConfig) GetPort() int64 {
	return t.Port
}
