package types

type Account struct{
	Address string
}

func NewAccount(address string) Account{
	return Account{
		Address:address,
	}
}