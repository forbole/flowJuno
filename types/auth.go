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
	Balance float64
	UnlockLimit float64
	Height int64
}

func NewLockedAccount(address string,lockedAddress string,balance float64,unlockLimit float64,height int64) LockedAccount{
	return LockedAccount{
		Address:address,
		LockedAddress:lockedAddress,
		Balance:balance,
		UnlockLimit: unlockLimit,
		Height:height,
	}
}

