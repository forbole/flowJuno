package utils

import (
	"github.com/forbole/flowJuno/types"
)

const (
	maxPostgreSQLParams = 65535
)

func SplitAccounts(accounts []types.Account, paramsNumber int) [][]types.Account {
	maxBalancesPerSlice := maxPostgreSQLParams / paramsNumber
	slices := make([][]types.Account, len(accounts)/maxBalancesPerSlice+1)

	sliceIndex := 0
	for index, account := range accounts {
		slices[sliceIndex] = append(slices[sliceIndex], account)

		if index > 0 && index%(maxBalancesPerSlice-1) == 0 {
			sliceIndex++
		}
	}

	return slices
}

func SplitAccountKeyList(accountKeyList []types.AccountKeyList, paramsNumber int) [][]types.AccountKeyList {
	maxBalancesPerSlice := maxPostgreSQLParams / paramsNumber
	slices := make([][]types.AccountKeyList, len(accountKeyList)/maxBalancesPerSlice+1)

	sliceIndex := 0
	for index, account := range accountKeyList {
		slices[sliceIndex] = append(slices[sliceIndex], account)

		if index > 0 && index%(maxBalancesPerSlice-1) == 0 {
			sliceIndex++
		}
	}

	return slices
}

/*
func SplitBalances(balances []types.AccountBalance, paramsNumber int) [][]types.AccountBalance {
	maxBalancesPerSlice := maxPostgreSQLParams / paramsNumber
	slices := make([][]types.AccountBalance, len(balances)/maxBalancesPerSlice+1)

	sliceIndex := 0
	for index, balance := range balances {
		slices[sliceIndex] = append(slices[sliceIndex], balance)

		if index > 0 && index%(maxBalancesPerSlice-1) == 0 {
			sliceIndex++
		}
	}

	return slices
} */
