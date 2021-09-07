package types

type Account struct{
	Address string
}

func NewAccount(address string) Account{
	return Account{
		Address:address,
	}
}

type LockedAccount struct{
	Address string
	LockedAddress string
}

func NewLockedAccount(address string,lockedAccount string) LockedAccount{
	return LockedAccount{
		Address:address,
		LockedAddress:lockedAccount,
	}
}