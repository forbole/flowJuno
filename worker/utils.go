package worker

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	tmctypes "github.com/tendermint/tendermint/rpc/core/types"
	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/forbole/flowJuno/types"
)

// findValidatorByAddr finds a validator by a consensus address given a set of
// Tendermint validators for a particular block. If no validator is found, nil
// is returned.
func findValidatorByAddr(consAddr string, vals *tmctypes.ResultValidators) *tmtypes.Validator {
	for _, val := range vals.Validators {
		if consAddr == sdk.ConsAddress(val.Address).String() {
			return val
		}
	}

	return nil
}

// sumGasTxs returns the total gas consumed by a set of transactions.
func sumGasTxs(txs types.Txs) uint64 {
	var totalGas uint64

	for _, tx := range txs {
		totalGas += tx.GasLimit
	}

	return totalGas
}

// WaitUntilQueueEmpty wait until queue is available
func WaitUntilQueueAvailable(queue types.HeightQueue) {
	sleeptime := 1
	for len(queue) == cap(queue) {
		time.Sleep(time.Second * time.Duration(sleeptime))
	}
}
