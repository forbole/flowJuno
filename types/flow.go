package types

import (
	"fmt"
	"reflect"

	"github.com/onflow/flow-go-sdk"
)

// Tx represents an already existing blockchain transaction
type Tx struct {
	Events []flow.Event
	Status string
}

// NewTx allows to create a new Tx instance from the given txResponse
func NewTx(tx flow.TransactionResult) Tx {
	return Tx{
		Events: tx.Events,
		Status: tx.Status.String(),
	}
}

// FindEventByType searches inside the given tx events for the message having the specified index, in order
// to find the event having the given type, and returns it.
// If no such event is found, returns an error instead.
func (tx Tx) FindEventByType(index int, eventType string) (flow.Event, error) {
	for _, ev := range tx.Events {
		if ev.Type == eventType {
			return ev, nil
		}
	}

	return flow.Event{}, fmt.Errorf("no %s event found inside tx", eventType, )
}

// Successful tells whether this tx is successful or not
func (tx Tx) Successful() bool {
	return tx.Status == flow.TransactionStatusSealed.String()
}

type NodeOperators struct {
	Height   int64
	NodeInfos []*NodeInfo
}

func NewNodeOperators(height int64,nodeInfos []*NodeInfo)NodeOperators{
	return NodeOperators{
		Height: height,
		NodeInfos: nodeInfos,
	}
}

type NodeInfo struct {
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

// NewNodeOperatorInfoFromInterface create a NodeOperatorInfo from []interface{}
// **Unsafe** Make sure it is FlowIDTableStaking.NodeInfo type from Candance
func NewNodeInfoFromCandance(node interface{}) (NodeInfo, error) {
	s := reflect.ValueOf(node)
	//Delegator:=s.Index(10).Interface()
	//delegations:=reflect.ValueOf(s.Index(10))

	d, ok := s.Index(10).Interface().([]interface{})
	if !ok {
		return NodeInfo{}, fmt.Errorf("delegators array cannot cast to []interface{}")
	}
	delegators := make([]uint32, len(d))
	for i, a := range d {
		println(a.(uint32))
		delegators[i], ok = a.(uint32)
		if !ok {
			return NodeInfo{}, fmt.Errorf("delegators are not uint32")
		}
	}
	id, ok := s.Index(0).Interface().(string)
	if !ok {
		return NodeInfo{},fmt.Errorf("id is not type string")
	}
	role, ok := s.Index(1).Interface().(uint8)
	if !ok {
		return NodeInfo{},fmt.Errorf("role is not type uint8")
	}
	networkingAddress, ok := s.Index(2).Interface().(string)
	if !ok {
		return NodeInfo{},fmt.Errorf("networkingAddress is not string")
	}
	networkingKey, ok := s.Index(3).Interface().(string)
	if !ok {
		return NodeInfo{},fmt.Errorf("networkingKey is not string")
	}
	stakingKey, ok := s.Index(4).Interface().(string)
	if !ok {
		return NodeInfo{},fmt.Errorf("stakingKey is not string")
	}
	tokensStaked, ok := s.Index(5).Interface().(uint64)
	if !ok {
		return NodeInfo{},fmt.Errorf("tokensStaked is not uint64")
	}
	tokensCommitted, ok := s.Index(6).Interface().(uint64)
	if !ok {
		return NodeInfo{},fmt.Errorf("tokensCommitted is not uint64")
	}
	tokensUnstaking, ok := s.Index(7).Interface().(uint64)
	if !ok {
		return NodeInfo{},fmt.Errorf("tokensUnstaking is not uint64")
	}
	tokensUnstaked, ok := s.Index(8).Interface().(uint64)
	if !ok {
		return NodeInfo{},fmt.Errorf("tokensUnstaked is not uint64")
	}
	tokensRewarded, ok := s.Index(9).Interface().(uint64)
	if !ok {
		return NodeInfo{},fmt.Errorf("tokensRewarded is not uint64")
	}

	delegatorIDCounter, ok := s.Index(11).Interface().(uint32)
	if !ok {
		return NodeInfo{},fmt.Errorf("delegatorIDCounter is not uint32")
	}
	tokensRequestedToUnstake, ok := s.Index(12).Interface().(uint64)
	if !ok {
		return NodeInfo{},fmt.Errorf("tokensRequestedToUnstake is not uint64")
	}
	initialWeight, ok := s.Index(13).Interface().(uint64)
	if !ok {
		return NodeInfo{},fmt.Errorf("initialWeight is not uint64")
	}

	nodeInfo := NewNodeInfo(id,role,networkingAddress,networkingKey,
		stakingKey,tokensStaked,tokensCommitted,tokensUnstaking,
		tokensUnstaked,tokensRewarded,delegators,delegatorIDCounter,
		tokensRequestedToUnstake,initialWeight)
		
	return nodeInfo, nil
}

func NewNodeInfo(id string, role uint8, networkingAddress string, networkingKey string,
	stakingKey string, tokensStaked uint64, tokensCommitted uint64,
	tokensUnstaking uint64, tokensUnstaked uint64, tokensRewarded uint64,
	delegators []uint32, delegatorIDCounter uint32, tokensRequestedToUnstake uint64,
	initialWeight uint64) NodeInfo {
	return NodeInfo{
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
