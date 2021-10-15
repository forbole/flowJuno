package types

import (
	"fmt"
	"reflect"

	"github.com/onflow/cadence"
)

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
	DelegatorNodeInfo DelegatorNodeInfo
}

func NewDelegatorAccount(address string, delegatorId int64, delegatorNodeId string, delegatorNodeInfo DelegatorNodeInfo) DelegatorAccount {
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
	StakerNodeInfo StakerNodeInfo
}

func NewStakerAccount(address, stakerNodeId string, stakerNodeInfo StakerNodeInfo) StakerAccount {
	return StakerAccount{
		Address:        address,
		StakerNodeId:   stakerNodeId,
		StakerNodeInfo: stakerNodeInfo,
	}
}

type DelegatorNodeInfo struct {
	Id                       uint32
	NodeID                   string
	TokensCommitted          uint64
	TokensStaked             uint64
	TokensUnstaking          uint64
	TokensRewarded           uint64
	TokensUnstaked           uint64
	TokensRequestedToUnstake uint64
}

func NewDelegatorNodeInfo(id uint32,
	nodeID string,
	tokensCommitted uint64,
	tokensStaked uint64,
	tokensUnstaking uint64,
	tokensRewarded uint64,
	tokensUnstaked uint64,
	tokensRequestedToUnstake uint64) DelegatorNodeInfo {
	return DelegatorNodeInfo{
		Id:                       id,
		NodeID:                   nodeID,
		TokensCommitted:          tokensCommitted,
		TokensStaked:             tokensStaked,
		TokensUnstaking:          tokensUnstaking,
		TokensRewarded:           tokensRewarded,
		TokensUnstaked:           tokensUnstaked,
		TokensRequestedToUnstake: tokensRequestedToUnstake,
	}
}

func DelegatorNodeInfoArrayFromCadence(value cadence.Value) ([]DelegatorNodeInfo, error) {
	arrayValue, ok := value.(cadence.Array)
	if !ok {
		return nil, fmt.Errorf("This is not an array")
	}
	stakers := make([]DelegatorNodeInfo, len(arrayValue.Values))
	for i, value := range arrayValue.Values {
		stakervalue, err := DelegatorNodeInfoFromCadence(value)
		if err != nil {
			return nil, err
		}
		stakers[i] = stakervalue
	}
	return stakers, nil
}

func DelegatorNodeInfoFromCadence(value cadence.Value) (DelegatorNodeInfo, error) {
	fields := value.(cadence.Struct).Fields

	id, ok := fields[0].ToGoValue().(uint32)
	if !ok {
		return DelegatorNodeInfo{}, fmt.Errorf("id is not a uint32!!")
	}

	nodeID, ok := fields[1].ToGoValue().(string)
	if !ok {
		return DelegatorNodeInfo{}, fmt.Errorf("nodeID is not a string!!")
	}
	tokenCommited, ok := fields[2].ToGoValue().(uint64)
	if !ok {
		return DelegatorNodeInfo{}, fmt.Errorf("tokenCommited is not a uint64!!")
	}
	tokenStaked, ok := fields[3].ToGoValue().(uint64)
	if !ok {
		return DelegatorNodeInfo{}, fmt.Errorf("tokenStaked is not a uint64!!")
	}
	tokensUnstaking, ok := fields[4].ToGoValue().(uint64)
	if !ok {
		return DelegatorNodeInfo{}, fmt.Errorf("tokensUnstaking is not a uint64!!")
	}
	tokenRewarded, ok := fields[5].ToGoValue().(uint64)
	if !ok {
		return DelegatorNodeInfo{}, fmt.Errorf("tokenRewarded is not a uint64!!")
	}
	tokenUnstaked, ok := fields[6].ToGoValue().(uint64)
	if !ok {
		return DelegatorNodeInfo{}, fmt.Errorf("tokenUnstaked is not a uint64!!")
	}
	tokenRequestedToUnstake, ok := fields[7].ToGoValue().(uint64)
	if !ok {
		return DelegatorNodeInfo{}, fmt.Errorf("tokenRequestedToUnstake is not a uint64!!")
	}

	return NewDelegatorNodeInfo(id, nodeID, tokenCommited, tokenStaked, tokensUnstaking,
		tokenRewarded, tokenUnstaked, tokenRequestedToUnstake), nil

}

type StakerNodeInfo struct {
	Id                string
	Role              uint8
	NetworkingAddress string
	NetworkingKey     string
	StakingKey        string
	TokensStaked      uint64
	TokensCommitted   uint64
	TokensUnstaking   uint64
	TokensUnstaked    uint64
	TokensRewarded    uint64

	Delegators               []uint32
	DelegatorIDCounter       uint32
	TokensRequestedToUnstake uint64
	InitialWeight            uint64
}

func (w StakerNodeInfo) Equals(v StakerNodeInfo) bool {
	return (w.Id == v.Id &&
		w.Role == v.Role &&
		w.NetworkingAddress == v.NetworkingAddress &&
		w.NetworkingKey == v.NetworkingKey &&
		w.StakingKey == v.StakingKey &&
		w.TokensStaked == v.TokensStaked &&
		w.TokensCommitted == v.TokensCommitted &&
		w.TokensUnstaking == v.TokensUnstaking &&
		w.TokensUnstaked == v.TokensUnstaked &&
		w.TokensRewarded == v.TokensRewarded &&
		reflect.DeepEqual(w.Delegators, v.Delegators) &&
		w.DelegatorIDCounter == v.DelegatorIDCounter &&
		w.TokensRequestedToUnstake == v.TokensRequestedToUnstake &&
		w.InitialWeight == v.InitialWeight)
}

