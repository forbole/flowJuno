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

// NodeStakedTokensRow represents a single row of the node_staked_tokens table
type NodeStakedTokensRow struct {
	NodeId           string `db:"node_id"`
	NodeStakedTokens uint64 `db:"node_staked_tokens"`
	Height           int64  `db:"height"`
}

// Equal tells whether v and w represent the same rows
func (v NodeStakedTokensRow) Equal(w NodeStakedTokensRow) bool {
	return v.NodeId == w.NodeId &&
		v.NodeStakedTokens == w.NodeStakedTokens &&
		v.Height == w.Height
}

// NodeStakedTokensRow allows to build a new NodeStakedTokensRow
func NewNodeStakedTokensRow(
	nodeId string,
	nodeStakedTokens uint64,
	height int64) NodeStakedTokensRow {
	return NodeStakedTokensRow{
		NodeId:           nodeId,
		NodeStakedTokens: nodeStakedTokens,
		Height:           height,
	}
}

// NodeRoleRow represents a single row of the node_role table
type NodeRoleRow struct {
	NodeId string `db:"node_id"`
	Role   uint8  `db:"role"`
	Height int64  `db:"height"`
}

// Equal tells whether v and w represent the same rows
func (v NodeRoleRow) Equal(w NodeRoleRow) bool {
	return v.NodeId == w.NodeId &&
		v.Role == w.Role &&
		v.Height == w.Height
}

// NodeRoleRow allows to build a new NodeRoleRow
func NewNodeRoleRow(
	nodeId string,
	role uint8,
	height int64) NodeRoleRow {
	return NodeRoleRow{
		NodeId: nodeId,
		Role:   role,
		Height: height,
	}
}

// NodeRewardedTokensRow represents a single row of the node_rewarded_tokens table
type NodeRewardedTokensRow struct {
	NodeId             string `db:"node_id"`
	NodeRewardedTokens uint64 `db:"node_rewarded_tokens"`
	Height             int64  `db:"height"`
}

// Equal tells whether v and w represent the same rows
func (v NodeRewardedTokensRow) Equal(w NodeRewardedTokensRow) bool {
	return v.NodeId == w.NodeId &&
		v.NodeRewardedTokens == w.NodeRewardedTokens &&
		v.Height == w.Height
}

// NodeRewardedTokensRow allows to build a new NodeRewardedTokensRow
func NewNodeRewardedTokensRow(
	nodeId string,
	nodeRewardedTokens uint64,
	height int64) NodeRewardedTokensRow {
	return NodeRewardedTokensRow{
		NodeId:             nodeId,
		NodeRewardedTokens: nodeRewardedTokens,
		Height:             height,
	}
}

// NodeNetworkingKeyRow represents a single row of the node_networking_key table
type NodeNetworkingKeyRow struct {
	NodeId        string `db:"node_id"`
	NetworkingKey string `db:"networking_key"`
	Height        int64  `db:"height"`
}

// Equal tells whether v and w represent the same rows
func (v NodeNetworkingKeyRow) Equal(w NodeNetworkingKeyRow) bool {
	return v.NodeId == w.NodeId &&
		v.NetworkingKey == w.NetworkingKey &&
		v.Height == w.Height
}

// NodeNetworkingKeyRow allows to build a new NodeNetworkingKeyRow
func NewNodeNetworkingKeyRow(
	nodeId string,
	networkingKey string,
	height int64) NodeNetworkingKeyRow {
	return NodeNetworkingKeyRow{
		NodeId:        nodeId,
		NetworkingKey: networkingKey,
		Height:        height,
	}
}

// NodeNetworkingAddressRow represents a single row of the node_networking_address table
type NodeNetworkingAddressRow struct {
	NodeId            string `db:"node_id"`
	NetworkingAddress string `db:"networking_address"`
	Height            int64  `db:"height"`
}

// Equal tells whether v and w represent the same rows
func (v NodeNetworkingAddressRow) Equal(w NodeNetworkingAddressRow) bool {
	return v.NodeId == w.NodeId &&
		v.NetworkingAddress == w.NetworkingAddress &&
		v.Height == w.Height
}

// NodeNetworkingAddressRow allows to build a new NodeNetworkingAddressRow
func NewNodeNetworkingAddressRow(
	nodeId string,
	networkingAddress string,
	height int64) NodeNetworkingAddressRow {
	return NodeNetworkingAddressRow{
		NodeId:            nodeId,
		NetworkingAddress: networkingAddress,
		Height:            height,
	}
}

// NodeInitialWeightRow represents a single row of the node_initial_weight table
type NodeInitialWeightRow struct {
	NodeId        string `db:"node_id"`
	InitialWeight uint64 `db:"initial_weight"`
	Height        int64  `db:"height"`
}

// Equal tells whether v and w represent the same rows
func (v NodeInitialWeightRow) Equal(w NodeInitialWeightRow) bool {
	return v.NodeId == w.NodeId &&
		v.InitialWeight == w.InitialWeight &&
		v.Height == w.Height
}

