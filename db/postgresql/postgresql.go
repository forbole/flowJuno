package postgresql

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/onflow/flow-go-sdk"
	"github.com/rs/zerolog/log"

	"github.com/cosmos/cosmos-sdk/simapp/params"
	"github.com/lib/pq"

	_ "github.com/lib/pq" // nolint

	"github.com/desmos-labs/juno/db"
	"github.com/desmos-labs/juno/types"
)

// Builder creates a database connection with the given database connection info
// from config. It returns a database connection handle or an error if the
// connection fails.
func Builder(cfg types.DatabaseConfig, encodingConfig *params.EncodingConfig) (db.Database, error) {
	sslMode := "disable"
	if cfg.GetSSLMode() != "" {
		sslMode = cfg.GetSSLMode()
	}

	schema := "public"
	if cfg.GetSchema() != "" {
		schema = cfg.GetSchema()
	}

	connStr := fmt.Sprintf(
		"host=%s port=%d dbname=%s user=%s sslmode=%s search_path=%s",
		cfg.GetHost(), cfg.GetPort(), cfg.GetName(), cfg.GetUser(), sslMode, schema,
	)

	if cfg.GetPassword() != "" {
		connStr += fmt.Sprintf(" password=%s", cfg.GetPassword())
	}

	postgresDb, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// Set max open connections
	postgresDb.SetMaxOpenConns(cfg.GetMaxOpenConnections())
	postgresDb.SetMaxIdleConns(cfg.GetMaxIdleConnections())

	return &Database{Sql: postgresDb, EncodingConfig: encodingConfig}, nil
}

// type check to ensure interface is properly implemented
var _ db.Database = &Database{}

// Database defines a wrapper around a SQL database and implements functionality
// for data aggregation and exporting.
type Database struct {
	Sql            *sql.DB
	EncodingConfig *params.EncodingConfig
}

// LastBlockHeight implements db.Database
func (db *Database) LastBlockHeight() (int64, error) {
	var height int64
	err := db.Sql.QueryRow(`SELECT coalesce(MAX(height),0) AS height FROM block;`).Scan(&height)
	return height, err
}

// HasBlock implements db.Database
func (db *Database) HasBlock(height int64) (bool, error) {
	var res bool
	err := db.Sql.QueryRow(`SELECT EXISTS(SELECT 1 FROM block WHERE height = $1);`, height).Scan(&res)
	return res, err
}

// SaveBlock implements db.Database
func (db *Database) SaveBlock(block *flow.Block) error {
	stmt := `INSERT INTO block (height,id,parent_id ,collection_guarantees,timestamp) VALUES ($1,$2,$3,$4,$5) ON CONFLICT DO NOTHING`
    
	grauntees:=make([]string,len(block.CollectionGuarantees))
	for i,collectionGuarantee:=range block.CollectionGuarantees{
		grauntees[i]=collectionGuarantee.CollectionID.String()
	}
	_, err := db.Sql.Exec(stmt,
		block.Height, block.ID, block.ParentID,grauntees,block.Timestamp,
	)
	if err!=nil{
		return err
	}

	var params []interface{}
	stmt=`INSERT INTO block_seal (height,execution_receipt_id ,execution_receipt_signatures) VALUES ($1,$2,$3,$4) ON CONFLICT DO NOTHING`
	for i, seal:=range block.Seals{
		vi := i * 3
		stmt += fmt.Sprintf("($%d, $%d, $%d),", vi+1, vi+2,vi+3)	
		params=append(params, block.Height, seal.ExecutionReceiptID.String(),seal.ExecutionReceiptSignatures,
		)	
	}

	stmt = stmt[:len(stmt)-1] // Remove trailing ,
	stmt += " ON CONFLICT DO NOTHING"
	_, err = db.Sql.Exec(stmt, params...)

	return err
}

// SaveTx implements db.Database
func (db *Database) SaveTx(tx *types.Txs) error {
	sqlStatement := `
INSERT INTO transaction 
    (hash, height, success, messages, memo, signatures, signer_infos, fee, gas_wanted, gas_used, raw_log, logs) 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) ON CONFLICT DO NOTHING`

	var sigs = make([]string, len(tx.Signatures))
	for index, sig := range sigs {
		sigs[index] = sig
	}

	var msgs = make([]string, len(tx.Body.Messages))
	for index, msg := range tx.Body.Messages {
		bz, err := db.EncodingConfig.Marshaler.MarshalJSON(msg)
		if err != nil {
			return err
		}
		msgs[index] = string(bz)
	}
	msgsBz := fmt.Sprintf("[%s]", strings.Join(msgs, ","))

	feeBz, err := db.EncodingConfig.Marshaler.MarshalJSON(tx.AuthInfo.Fee)
	if err != nil {
		return fmt.Errorf("failed to JSON encode tx fee: %s", err)
	}

	var sigInfos = make([]string, len(tx.AuthInfo.SignerInfos))
	for index, info := range tx.AuthInfo.SignerInfos {
		bz, err := db.EncodingConfig.Marshaler.MarshalJSON(info)
		if err != nil {
			return err
		}
		sigInfos[index] = string(bz)
	}
	sigInfoBz := fmt.Sprintf("[%s]", strings.Join(sigInfos, ","))

	logsBz, err := db.EncodingConfig.Amino.MarshalJSON(tx.Logs)
	if err != nil {
		return err
	}

	_, err = db.Sql.Exec(sqlStatement,
		tx.TxHash, tx.Height, tx.Successful(),
		msgsBz, tx.Body.Memo, pq.Array(sigs),
		sigInfoBz, string(feeBz),
		tx.GasWanted, tx.GasUsed, tx.RawLog, string(logsBz),
	)
	return err
}

