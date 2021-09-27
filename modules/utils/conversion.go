package utils

import (
	"fmt"

	"github.com/onflow/cadence"
)

func CadenceConvertStringArray(value cadence.Value) ([]string, error) {
	valueArray, ok := value.(cadence.Array)
	if !ok {
		return nil, fmt.Errorf("the cadence value is not an array")
	}

	table := make([]string, len(valueArray.Values))
	for i, val := range valueArray.Values {
		table[i] = val.String()
	}
	return table, nil
}

func CadenceConvertUint64(value cadence.Value) (uint64, error) {
	val, ok := value.ToGoValue().(uint64)
	if !ok {
		return 0, fmt.Errorf("the cadance value is not a uint64 value")
	}

	return val, nil
}