func NewStakerNodeInfoArrayFromCadence(value cadence.Value) ([]StakerNodeInfo, error) {
	arrayValue, ok := value.(cadence.Array)
	if !ok {
		return nil, fmt.Errorf("This is not an array")
	}
	stakers := make([]StakerNodeInfo, len(arrayValue.Values))
	for i, value := range arrayValue.Values {
		stakervalue, err := NewStakerNodeInfoFromCadence(value)
		if err != nil {
			return nil, err
		}
		stakers[i] = stakervalue
	}
	return stakers, nil
}

// NewNodeOperatorInfoFromInterface create a NodeOperatorInfo from []interface{}
func NewStakerNodeInfoFromCadence(value cadence.Value) (StakerNodeInfo, error) {
	fields := value.(cadence.Struct).Fields

	d, ok := fields[10].(cadence.Array)
	if !ok {
		fmt.Println("delegator not an cadence array")
	}
	delegators := make([]uint32, len(d.Values))
	for i, val := range d.Values {
		delegators[i], ok = val.ToGoValue().(uint32)
		if !ok {
			return StakerNodeInfo{}, fmt.Errorf("id is not type uint32")
		}
	}
	//fmt.Println(d)
	//delegators:=[]uint32{}

	id, ok := fields[0].ToGoValue().(string)
	if !ok {
		return StakerNodeInfo{}, fmt.Errorf("id is not type string")
	}
	role, ok := fields[1].ToGoValue().(uint8)
	if !ok {
		return StakerNodeInfo{}, fmt.Errorf("role is not type uint8")
	}
	networkingAddress, ok := fields[2].ToGoValue().(string)
	if !ok {
		return StakerNodeInfo{}, fmt.Errorf("networkingAddress is not string")
	}
	networkingKey, ok := fields[3].ToGoValue().(string)
	if !ok {
		return StakerNodeInfo{}, fmt.Errorf("networkingKey is not string")
	}
	stakingKey, ok := fields[4].ToGoValue().(string)
	if !ok {
		return StakerNodeInfo{}, fmt.Errorf("stakingKey is not string")
	}
	tokensStaked, ok := fields[5].ToGoValue().(uint64)
	if !ok {
		return StakerNodeInfo{}, fmt.Errorf("tokensStaked is not uint64")
	}
	tokensCommitted, ok := fields[6].ToGoValue().(uint64)
	if !ok {
		return StakerNodeInfo{}, fmt.Errorf("tokensCommitted is not uint64")
	}
	tokensUnstaking, ok := fields[7].ToGoValue().(uint64)
	if !ok {
		return StakerNodeInfo{}, fmt.Errorf("tokensUnstaking is not uint64")
	}
	tokensUnstaked, ok := fields[8].ToGoValue().(uint64)
	if !ok {
		return StakerNodeInfo{}, fmt.Errorf("tokensUnstaked is not uint64")
	}
	tokensRewarded, ok := fields[9].ToGoValue().(uint64)
	if !ok {
		return StakerNodeInfo{}, fmt.Errorf("tokensRewarded is not uint64")
	}

	delegatorIDCounter, ok := fields[11].ToGoValue().(uint32)
	if !ok {
		return StakerNodeInfo{}, fmt.Errorf("delegatorIDCounter is not uint32")
	}
	tokensRequestedToUnstake, ok := fields[12].ToGoValue().(uint64)
	if !ok {
		return StakerNodeInfo{}, fmt.Errorf("tokensRequestedToUnstake is not uint64")
	}
	initialWeight, ok := fields[13].ToGoValue().(uint64)
	if !ok {
		return StakerNodeInfo{}, fmt.Errorf("initialWeight is not uint64")
	}

	stakerNodeInfo := NewStakerNodeInfo(id, role, networkingAddress, networkingKey,
		stakingKey, tokensStaked, tokensCommitted, tokensUnstaking,
		tokensUnstaked, tokensRewarded, delegators, delegatorIDCounter,
		tokensRequestedToUnstake, initialWeight)

	return stakerNodeInfo, nil
}

func NewStakerNodeInfo(id string, role uint8, networkingAddress string, networkingKey string,
	stakingKey string, tokensStaked uint64, tokensCommitted uint64,
	tokensUnstaking uint64, tokensUnstaked uint64, tokensRewarded uint64,
	delegators []uint32, delegatorIDCounter uint32, tokensRequestedToUnstake uint64,
	initialWeight uint64) StakerNodeInfo {
	return StakerNodeInfo{
		Id:                id,
		Role:              role,
		NetworkingAddress: networkingAddress,
		NetworkingKey:     networkingKey,
		StakingKey:        stakingKey,
		TokensStaked:      tokensStaked,
		TokensCommitted:   tokensCommitted,
		TokensUnstaking:   tokensUnstaking,
		TokensUnstaked:    tokensUnstaked,
		TokensRewarded:    tokensRewarded,

		Delegators:               delegators,
		DelegatorIDCounter:       delegatorIDCounter,
		TokensRequestedToUnstake: tokensRequestedToUnstake,
		InitialWeight:            initialWeight,
	}
}
