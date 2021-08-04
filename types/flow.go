package types

import (
	"fmt"
	"time"

	"github.com/onflow/flow-go-sdk"
	sdk "github.com/cosmos/cosmos-sdk/types"

)

// Tx represents an already existing blockchain transaction
type Tx struct{
	Events []flow.Event
	Status string
}

// NewTx allows to create a new Tx instance from the given txResponse
func NewTx(tx flow.TransactionResult)Tx{
	return Tx{
		Events:tx.Events,
		Status:tx.Status.String(),
	}
}


// FindEventByType searches inside the given tx events for the message having the specified index, in order
// to find the event having the given type, and returns it.
// If no such event is found, returns an error instead.
func (tx Tx) FindEventByType(index int, eventType string) (flow.Event, error) {
	for _, ev := range tx.Events{
		if ev.Type == eventType {
			return ev, nil
		}
	}

	return flow.Event{}, fmt.Errorf("no %s event found inside tx with hash %s", eventType, tx.TxHash)
}

// FindAttributeByKey searches inside the specified event of the given tx to find the attribute having the given key.
// If the specified event does not contain a such attribute, returns an error instead.
func (tx Tx) FindAttributeByKey(event flow.Event, attrKey string) (string, error) {
	for _, attr := range event.TransactionIndex{
		if attr.Key == attrKey {
			return attr.Value, nil
		}
	}

	return "", fmt.Errorf("no event with attribute %s found inside tx with hash %s", attrKey, tx.TxHash)
}

// Successful tells whether this tx is successful or not
func (tx Tx) Successful() bool {
	return tx.TxResponse.Code == 0
}
