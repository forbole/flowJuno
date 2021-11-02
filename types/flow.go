package types

import (
	"reflect"

	"github.com/onflow/cadence"
	"github.com/onflow/flow-go-sdk"
)

// Tx represents an already existing blockchain transaction
type Txs []Tx

type Tx struct {
	Height        uint64
	TransactionID string

	//Transaction Details
	Script             []byte
	Arguments          [][]byte
	ReferenceBlockID   string
	GasLimit           uint64
	ProposalKey        string
	Payer              string
	Authorizers        []string
	PayloadSignatures  []byte
	EnvelopeSignatures []byte
}

func NewTx( height uint64, transactionID string,
	script []byte, arguments [][]byte, referenceBlockID string,
	gasLimit uint64, proposalKey string, payer string, authorizers []string, payloadSignatures []byte,
	envelopeSignatures []byte) Tx {
	return Tx{
		Height:        height,
		TransactionID: transactionID,

		//Transaction Details
		Script:             script,
		Arguments:          arguments,
		ReferenceBlockID:   referenceBlockID,
		GasLimit:           gasLimit,
		ProposalKey:        proposalKey,
		Payer:              payer,
		Authorizers:        authorizers,
		PayloadSignatures:  payloadSignatures,
		EnvelopeSignatures: envelopeSignatures,
	}
}

type Event struct {
	//Transaction Result Event
	Height           int
	Type             string
	TransactionID    string
	TransactionIndex int
	EventIndex       int
	Value            cadence.Event
}

func NewEvent(height int, t string, transactionID string, transactionIndex int, eventIndex int,
	value cadence.Event) Event {

	return Event{
		//Transaction Result Event
		Height:           height,
		Type:             t,
		TransactionID:    transactionID,
		TransactionIndex: transactionIndex,
		EventIndex:       eventIndex,
		Value:            value,
	}
}

// Successful tells whether this tx is successful or not
func (tx Tx) Successful() bool {
	return true 
}

type Collection struct {
	Height         uint64
	Id             string
	Processed      bool
	TransactionIds []flow.Identifier
}

// Equal tells whether v and w represent the same rows
func (v Collection) Equal(w Collection) bool {
	return v.Height == w.Height &&
		v.Id == w.Id &&
		v.Processed == w.Processed &&
		reflect.DeepEqual(v.TransactionIds, w.TransactionIds)
}

// Collection allows to build a new Collection
func NewCollection(
	height uint64,
	id string,
	processed bool,
	transactionIds []flow.Identifier) Collection {

	return Collection{
		Height:         height,
		Id:             id,
		Processed:      processed,
		TransactionIds: transactionIds,
	}
}
