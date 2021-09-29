package types

// TotalStakeRow represents a single row of the total_stake table
type TotalStakeByTypeRow struct {
	Height     int64  `db:"height"`
	Role       int8   `db:"role"`
	TotalStake uint64 `db:"total_stake"`
}

// Equal tells whether v and w represent the same rows
func (v TotalStakeByTypeRow) Equal(w TotalStakeByTypeRow) bool {
	return v.Height == w.Height &&
		v.Role == w.Role &&
		v.TotalStake == w.TotalStake
}

// TotalStakeRow allows to build a new TotalStakeRow
func NewTotalStakeByTypeRow(
	height int64,
	role int8,
	totalStake uint64) TotalStakeByTypeRow {
	return TotalStakeByTypeRow{
		Height:     height,
		Role:       role,
		TotalStake: totalStake,
	}
}

// StakeRequirementsRow represents a single row of the stake_requirements table
type StakeRequirementsRow struct {
	Height       int64  `db:"height"`
	Role         int8   `db:"role"`
	Requirements uint64 `db:"requirements"`
}

// Equal tells whether v and w represent the same rows
func (v StakeRequirementsRow) Equal(w StakeRequirementsRow) bool {
	return v.Height == w.Height &&
		v.Role == w.Role &&
		v.Requirements == w.Requirements
}

// StakeRequirementsRow allows to build a new StakeRequirementsRow
func NewStakeRequirementsRow(
	height int64,
	role int8,
	requirements uint64) StakeRequirementsRow {
	return StakeRequirementsRow{
		Height:       height,
		Role:         role,
		Requirements: requirements,
	}
}

// WeeklyPayoutRow represents a single row of the weekly_payout table
type WeeklyPayoutRow struct {
	Height int64  `db:"height"`
	Payout uint64 `db:"payout"`
}

// Equal tells whether v and w represent the same rows
func (v WeeklyPayoutRow) Equal(w WeeklyPayoutRow) bool {
	return v.Height == w.Height &&
		v.Payout == w.Payout
}

// WeeklyPayoutRow allows to build a new WeeklyPayoutRow
func NewWeeklyPayoutRow(
	height int64,
	payout uint64) WeeklyPayoutRow {
	return WeeklyPayoutRow{
		Height: height,
		Payout: payout,
	}
}

// TotalStakeRow represents a single row of the total_stake table
type TotalStakeRow struct {
	Height     int64  `db:"height"`
	TotalStake uint64 `db:"total_stake"`
}

// Equal tells whether v and w represent the same rows
func (v TotalStakeRow) Equal(w TotalStakeRow) bool {
	return v.Height == w.Height &&
		v.TotalStake == w.TotalStake
}

// TotalStakeRow allows to build a new TotalStakeRow
func NewTotalStakeRow(
	height int64,
	totalStake uint64) TotalStakeRow {
	return TotalStakeRow{
		Height:     height,
		TotalStake: totalStake,
	}
}

// StakingTableRow represents a single row of the staking_table table
type StakingTableRow struct {
	Height       int64  `db:"height"`
	StakingTable string `db:"staking_table"`
}

// Equal tells whether v and w represent the same rows
func (v StakingTableRow) Equal(w StakingTableRow) bool {
	return v.Height == w.Height &&
		v.StakingTable == w.StakingTable
}

// StakingTableRow allows to build a new StakingTableRow
func NewStakingTableRow(
	height int64,
	stakingTable string) StakingTableRow {
	return StakingTableRow{
		Height:       height,
		StakingTable: stakingTable,
	}
}

// ProposedTableRow represents a single row of the proposed_table table
type ProposedTableRow struct {
	Height        int64  `db:"height"`
	ProposedTable string `db:"proposed_table"`
}

// Equal tells whether v and w represent the same rows
func (v ProposedTableRow) Equal(w ProposedTableRow) bool {
	return v.Height == w.Height &&
		v.ProposedTable == w.ProposedTable
}

// ProposedTableRow allows to build a new ProposedTableRow
func NewProposedTableRow(
	height int64,
	proposedTable string) ProposedTableRow {
	return ProposedTableRow{
		Height:        height,
		ProposedTable: proposedTable,
	}
}

// CurrentTableRow represents a single row of the current_table table
type CurrentTableRow struct {
	Height int64  `db:"height"`
	Table  string `db:"current_table"`
}

