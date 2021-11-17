package types

import (
	"reflect"

	"github.com/lib/pq"
)

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
	NodeId string `db:"node_id"`
}

// Equal tells whether v and w represent the same rows
func (v StakingTableRow) Equal(w StakingTableRow) bool {
	return v.NodeId == w.NodeId
}

// StakingTableRow allows to build a new StakingTableRow
func NewStakingTableRow(
	nodeId string) StakingTableRow {
	return StakingTableRow{
		NodeId: nodeId,
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

func NewNodeInfosFromTableRow(
	id string,
	role uint8,
	networkingAddress string,
	networkingKey string,
	stakingKey string,
	tokensStaked uint64,
	tokensCommitted uint64,
	tokensUnstaking uint64,
	tokensUnstaked uint64,
	tokensRewarded uint64,
	delegators pq.Int32Array,
	delegatorIDCounter uint32,
	tokensRequestedToUnstake uint64,
	initialWeight uint64,
	height uint64) NodeInfosFromTableRow {
	return NodeInfosFromTableRow{
		Id:                       id,
		Role:                     role,
		NetworkingAddress:        networkingAddress,
		NetworkingKey:            networkingKey,
		StakingKey:               stakingKey,
		TokensStaked:             tokensStaked,
		TokensCommitted:          tokensCommitted,
		TokensUnstaking:          tokensUnstaking,
		TokensUnstaked:           tokensUnstaked,
		TokensRewarded:           tokensRewarded,
		Delegators:               delegators,
		DelegatorIDCounter:       delegatorIDCounter,
		TokensRequestedToUnstake: tokensRequestedToUnstake,
		InitialWeight:            initialWeight,
		Height:                   height,
	}
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

// NodeInfosFromTableRow represents a single row of the node_infos_from_table table
type NodeInfosFromTableRow struct {
	Id                       string        `db:"id"`
	Role                     uint8         `db:"role"`
	NetworkingAddress        string        `db:"networking_address"`
	NetworkingKey            string        `db:"networking_key"`
	StakingKey               string        `db:"staking_key"`
	TokensStaked             uint64        `db:"tokens_staked"`
	TokensCommitted          uint64        `db:"tokens_committed"`
	TokensUnstaking          uint64        `db:"tokens_unstaking"`
	TokensUnstaked           uint64        `db:"tokens_unstaked"`
	TokensRewarded           uint64        `db:"tokens_rewarded"`
	Delegators               pq.Int32Array `db:"delegators"`
	DelegatorIDCounter       uint32        `db:"delegator_i_d_counter"`
	TokensRequestedToUnstake uint64        `db:"tokens_requested_to_unstake"`
	InitialWeight            uint64        `db:"initial_weight"`
	Height                   uint64        `db:"height"`
}

// Equal tells whether v and w represent the same rows
func (v NodeInfosFromTableRow) Equal(w NodeInfosFromTableRow) bool {
	return v.Id == w.Id &&
		v.Role == w.Role &&
		v.NetworkingAddress == w.NetworkingAddress &&
		v.NetworkingKey == w.NetworkingKey &&
		v.StakingKey == w.StakingKey &&
		v.TokensStaked == w.TokensStaked &&
		v.TokensCommitted == w.TokensCommitted &&
		v.TokensUnstaking == w.TokensUnstaking &&
		v.TokensUnstaked == w.TokensUnstaked &&
		v.TokensRewarded == w.TokensRewarded &&
		reflect.DeepEqual(v.Delegators, w.Delegators) &&
		v.DelegatorIDCounter == w.DelegatorIDCounter &&
		v.TokensRequestedToUnstake == w.TokensRequestedToUnstake &&
		v.InitialWeight == w.InitialWeight &&
		v.Height == w.Height
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
	CutPercentage uint64 `db:"cut_percentage"`
	Height        int64  `db:"height"`
}

// Equal tells whether v and w represent the same rows
func (v CutPercentageRow) Equal(w CutPercentageRow) bool {
	return v.CutPercentage == w.CutPercentage &&
		v.Height == w.Height
}

// CutPercentageRow allows to build a new CutPercentageRow
func NewCutPercentageRow(
	cutPercentage uint64,
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
	DelegatorID uint32 `db:"delegator_id"`
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

// DelegatorInfoRow represents a single row of the delegator_info table
type DelegatorInfoRow struct {
	Id                       uint32 `db:"id"`
	NodeId                   string `db:"node_id"`
	TokensCommitted          uint64 `db:"tokens_committed"`
	TokensStaked             uint64 `db:"tokens_staked"`
	TokensUnstaking          uint64 `db:"tokens_unstaking"`
	TokensRewarded           uint64 `db:"tokens_rewarded"`
	TokensUnstaked           uint64 `db:"tokens_unstaked"`
	TokensRequestedToUnstake uint64 `db:"tokens_requested_to_unstake"`
	Height                   uint64 `db:"height"`
}

// Equal tells whether v and w represent the same rows
func (v DelegatorInfoRow) Equal(w DelegatorInfoRow) bool {
	return v.Id == w.Id &&
		v.NodeId == w.NodeId &&
		v.TokensCommitted == w.TokensCommitted &&
		v.TokensStaked == w.TokensStaked &&
		v.TokensUnstaking == w.TokensUnstaking &&
		v.TokensRewarded == w.TokensRewarded &&
		v.TokensUnstaked == w.TokensUnstaked &&
		v.TokensRequestedToUnstake == w.TokensRequestedToUnstake &&
		v.Height == w.Height
}

// DelegatorInfoRow allows to build a new DelegatorInfoRow
func NewDelegatorInfoRow(
	id uint32,
	nodeId string,
	tokensCommitted uint64,
	tokensStaked uint64,
	tokensUnstaking uint64,
	tokensRewarded uint64,
	tokensUnstaked uint64,
	tokensRequestedToUnstake uint64,
	height uint64) DelegatorInfoRow {
	return DelegatorInfoRow{
		Id:                       id,
		NodeId:                   nodeId,
		TokensCommitted:          tokensCommitted,
		TokensStaked:             tokensStaked,
		TokensUnstaking:          tokensUnstaking,
		TokensRewarded:           tokensRewarded,
		TokensUnstaked:           tokensUnstaked,
		TokensRequestedToUnstake: tokensRequestedToUnstake,
		Height:                   height,
	}
}

// DelegatorInfoFromAddressRow represents a single row of the delegator_info_from_address table
type DelegatorInfoFromAddressRow struct {
	DelegatorInfo string `db:"delegator_info"`
	Height        int64  `db:"height"`
	Address       string `db:"address"`
}

// Equal tells whether v and w represent the same rows
func (v DelegatorInfoFromAddressRow) Equal(w DelegatorInfoFromAddressRow) bool {
	return v.DelegatorInfo == w.DelegatorInfo &&
		v.Height == w.Height &&
		v.Address == w.Address
}

// DelegatorInfoFromAddressRow allows to build a new DelegatorInfoFromAddressRow
func NewDelegatorInfoFromAddressRow(
	delegatorInfo string,
	height int64,
	address string) DelegatorInfoFromAddressRow {
	return DelegatorInfoFromAddressRow{
		DelegatorInfo: delegatorInfo,
		Height:        height,
		Address:       address,
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
