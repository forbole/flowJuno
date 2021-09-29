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

type ProposedTable struct {
	Height        int64
	ProposedTable []string
}

// ProposedTable allows to build a new ProposedTable
func NewProposedTable(
	height int64,
	proposedTable []string) ProposedTable {
	return ProposedTable{
		Height:        height,
		ProposedTable: proposedTable,
	}
}

type CurrentTable struct {
	Height int64
	Table  []string
}

// CurrentTable allows to build a new CurrentTable
func NewCurrentTable(
	height int64,
	table []string) CurrentTable {
	return CurrentTable{
		Height: height,
		Table:  table,
	}
}

type NodeUnstakingTokens struct {
	NodeId         string
	TokenUnstaking uint64
	Height         int64
}

// Equal tells whether v and w represent the same rows
func (v NodeUnstakingTokens) Equal(w NodeUnstakingTokens) bool {
	return v.NodeId == w.NodeId &&
		v.TokenUnstaking == w.TokenUnstaking &&
		v.Height == w.Height
}

// NodeUnstakingTokens allows to build a new NodeUnstakingTokens
func NewNodeUnstakingTokens(
	nodeId string,
	tokenUnstaking uint64,
	height int64) NodeUnstakingTokens {
	return NodeUnstakingTokens{
		NodeId:         nodeId,
		TokenUnstaking: tokenUnstaking,
		Height:         height,
	}
}

type NodeTotalCommitment struct {
	NodeId          string
	TotalCommitment uint64
	Height          int64
}

// Equal tells whether v and w represent the same rows
func (v NodeTotalCommitment) Equal(w NodeTotalCommitment) bool {
	return v.NodeId == w.NodeId &&
		v.TotalCommitment == w.TotalCommitment &&
		v.Height == w.Height
}

// NodeTotalCommitment allows to build a new NodeTotalCommitment
func NewNodeTotalCommitment(
	nodeId string,
	totalCommitment uint64,
	height int64) NodeTotalCommitment {
	return NodeTotalCommitment{
		NodeId:          nodeId,
		TotalCommitment: totalCommitment,
		Height:          height,
	}
}

type NodeTotalCommitmentWithoutDelegators struct {
	NodeId                           string
	TotalCommitmentWithoutDelegators uint64
	Height                           int64
}

// Equal tells whether v and w represent the same rows
func (v NodeTotalCommitmentWithoutDelegators) Equal(w NodeTotalCommitmentWithoutDelegators) bool {
	return v.NodeId == w.NodeId &&
		v.TotalCommitmentWithoutDelegators == w.TotalCommitmentWithoutDelegators &&
		v.Height == w.Height
}

// NodeTotalCommitmentWithoutDelegators allows to build a new NodeTotalCommitmentWithoutDelegators
func NewNodeTotalCommitmentWithoutDelegators(
	nodeId string,
	totalCommitmentWithoutDelegators uint64,
	height int64) NodeTotalCommitmentWithoutDelegators {
	return NodeTotalCommitmentWithoutDelegators{
		NodeId:                           nodeId,
		TotalCommitmentWithoutDelegators: totalCommitmentWithoutDelegators,
		Height:                           height,
	}
}
