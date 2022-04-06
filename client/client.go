package client

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/cosmos/cosmos-sdk/simapp/params"
	"github.com/cosmos/cosmos-sdk/types/tx"

	"github.com/forbole/flowJuno/types"

	"google.golang.org/grpc"

	"github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/client"
)

// Proxy implements a wrapper around both a Tendermint RPC client and a
// Cosmos Sdk REST client that allows for essential data queries.
type Proxy struct {
	ctx            context.Context
	encodingConfig *params.EncodingConfig
	contract       Contracts

	flowClient client.Client

	grpConnection   *grpc.ClientConn
	txServiceClient tx.ServiceClient
	genesisHeight   uint64
}

// NewClientProxy allows to build a new Proxy instance
func NewClientProxy(cfg types.Config, encodingConfig *params.EncodingConfig) (*Proxy, error) {
	flowClient, err := client.New(cfg.GetRPCConfig().GetAddress(), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	ctx := context.Background()

	var contracts Contracts
	if cfg.GetRPCConfig().GetContracts() == "Mainnet" {
		contracts = MainnetContracts()
	} else if cfg.GetRPCConfig().GetContracts() == "Testnet" {
		contracts = TestnetContracts()
	}

	return &Proxy{
		encodingConfig:  encodingConfig,
		ctx:             ctx,
		flowClient:      *flowClient,
		grpConnection:   nil,
		txServiceClient: nil,
		contract:        contracts,
		genesisHeight:   cfg.GetCosmosConfig().GetGenesisHeight(),
	}, nil
}

// GetGeneisisBlock parse the specific block as genesis block
func (cp *Proxy) GetGenesisHeight() uint64 {
	return cp.genesisHeight
}

// LatestHeight returns the latest block height on the active chain. An error
// is returned if the query fails.
func (cp *Proxy) LatestHeight() (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	block, err := cp.flowClient.GetLatestBlock(ctx, true)
	if err != nil {
		return -1, err
	}

	height := int64(block.Height)
	return height, nil
}

// Block queries for a block by height. An error is returned if the query fails.
func (cp *Proxy) Block(height int64) (*flow.Block, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	block, err := cp.flowClient.GetBlockByHeight(ctx, uint64(height))
	if err != nil {
		return nil, err
	}
	return block, nil
}

// GetTransaction queries for a transaction by hash. An error is returned if the
// query fails.
func (cp *Proxy) GetTransaction(hash string) (*flow.Transaction, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	transaction, err := cp.flowClient.GetTransaction(ctx, flow.HashToID([]byte(hash)))
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

/*
// NodeOperators returns all the known flow node operators for a given block
// height. An error is returned if the query fails.
func (cp *Proxy) NodeOperators(height int64) (*types.NodeOperators, error) {
	script := fmt.Sprintf(`
	import FlowIDTableStaking from %s
	pub fun main(): [FlowIDTableStaking.NodeInfo] {
		let nodes:[FlowIDTableStaking.NodeInfo]=[]
		for node in FlowIDTableStaking.getStakedNodeIDs() {
			nodes.append(FlowIDTableStaking.NodeInfo(node))
		}
		return nodes
	}`, cp.contract.StakingTable)

	result, err := cp.flowClient.ExecuteScriptAtLatestBlock(cp.ctx, []byte(script), nil)
	if err != nil {
		return nil, err
	}
	value := result.ToGoValue()
	nodes, ok := value.([]interface{})
	if !ok {
		return nil, fmt.Errorf("candance value cannot change to valid []interface{}")
	}
	nodeInfos := make([]*types.NodeInfo, len(nodes))
	for i, node := range nodes {
		nodeInfo, err := types.NewNodeInfoFromCandance(node)
		if err != nil {
			return nil, err
		}
		nodeInfos[i] = &nodeInfo
	}

	nodeOperators := types.NewNodeOperators(height, nodeInfos)
	return &nodeOperators, nil
} */

func (cp *Proxy) Client() *client.Client {
	return &cp.flowClient
}

func (cp *Proxy) Ctx() context.Context {
	return cp.ctx
}

func (cp *Proxy) Contract() Contracts {
	return cp.contract
}

func (cp *Proxy) GetChainID() string {
	// There is GetNetworkParameters rpc method that not implenment yet in flow-go-sdk.
	return cp.contract.ChainID
}

/*
// Genesis returns the genesis state
func (cp *Proxy) Genesis() (*tmctypes.ResultGenesis, error) {
	return cp.flowClient.Genesis(cp.ctx)
}

// ConsensusState returns the consensus state of the chain
func (cp *Proxy) ConsensusState() (*constypes.RoundStateSimple, error) {
	state, err := cp.flowClient.ConsensusState(context.Background())
	if err != nil {
		return nil, err
	}

	var data constypes.RoundStateSimple
	err = tmjson.Unmarshal(state.RoundState, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
*/
// SubscribeEvents subscribes to new events with the given query through the RPC
// client with the given subscriber name. A receiving only channel, context
// cancel function and an error is returned. It is up to the caller to cancel
// the context and handle any errors appropriately.
/* func (cp *Proxy) SubscribeEvents(subscriber, query string) (<-chan tmctypes.ResultEvent, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	eventCh, err := cp.flowClient.event
	return eventCh, cancel, err
}

// SubscribeNewBlocks subscribes to the new block event handler through the RPC
// client with the given subscriber name. An receiving only channel, context
// cancel function and an error is returned. It is up to the caller to cancel
// the context and handle any errors appropriately.
func (cp *Proxy) SubscribeNewBlocks(subscriber string) (<-chan tmctypes.ResultEvent, context.CancelFunc, error) {
	return cp.SubscribeEvents(subscriber, "tm.event = 'NewBlock'")
} */

// Collections get all the collection from block
func (cp *Proxy) Collections(block *flow.Block) []types.Collection {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collectionsID := block.CollectionGuarantees
	collections := make([]types.Collection, len(block.CollectionGuarantees))
	for i, c := range collectionsID {
		collection, err := cp.flowClient.GetCollection(ctx, c.CollectionID)

			// When it do not have a collection transaction yet at that block. It do not have a transaction ID
			// Retry until collection is produced

			sleeptime := 1
			for err != nil && strings.Contains(err.Error(), "please retry for collection in finalized block") {
				// It is not because block not finalized. Just unprocessable block
				if sleeptime>=16{
					break 
				}
				time.Sleep(time.Duration(sleeptime) * time.Second)
				collection, err = cp.flowClient.GetCollection(ctx, c.CollectionID)
				sleeptime = sleeptime * 2
			}

		
		if err != nil {
			collections[i] = types.NewCollection(block.Height, c.CollectionID.String(), false, nil)
			continue
		}

		collections[i] = types.NewCollection(block.Height, c.CollectionID.String(), true, collection.TransactionIDs)
	}
	return collections
}

// Txs queries for all the transactions in a block. Transactions are returned
// in the TransactionResult format which internally contains an array of Transactions. An error is
// returned if any query fails.
func (cp *Proxy) Txs(block *flow.Block) (types.Txs, error) {
	var transactionIDs []flow.Identifier
	collections := cp.Collections(block)
	for _, collection := range collections {
		transactionIDs = append(transactionIDs, (collection.TransactionIds)...)
	}

	txResponses := make([]types.Tx, len(transactionIDs))
	for i, txID := range transactionIDs {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		transaction, err := cp.flowClient.GetTransaction(ctx, txID)
		cancel()

		if err != nil {
			return nil, err
		}

		authoriser := make([]string, len(transaction.Authorizers))
		for i, auth := range transaction.Authorizers {
			authoriser[i] = auth.String()
		}

		payloadSignitures, err := json.Marshal(transaction.PayloadSignatures)
		if err != nil {
			return nil, err
		}

		envelopeSigniture, err := json.Marshal(transaction.EnvelopeSignatures)
		if err != nil {
			return nil, err
		}

		tx := types.NewTx(block.Height, txID.String(), transaction.Script, transaction.Arguments,
			transaction.ReferenceBlockID.String(), transaction.GasLimit, transaction.ProposalKey.Address.String(), transaction.Payer.String(),
			authoriser, payloadSignitures, envelopeSigniture)

		txResponses[i] = tx
	}
	return txResponses, nil
}

func (cp *Proxy) TransactionResult(transactionIds []flow.Identifier) ([]types.TransactionResult, error) {
	if len(transactionIds) == 0 {
		return nil, nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	txResults := make([]types.TransactionResult, len(transactionIds))
	for i, txid := range transactionIds {
		result, err := cp.flowClient.GetTransactionResult(ctx, txid)
		if err != nil {
			return nil, err
		}
		errStr := ""
		if result.Error != nil {
			errStr = result.Error.Error()
		}
		txResults[i] = types.NewTransactionResult(txid.String(), result.Status.String(), errStr)
	}
	return txResults, nil
}

func (cp *Proxy) EventsInBlock(block *flow.Block) ([]types.Event, error) {
	txs, err := cp.Txs(block)
	if err != nil {
		return nil, err
	}
	var event []types.Event
	for _, tx := range txs {
		ev, err := cp.Events(tx.TransactionID, int(tx.Height))
		if err != nil {
			return []types.Event{}, err
		}
		event = append(event, ev...)
	}
	return event, nil
}

func (cp *Proxy) EventsInTransaction(tx types.Tx) ([]types.Event, error) {
	var event []types.Event

	ev, err := cp.Events(tx.TransactionID, int(tx.Height))
	if err != nil {
		return []types.Event{}, err
	}
	event = append(event, ev...)
	return event, nil
}

// Events get events from a transaction ID
func (cp *Proxy) Events(transactionID string, height int) ([]types.Event, error) {
	sleeptime := 1
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(sleeptime)*time.Second)
	transactionResult, err := cp.flowClient.GetTransactionResult(ctx, flow.HexToID(transactionID))
	cancel()

	for err != nil && strings.Contains(err.Error(), "context deadline exceeded") {
		if sleeptime>=16{
			// fail
			return []types.Event{}, err
		}
		sleeptime = sleeptime * 2
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(sleeptime)*time.Second)
		transactionResult, err = cp.flowClient.GetTransactionResult(ctx, flow.HexToID(transactionID))
		cancel()
	}
	if err != nil {
		return []types.Event{}, err
	}

	ev := make([]types.Event, len(transactionResult.Events))
	for i, event := range transactionResult.Events {
		ev[i] = types.NewEvent(height, event.Type, event.TransactionID.String(), event.TransactionIndex,
			event.EventIndex, event.Value)
	}
	return ev, nil
}

// Stop defers the node stop execution to the RPC client.
func (cp *Proxy) Stop() {
	err := cp.flowClient.Close()
	if err != nil {
		log.Fatal().Str("module", "client proxy").Err(err).Msg("error while stopping proxy")
	}
}
