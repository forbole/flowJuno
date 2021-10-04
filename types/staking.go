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

type NodeStakingKey struct {
	NodeId         string
	NodeStakingKey string
	Height         int64
}

// Equal tells whether v and w represent the same rows
func (v NodeStakingKey) Equal(w NodeStakingKey) bool {
	return v.NodeId == w.NodeId &&
		v.NodeStakingKey == w.NodeStakingKey &&
		v.Height == w.Height
}

// NodeStakingKey allows to build a new NodeStakingKey
func NewNodeStakingKey(
	nodeId string,
	nodeStakingKey string,
	height int64) NodeStakingKey {
	return NodeStakingKey{
		NodeId:         nodeId,
		NodeStakingKey: nodeStakingKey,
		Height:         height,
	}
}

type NodeStakedTokens struct {
	NodeId           string
	NodeStakedTokens uint64
	Height           int64
}

// Equal tells whether v and w represent the same rows
func (v NodeStakedTokens) Equal(w NodeStakedTokens) bool {
	return v.NodeId == w.NodeId &&
		v.NodeStakedTokens == w.NodeStakedTokens &&
		v.Height == w.Height
}

// NodeStakedTokens allows to build a new NodeStakedTokens
func NewNodeStakedTokens(
	nodeId string,
	nodeStakedTokens uint64,
	height int64) NodeStakedTokens {
	return NodeStakedTokens{
		NodeId:           nodeId,
		NodeStakedTokens: nodeStakedTokens,
		Height:           height,
	}
}

type NodeRole struct {
	NodeId string
	Role   uint8
	Height int64
}

// Equal tells whether v and w represent the same rows
func (v NodeRole) Equal(w NodeRole) bool {
	return v.NodeId == w.NodeId &&
		v.Role == w.Role &&
		v.Height == w.Height
}

// NodeRole allows to build a new NodeRole
func NewNodeRole(
	nodeId string,
	role uint8,
	height int64) NodeRole {
	return NodeRole{
		NodeId: nodeId,
		Role:   role,
		Height: height,
	}
}

type NodeRewardedTokens struct {
	NodeId             string
	NodeRewardedTokens uint64
	Height             int64
}

// Equal tells whether v and w represent the same rows
func (v NodeRewardedTokens) Equal(w NodeRewardedTokens) bool {
	return v.NodeId == w.NodeId &&
		v.NodeRewardedTokens == w.NodeRewardedTokens &&
		v.Height == w.Height
}

// NodeRewardedTokens allows to build a new NodeRewardedTokens
func NewNodeRewardedTokens(
	nodeId string,
	nodeRewardedTokens uint64,
	height int64) NodeRewardedTokens {
	return NodeRewardedTokens{
		NodeId:             nodeId,
		NodeRewardedTokens: nodeRewardedTokens,
		Height:             height,
	}
}

type NodeNetworkingKey struct {
	NodeId        string
	NetworkingKey string
	Height        int64
}

// Equal tells whether v and w represent the same rows
func (v NodeNetworkingKey) Equal(w NodeNetworkingKey) bool {
	return v.NodeId == w.NodeId &&
		v.NetworkingKey == w.NetworkingKey &&
		v.Height == w.Height
}

// NodeNetworkingKey allows to build a new NodeNetworkingKey
func NewNodeNetworkingKey(
	nodeId string,
	networkingKey string,
	height int64) NodeNetworkingKey {
	return NodeNetworkingKey{
		NodeId:        nodeId,
		NetworkingKey: networkingKey,
		Height:        height,
	}
}

type NodeNetworkingAddress struct {
	NodeId            string
	NetworkingAddress string
	Height            int64
}

// Equal tells whether v and w represent the same rows
func (v NodeNetworkingAddress) Equal(w NodeNetworkingAddress) bool {
	return v.NodeId == w.NodeId &&
		v.NetworkingAddress == w.NetworkingAddress &&
		v.Height == w.Height
}

// NodeNetworkingAddress allows to build a new NodeNetworkingAddress
func NewNodeNetworkingAddress(
	nodeId string,
	networkingAddress string,
	height int64) NodeNetworkingAddress {
	return NodeNetworkingAddress{
		NodeId:            nodeId,
		NetworkingAddress: networkingAddress,
		Height:            height,
	}
}

type NodeInitialWeight struct {
	NodeId        string
	InitialWeight uint64
	Height        int64
}

// Equal tells whether v and w represent the same rows
func (v NodeInitialWeight) Equal(w NodeInitialWeight) bool {
	return v.NodeId == w.NodeId &&
		v.InitialWeight == w.InitialWeight &&
		v.Height == w.Height
}

// NodeInitialWeight allows to build a new NodeInitialWeight
func NewNodeInitialWeight(
	nodeId string,
	initialWeight uint64,
	height int64) NodeInitialWeight {
	return NodeInitialWeight{
		NodeId:        nodeId,
		InitialWeight: initialWeight,
		Height:        height,
	}
}

type NodeInfoFromAddress struct {
	Address  string
	NodeInfo StakerNodeInfo
	Height   int64
}

// NodeInfoFromNodeIDs allows to build a new NodeInfoFromNodeIDs
func NewNodeInfoFromAddress(
	address string,
	nodeInfo StakerNodeInfo,
	height int64) NodeInfoFromAddress {
	return NodeInfoFromAddress{
		Address:  address,
		NodeInfo: nodeInfo,
		Height:   height,
	}
}

type NodeInfoFromNodeID struct {
	NodeId   string
	NodeInfo StakerNodeInfo
	Height   int64
}

// NodeInfoFromNodeID allows to build a new NodeInfoFromNodeID
func NewNodeInfoFromNodeID(
	nodeId string,
	nodeInfo StakerNodeInfo,
	height int64) NodeInfoFromNodeID {
	return NodeInfoFromNodeID{
		NodeId:   nodeId,
		NodeInfo: nodeInfo,
		Height:   height,
	}
}

type NodeCommittedTokens struct {
	NodeId          string
	CommittedTokens uint64
	Height          int64
}

// NodeCommittedTokens allows to build a new NodeCommittedTokens
func NewNodeCommittedTokens(
	nodeId string,
	committedTokens uint64,
	height int64) NodeCommittedTokens {
	return NodeCommittedTokens{
		NodeId:          nodeId,
		CommittedTokens: committedTokens,
		Height:          height,
	}
}
