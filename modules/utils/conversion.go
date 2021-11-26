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
		strvalue, err := CadanceConvertString(val)
		if err != nil {
			return nil, err
		}
		table[i] = strvalue

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

func CadenceConvertUint32(value cadence.Value) (uint32, error) {
	val, ok := value.ToGoValue().(uint32)
	if !ok {
		return 0, fmt.Errorf("the cadance value is not a uint32 value")
	}

	return val, nil
}

func CadenceConvertUint8(value cadence.Value) (uint8, error) {
	val, ok := value.ToGoValue().(uint8)
	if !ok {
		return 0, fmt.Errorf("the cadance value is not a uint8 value")
	}

	return val, nil
}

func CadanceConvertString(val cadence.Value) (string, error) {
	strvalue, ok := val.ToGoValue().(string)
	if !ok {
		return "", fmt.Errorf("the cadance value is not a string value")
	}

	return strvalue, nil
}
