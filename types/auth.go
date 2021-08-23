package types

import "github.com/onflow/flow-go-sdk"

type Account struct{
	Address string
}

func NewAccount(address flow.Address) Account{
	return Account{
		Address:address.String(),
	}
}