// HasValidator implements db.Database
func (db *Database) HasValidator(addr string) (bool, error) {
	var res bool
	stmt := `SELECT EXISTS(SELECT 1 FROM validator WHERE consensus_address = $1);`
	err := db.Sql.QueryRow(stmt, addr).Scan(&res)
	return res, err
}

// SaveValidators implements db.Database
func (db *Database) SaveValidators(validators []*types.Validator) error {
	if len(validators) == 0 {
		return nil
	}

	stmt := `INSERT INTO validator (consensus_address, consensus_pubkey) VALUES `

	var vparams []interface{}
	for i, val := range validators {
		vi := i * 2

		stmt += fmt.Sprintf("($%d, $%d),", vi+1, vi+2)
		vparams = append(vparams, val.ConsAddr, val.ConsPubKey)
	}

	stmt = stmt[:len(stmt)-1] // Remove trailing ,
	stmt += " ON CONFLICT DO NOTHING"
	_, err := db.Sql.Exec(stmt, vparams...)
	return err
}

func (db *Database) SaveNodeInfos(infos []*types.NodeInfo) error {
	if len(infos) == 0 {
		return nil
	}

	stmt := `INSERT INTO node_info (
		id ,role,networkingAddress,networkingKey ,stakingKey ,tokensStaked ,
		tokensCommitted ,tokensUnstaking ,tokensUnstaked ,
		tokensRewarded ,delegators ,delegatorIDCounter ,
		tokensRequestedToUnstake, initialWeight
	) VALUES `

	var vparams []interface{}
	for i, val := range infos {
		vi := i * 13
		

		stmt += fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d),",
		 vi+1, vi+2,vi+3,vi+4,vi+5,vi+6,
		 vi+7,vi+8,vi+9,vi+10,vi+11,vi+12,vi+13)
		vparams = append(vparams,val.Id ,val.Role,val.NetworkingAddress,val.NetworkingKey ,val.StakingKey ,
			val.TokensStaked ,val.TokensCommitted ,val.TokensUnstaking ,val.TokensUnstaked ,val.TokensRewarded ,
			val.Delegators ,val.DelegatorIDCounter ,val.TokensRequestedToUnstake,val. InitialWeight)
	}

	stmt = stmt[:len(stmt)-1] // Remove trailing ,
	stmt += " ON CONFLICT DO NOTHING"
	_, err := db.Sql.Exec(stmt, vparams...)
	return err
}

// SaveCommitSignatures implements db.Database
func (db *Database) SaveCommitSignatures(signatures []*types.CommitSig) error {
	if len(signatures) == 0 {
		return nil
	}

	stmt := `INSERT INTO pre_commit (validator_address, height, timestamp, voting_power, proposer_priority) VALUES `

	var sparams []interface{}
	for i, sig := range signatures {
		si := i * 5

		stmt += fmt.Sprintf("($%d, $%d, $%d, $%d, $%d),", si+1, si+2, si+3, si+4, si+5)
		sparams = append(sparams, sig.ValidatorAddress, sig.Height, sig.Timestamp, sig.VotingPower, sig.ProposerPriority)
	}

	stmt = stmt[:len(stmt)-1]
	stmt += " ON CONFLICT (validator_address, timestamp) DO NOTHING"
	_, err := db.Sql.Exec(stmt, sparams...)
	return err
}

// SaveMessage implements db.Database
func (db *Database) SaveMessage(msg *types.Message) error {
	stmt := `
INSERT INTO message(transaction_hash, index, type, value, involved_accounts_addresses) 
VALUES ($1, $2, $3, $4, $5)`

	_, err := db.Sql.Exec(stmt, msg.TxHash, msg.Index, msg.Type, msg.Value, pq.Array(msg.Addresses))
	return err
}

// Close implements db.Database
func (db *Database) Close() {
	err := db.Sql.Close()
	if err != nil {
		log.Error().Str("module", "psql database").Err(err).Msg("error while closing connection")
	}
}

// -------------------------------------------------------------------------------------------------------------------

// GetLastPruned implements db.PruningDb
func (db *Database) GetLastPruned() (int64, error) {
	var lastPrunedHeight int64
	err := db.Sql.QueryRow(`SELECT coalesce(MAX(last_pruned_height),0) FROM pruning LIMIT 1;`).Scan(&lastPrunedHeight)
	return lastPrunedHeight, err
}

// StoreLastPruned implements db.PruningDb
func (db *Database) StoreLastPruned(height int64) error {
	_, err := db.Sql.Exec(`DELETE FROM pruning`)
	if err != nil {
		return err
	}

	_, err = db.Sql.Exec(`INSERT INTO pruning (last_pruned_height) VALUES ($1)`, height)
	return err
}

// Prune implements db.PruningDb
func (db *Database) Prune(height int64) error {
	_, err := db.Sql.Exec(`DELETE FROM pre_commit WHERE height = $1`, height)
	if err != nil {
		return err
	}

	_, err = db.Sql.Exec(`
DELETE FROM message 
USING transaction 
WHERE message.transaction_hash = transaction.hash AND transaction.height = $1
`, height)
	return err
}

