package database

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/forbole/flowJuno/logging"

	"github.com/onflow/flow-go-sdk"
	"github.com/rs/zerolog/log"

	"github.com/cosmos/cosmos-sdk/simapp/params"
	"github.com/lib/pq"

	"github.com/forbole/flowJuno/db"
	"github.com/forbole/flowJuno/types"
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
	Logger         logging.Logger
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

	grauntees := make([]string, len(block.CollectionGuarantees))
	for i, collectionGuarantee := range block.CollectionGuarantees {
		grauntees[i] = collectionGuarantee.CollectionID.String()
	}

	collectionGuarantees, err := json.Marshal(grauntees)
	if err != nil {
		return err
	}

	_, err = db.Sql.Exec(stmt,
		block.Height, block.ID.String(), block.ParentID.String(), collectionGuarantees, block.Timestamp,
	)
	if err != nil {
		return err
	}

	if len(block.Seals) == 0 {
		return nil
	}

	var params []interface{}
	stmt = `INSERT INTO block_seal (height,execution_receipt_id ,execution_receipt_signatures) VALUES `
	for i, seal := range block.Seals {
		vi := i * 3
		stmt += fmt.Sprintf("($%d, $%d, $%d),", vi+1, vi+2, vi+3)
		params = append(params, block.Height, seal.ExecutionReceiptID.String(), pq.ByteaArray(seal.ExecutionReceiptSignatures))
	}

	stmt = stmt[:len(stmt)-1] // Remove trailing ,
	stmt += " ON CONFLICT DO NOTHING"
	_, err = db.Sql.Exec(stmt, params...)

	return err
}

// SaveTx implements db.Database
func (db *Database) SaveTxs(txs types.Txs) error {
	if len(txs) == 0 {
		return nil
	}
	sqlStatement := `
INSERT INTO transaction 
    (height,transaction_id,script,arguments,reference_block_id,gas_limit,proposal_key ,payer,authorizers,payload_signature,envelope_signatures ) 
VALUES `

	var vparams []interface{}
	for i, tx := range txs {
		vi := i * 11
		sqlStatement += fmt.Sprintf(`($%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d),`,
			vi+1, vi+2, vi+3, vi+4, vi+5, vi+6, vi+7, vi+8, vi+9, vi+10, vi+11)
		vparams = append(vparams, tx.Height, tx.TransactionID, tx.Script, pq.ByteaArray(tx.Arguments), tx.ReferenceBlockID, tx.GasLimit, tx.ProposalKey, tx.Payer, pq.StringArray(tx.Authorizers),
			tx.PayloadSignatures, tx.EnvelopeSignatures)

	}
	sqlStatement = sqlStatement[:len(sqlStatement)-1] // Remove trailing ,

	sqlStatement += `ON CONFLICT DO NOTHING`
	_, err := db.Sql.Exec(sqlStatement, vparams...)
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

func (db *Database) SaveNodeInfos(infos []*types.StakerNodeInfo) error {
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
			vi+1, vi+2, vi+3, vi+4, vi+5, vi+6,
			vi+7, vi+8, vi+9, vi+10, vi+11, vi+12, vi+13)
		vparams = append(vparams, val.Id, val.Role, val.NetworkingAddress, val.NetworkingKey, val.StakingKey,
			val.TokensStaked, val.TokensCommitted, val.TokensUnstaking, val.TokensUnstaked, val.TokensRewarded,
			val.Delegators, val.DelegatorIDCounter, val.TokensRequestedToUnstake, val.InitialWeight)
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

func (db *Database) SaveEvents(events []types.Event) error {
	if len(events) == 0 {
		return nil
	}

	stmt := `INSERT INTO event (
		height,type,transaction_id,transaction_index,event_index,value
	) VALUES `

	var vparams []interface{}
	for i, event := range events {
		vi := i * 6

		stmt += fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d),",
			vi+1, vi+2, vi+3, vi+4, vi+5, vi+6)
		vparams = append(vparams, event.Height, event.Type, event.TransactionID, event.TransactionIndex, event.EventIndex, event.Value.String())
	}

	stmt = stmt[:len(stmt)-1] // Remove trailing ,
	stmt += " ON CONFLICT DO NOTHING"
	_, err := db.Sql.Exec(stmt, vparams...)
	return err
}

func (db *Database) SaveCollection(collection []types.Collection) error {
	stmt := `INSERT INTO collection(height,id,processed,transaction_id) VALUES `

	var params []interface{}

	i := 0
	for _, rows := range collection {
		for _, txid := range rows.TransactionIds {
			ai := i * 4
			stmt += fmt.Sprintf("($%d,$%d,$%d,$%d),", ai+1, ai+2, ai+3, ai+4)
			params = append(params, rows.Height, rows.Id, rows.Processed, txid.String())
			i++
		}
	}
	stmt = stmt[:len(stmt)-1]
	stmt += ` ON CONFLICT DO NOTHING`

	_, err := db.Sql.Exec(stmt, params...)
	if err != nil {
		return err
	}

	return nil
}
func (db *Database) SaveTransactionResult(transactionResult []types.TransactionResult, height uint64) error {
	stmt := `INSERT INTO transaction_result(height,transaction_id,status,error) VALUES `

	var params []interface{}

	for i, rows := range transactionResult {
		ai := i * 4
		stmt += fmt.Sprintf("($%d,$%d,$%d,$%d),", ai+1, ai+2, ai+3, ai+4)

		params = append(params, height, rows.TransactionId, rows.Status, rows.Error)

	}
	stmt = stmt[:len(stmt)-1]
	stmt += ` ON CONFLICT DO NOTHING`

	_, err := db.Sql.Exec(stmt, params...)
	if err != nil {
		return err
	}

	return nil
}

// createPartition allows to create a partition with the id for the given table name
func (db *Database) CreatePartition(table string, id int) error {
	stmt := fmt.Sprintf(
		"CREATE TABLE IF NOT EXISTS %v_%d PARTITION OF %v FOR VALUES IN (%d)",
		table,
		id,
		table,
		id,
	)
	_, err := db.Sql.Exec(stmt)
	return err
}

// dropPartition allows to drop a partition with the given partition name
func (db *Database) DropPartition(name string) error {
	stmt := fmt.Sprintf(
		"DROP TABLE IF EXISTS %v",
		name,
	)
	_, err := db.Sql.Exec(stmt)
	return err
}
