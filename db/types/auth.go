package types

// AccountRow represents a single row of the account table
type AccountRow struct {
	Address string `db:"address"`
}

// Equal tells whether v and w represent the same rows
func (v AccountRow) Equal(w AccountRow) bool {
	return v.Address == w.Address
}

// AccountRow allows to build a new AccountRow
func NewAccountRow(
	address string) AccountRow {
	return AccountRow{
		Address: address,
	}
}

// AccountBalanceRow represents a single row inside the account table
type AccountBalanceRow struct {
	Address     string  `db:"address"`
	Balance     float64 `db:"balance"`
	Code        string  `db:"code"`
	KeysList    string  `db:"keys_list"`
	ContractMap string  `db:"contract_map"`
}

// NewAccountBalanceRow allows to easily build a new AccountBalanceRow
func NewAccountBalanceRow(address string, balance float64, code, keysList, contractMap string) AccountBalanceRow {
	return AccountBalanceRow{
		Address:     address,
		Balance:     balance,
		Code:        code,
		KeysList:    keysList,
		ContractMap: contractMap,
	}
}

// Equal tells whether a and b contain the same data
func (a AccountBalanceRow) Equal(b AccountBalanceRow) bool {
	return (a.Address == b.Address &&
		a.Balance == b.Balance &&
		a.Code == b.Code &&
		a.ContractMap == b.ContractMap &&
		a.KeysList == b.KeysList)
}

// LockedAccountRow represents a single row of the locked_account table
type LockedAccountRow struct {
	Address       string `db:"address"`
	LockedAddress string `db:"locked_address"`
}

// Equal tells whether v and w represent the same rows
func (v LockedAccountRow) Equal(w LockedAccountRow) bool {
	return v.Address == w.Address &&
		v.LockedAddress == w.LockedAddress
}

// LockedAccountRow allows to build a new LockedAccountRow
func NewLockedAccountRow(
	address string,
	lockedAddress string) LockedAccountRow {
	return LockedAccountRow{
		Address:       address,
		LockedAddress: lockedAddress,
	}
}

type LockedAccountBalanceRow struct {
	LockedAddress string `db:"locked_address"`
	Balance       int    `db:"balance"`
	UnlockLimit   int    `db:"unlock_limit"`
	Height        int    `db:"height"`
}

func NewLockedAccountBalanceRow(lockedAddress string, balance, unlockLimit, height int) LockedAccountBalanceRow {
	return LockedAccountBalanceRow{
		LockedAddress: lockedAddress,
		Balance:       balance,
		UnlockLimit:   unlockLimit,
		Height:        height,
	}
}

func (a LockedAccountBalanceRow) Equal(b LockedAccountBalanceRow) bool {
	return (a.Balance == b.Balance &&
		a.LockedAddress == b.LockedAddress &&
		a.UnlockLimit == b.UnlockLimit &&
		a.Height == b.Height)
}

type DelegatorAccountRow struct {
	AccountAddress    string `db:"account_address"`
	DelegatorId       int64  `db:"delegator_id"`
	DelegatorNodeId   string `db:"delegator_node_id"`
	DelegatorNodeInfo string `db:"delegator_node_info"`
}

// Equal tells whether v and w represent the same rows
func (v DelegatorAccountRow) Equal(w DelegatorAccountRow) bool {
	return v.AccountAddress == w.AccountAddress &&
		v.DelegatorId == w.DelegatorId &&
		v.DelegatorNodeId == w.DelegatorNodeId &&
		v.DelegatorNodeInfo == w.DelegatorNodeInfo
}

// DelegatorAccountRow allows to build a new DelegatorAccountRow
func NewDelegatorAccountRow(
	accountAddress string,
	delegatorId int64,
	delegatorNodeId string,
	delegatorNodeInfo string) DelegatorAccountRow {
	return DelegatorAccountRow{
		AccountAddress:    accountAddress,
		DelegatorId:       delegatorId,
		DelegatorNodeId:   delegatorNodeId,
		DelegatorNodeInfo: delegatorNodeInfo,
	}
}

// StakerAccountRow represents a single row of the staker_account table
type StakerAccountRow struct {
	AccountAddress string `db:"account_address"`
	StakerNodeId   string `db:"staker_node_id"`
	StakerNodeInfo string `db:"staker_node_info"`
}

// Equal tells whether v and w represent the same rows
func (v StakerAccountRow) Equal(w StakerAccountRow) bool {
	return v.AccountAddress == w.AccountAddress &&
		v.StakerNodeId == w.StakerNodeId &&
		v.StakerNodeInfo == w.StakerNodeInfo
}

// StakerAccountRow allows to build a new StakerAccountRow
func NewStakerAccountRow(
	accountAddress string,
	stakerNodeId string,
	stakerNodeInfo string) StakerAccountRow {
	return StakerAccountRow{
		AccountAddress: accountAddress,
		StakerNodeId:   stakerNodeId,
		StakerNodeInfo: stakerNodeInfo,
	}
}
