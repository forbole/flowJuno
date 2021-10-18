package utils

import(
	"github.com/forbole/flowJuno/types"
)

const (
	maxPostgreSQLParams = 65535
)

func SplitDelegatorNodeInfo(inputarr []types.DelegatorNodeInfo, paramsNumber int) [][]types.DelegatorNodeInfo {
	maxBalancesPerSlice := maxPostgreSQLParams / paramsNumber

	if len(inputarr)<maxBalancesPerSlice{
		return [][]types.DelegatorNodeInfo{inputarr}
	}
	
	slices := make([][]types.DelegatorNodeInfo, len(inputarr)/maxBalancesPerSlice+1)

	sliceIndex := 0
	for index, element := range inputarr {
		slices[sliceIndex] = append(slices[sliceIndex], element)

		if index > 0 && index%(maxBalancesPerSlice-1) == 0 {
			sliceIndex++
		}
	}

	return slices
}