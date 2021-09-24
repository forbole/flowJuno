package types

type StakeRequirements struct {
	Height      int64
	Role        uint8
	Requirement uint64
}

// Equal tells whether v and w represent the same rows
func (v StakeRequirements) Equal(w StakeRequirements) bool {
	return v.Height == w.Height &&
		v.Role == w.Role &&
		v.Requirement == w.Requirement
}

// StakeRequirements allows to build a new StakeRequirements
func NewStakeRequirements(
	height int64,
	role uint8,
	requirement uint64) StakeRequirements {
	return StakeRequirements{
		Height:      height,
		Role:        role,
		Requirement: requirement,
	}
}

type TotalStakeByType struct {
	Height     int64
	Role       int8
	TotalStake uint64
}

// Equal tells whether v and w represent the same rows
func (v TotalStakeByType) Equal(w TotalStakeByType) bool {
	return v.Height == w.Height &&
		v.Role == w.Role &&
		v.TotalStake == w.TotalStake
}

// TotalStake allows to build a new TotalStake
func NewTotalStakeByType(
	height int64,
	role int8,
	totalStake uint64) TotalStakeByType {
	return TotalStakeByType{
		Height:     height,
		Role:       role,
		TotalStake: totalStake,
	}
}

type WeeklyPayout struct {
	Height int64
	Payout uint64
}

// Equal tells whether v and w represent the same rows
func (v WeeklyPayout) Equal(w WeeklyPayout) bool {
	return v.Height == w.Height &&
		v.Payout == w.Payout
}

// WeeklyPayout allows to build a new WeeklyPayout
func NewWeeklyPayout(
	height int64,
	payout uint64) WeeklyPayout {
	return WeeklyPayout{
		Height: height,
		Payout: payout,
	}
}

type TotalStake struct {
	Height     int64
	TotalStake uint64
}

// Equal tells whether v and w represent the same rows
func (v TotalStake) Equal(w TotalStake) bool {
	return v.Height == w.Height &&
		v.TotalStake == w.TotalStake
}

// TotalStake allows to build a new TotalStake
func NewTotalStake(
	height int64,
	totalStake uint64) TotalStake {
	return TotalStake{
		Height:     height,
		TotalStake: totalStake,
	}
}

type StakingTable struct {
	Height       int64
	StakingTable []string
}

// StakingTable allows to build a new StakingTable
func NewStakingTable(
	height int64,
	stakingTable []string) StakingTable {
	return StakingTable{
		Height:       height,
		StakingTable: stakingTable,
	}
}
