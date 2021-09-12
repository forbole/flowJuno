package types

type Account struct {
	Address string
}

func NewAccount(address string) Account {
	return Account{
		Address: address,
	}
}

type LockedAccount struct {
	Address       string
	LockedAddress string
	Balance       uint64
	UnlockLimit   uint64
}

func NewLockedAccount(address string, lockedAddress string, balance, unlockLimit uint64) LockedAccount {
	return LockedAccount{
		Address:       address,
		LockedAddress: lockedAddress,
		Balance:       balance,
		UnlockLimit:   unlockLimit,
	}
}

type DelegatorAccount struct {
	Address           string
	DelegatorId       int64
	DelegatorNodeId   string
	DelegatorNodeInfo string
}

func NewDelegatorAccount(address string, delegatorId int64, delegatorNodeId string, delegatorNodeInfo string) DelegatorAccount {
	return DelegatorAccount{
		Address:           address,
		DelegatorId:       delegatorId,
		DelegatorNodeId:   delegatorNodeId,
		DelegatorNodeInfo: delegatorNodeInfo,
	}
}

type StakerAccount struct {
	Address        string
	StakerNodeId   string
	StakerNodeInfo string
}

func NewStakerAccount(address, stakerNodeId, stakerNodeInfo string) StakerAccount {
	return StakerAccount{
		Address:        address,
		StakerNodeId:   stakerNodeId,
		StakerNodeInfo: stakerNodeInfo,
	}
}