// NodeInitialWeightRow allows to build a new NodeInitialWeightRow
func NewNodeInitialWeightRow(
	nodeId string,
	initialWeight uint64,
	height int64) NodeInitialWeightRow {
	return NodeInitialWeightRow{
		NodeId:        nodeId,
		InitialWeight: initialWeight,
		Height:        height,
	}
}

// NodeInfoFromNodeIDsRow represents a single row of the node_info_from_address table
type NodeInfoFromAddressRow struct {
	Address  string `db:"address"`
	NodeInfo string `db:"node_info"`
	Height   int64  `db:"height"`
}

// Equal tells whether v and w represent the same rows
func (v NodeInfoFromAddressRow) Equal(w NodeInfoFromAddressRow) bool {
	return v.Address == w.Address &&
		v.NodeInfo == w.NodeInfo &&
		v.Height == w.Height
}

// NodeInfoFromNodeIDsRow allows to build a new NodeInfoFromNodeIDsRow
func NewNodeInfoFromAddressRow(
	address string,
	nodeInfo string,
	height int64) NodeInfoFromAddressRow {
	return NodeInfoFromAddressRow{
		Address:  address,
		NodeInfo: nodeInfo,
		Height:   height,
	}
}

// NodeInfoFromNodeIDRow represents a single row of the node_info_from_node_i_d table
type NodeInfoFromNodeIDRow struct {
	NodeId   string `db:"node_id"`
	NodeInfo string `db:"node_info"`
	Height   int64  `db:"height"`
}

// Equal tells whether v and w represent the same rows
func (v NodeInfoFromNodeIDRow) Equal(w NodeInfoFromNodeIDRow) bool {
	return v.NodeId == w.NodeId &&
		v.NodeInfo == w.NodeInfo &&
		v.Height == w.Height
}

// NodeInfoFromNodeIDRow allows to build a new NodeInfoFromNodeIDRow
func NewNodeInfoFromNodeIDRow(
	nodeId string,
	nodeInfo string,
	height int64) NodeInfoFromNodeIDRow {
	return NodeInfoFromNodeIDRow{
		NodeId:   nodeId,
		NodeInfo: nodeInfo,
		Height:   height,
	}
}

// NodeCommittedTokensRow represents a single row of the node_committed_tokens table
type NodeCommittedTokensRow struct {
	NodeId          string `db:"node_id"`
	CommittedTokens uint64 `db:"committed_tokens"`
	Height          int64  `db:"height"`
}

// Equal tells whether v and w represent the same rows
func (v NodeCommittedTokensRow) Equal(w NodeCommittedTokensRow) bool {
	return v.NodeId == w.NodeId &&
		v.CommittedTokens == w.CommittedTokens &&
		v.Height == w.Height
}

// NodeCommittedTokensRow allows to build a new NodeCommittedTokensRow
func NewNodeCommittedTokensRow(
	nodeId string,
	committedTokens uint64,
	height int64) NodeCommittedTokensRow {
	return NodeCommittedTokensRow{
		NodeId:          nodeId,
		CommittedTokens: committedTokens,
		Height:          height,
	}
}

// CutPercentageRow represents a single row of the cut_percentage table
type CutPercentageRow struct {
	CutPercentage string `db:"cut_percentage"`
	Height        int64  `db:"height"`
}

// Equal tells whether v and w represent the same rows
func (v CutPercentageRow) Equal(w CutPercentageRow) bool {
	return v.CutPercentage == w.CutPercentage &&
		v.Height == w.Height
}

// CutPercentageRow allows to build a new CutPercentageRow
func NewCutPercentageRow(
	cutPercentage string,
	height int64) CutPercentageRow {
	return CutPercentageRow{
		CutPercentage: cutPercentage,
		Height:        height,
	}
}

// DelegatorCommittedRow represents a single row of the delegator_committed table
type DelegatorCommittedRow struct {
	Committed   uint64 `db:"committed"`
	Height      uint64 `db:"height"`
	NodeId      string `db:"node_id"`
	DelegatorID uint32 `db:"delegator_i_d"`
}

// Equal tells whether v and w represent the same rows
func (v DelegatorCommittedRow) Equal(w DelegatorCommittedRow) bool {
	return v.Committed == w.Committed &&
		v.Height == w.Height &&
		v.NodeId == w.NodeId &&
		v.DelegatorID == w.DelegatorID
}

// DelegatorCommittedRow allows to build a new DelegatorCommittedRow
func NewDelegatorCommittedRow(
	committed uint64,
	height uint64,
	nodeId string,
	delegatorID uint32) DelegatorCommittedRow {
	return DelegatorCommittedRow{
		Committed:   committed,
		Height:      height,
		NodeId:      nodeId,
		DelegatorID: delegatorID,
	}
}