// Equal tells whether v and w represent the same rows
func (v CurrentTableRow) Equal(w CurrentTableRow) bool {
	return v.Height == w.Height &&
		v.Table == w.Table
}

// CurrentTableRow allows to build a new CurrentTableRow
func NewCurrentTableRow(
	height int64,
	table string) CurrentTableRow {
	return CurrentTableRow{
		Height: height,
		Table:  table,
	}
}

// NodeUnstakingTokensRow represents a single row of the node_unstaking_tokens table
type NodeUnstakingTokensRow struct {
	NodeId         string `db:"node_id"`
	TokenUnstaking uint64 `db:"token_unstaking"`
	Height         int64  `db:"height"`
}

// Equal tells whether v and w represent the same rows
func (v NodeUnstakingTokensRow) Equal(w NodeUnstakingTokensRow) bool {
	return v.NodeId == w.NodeId &&
		v.TokenUnstaking == w.TokenUnstaking &&
		v.Height == w.Height
}

// NodeUnstakingTokensRow allows to build a new NodeUnstakingTokensRow
func NewNodeUnstakingTokensRow(
	nodeId string,
	tokenUnstaking uint64,
	height int64) NodeUnstakingTokensRow {
	return NodeUnstakingTokensRow{
		NodeId:         nodeId,
		TokenUnstaking: tokenUnstaking,
		Height:         height,
	}
}

// NodeTotalCommitmentRow represents a single row of the node_total_commitment table
type NodeTotalCommitmentRow struct {
	NodeId          string `db:"node_id"`
	TotalCommitment uint64 `db:"total_commitment"`
	Height          int64  `db:"height"`
}

// Equal tells whether v and w represent the same rows
func (v NodeTotalCommitmentRow) Equal(w NodeTotalCommitmentRow) bool {
	return v.NodeId == w.NodeId &&
		v.TotalCommitment == w.TotalCommitment &&
		v.Height == w.Height
}

// NodeTotalCommitmentRow allows to build a new NodeTotalCommitmentRow
func NewNodeTotalCommitmentRow(
	nodeId string,
	totalCommitment uint64,
	height int64) NodeTotalCommitmentRow {
	return NodeTotalCommitmentRow{
		NodeId:          nodeId,
		TotalCommitment: totalCommitment,
		Height:          height,
	}
}

// NodeTotalCommitmentWithoutDelegatorsRow represents a single row of the node_total_commitment_without_delegators table
type NodeTotalCommitmentWithoutDelegatorsRow struct {
	NodeId                           string `db:"node_id"`
	TotalCommitmentWithoutDelegators uint64 `db:"total_commitment_without_delegators"`
	Height                           int64  `db:"height"`
}

// Equal tells whether v and w represent the same rows
func (v NodeTotalCommitmentWithoutDelegatorsRow) Equal(w NodeTotalCommitmentWithoutDelegatorsRow) bool {
	return v.NodeId == w.NodeId &&
		v.TotalCommitmentWithoutDelegators == w.TotalCommitmentWithoutDelegators &&
		v.Height == w.Height
}

// NodeTotalCommitmentWithoutDelegatorsRow allows to build a new NodeTotalCommitmentWithoutDelegatorsRow
func NewNodeTotalCommitmentWithoutDelegatorsRow(
	nodeId string,
	totalCommitmentWithoutDelegators uint64,
	height int64) NodeTotalCommitmentWithoutDelegatorsRow {
	return NodeTotalCommitmentWithoutDelegatorsRow{
		NodeId:                           nodeId,
		TotalCommitmentWithoutDelegators: totalCommitmentWithoutDelegators,
		Height:                           height,
	}
}

// NodeStakingKeyRow represents a single row of the node_staking_key table
type NodeStakingKeyRow struct {
	NodeId         string `db:"node_id"`
	NodeStakingKey string `db:"node_staking_key"`
	Height         int64  `db:"height"`
}

// Equal tells whether v and w represent the same rows
func (v NodeStakingKeyRow) Equal(w NodeStakingKeyRow) bool {
	return v.NodeId == w.NodeId &&
		v.NodeStakingKey == w.NodeStakingKey &&
		v.Height == w.Height
}

// NodeStakingKeyRow allows to build a new NodeStakingKeyRow
func NewNodeStakingKeyRow(
	nodeId string,
	nodeStakingKey string,
	height int64) NodeStakingKeyRow {
	return NodeStakingKeyRow{
		NodeId:         nodeId,
		NodeStakingKey: nodeStakingKey,
		Height:         height,
	}
}
