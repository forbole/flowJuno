package types

import (
	"fmt"

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
	StakerNodeInfo string
}

func NewStakerAccount(address, stakerNodeId, stakerNodeInfo string) StakerAccount {
	return StakerAccount{
		Address:        address,
		StakerNodeId:   stakerNodeId,
		StakerNodeInfo: stakerNodeInfo,
	}
}

type DelegatorNodeInfo struct{
	Id uint32
	NodeID string
	TokensCommitted uint64
	TokensStaked uint64
	TokensUnstaking uint64 
	TokensRewarded uint64
	TokensUnstaked uint64 
	TokensRequestedToUnstake uint64
}

func NewDelegatorNodeInfo(id uint32,
	nodeID string,
	tokensCommitted uint64,
	tokensStaked uint64,
	tokensUnstaking uint64 ,
	tokensRewarded uint64,
	tokensUnstaked uint64 ,
	tokensRequestedToUnstake uint64) DelegatorNodeInfo {
		return DelegatorNodeInfo{
			Id :id ,
			NodeID :nodeID,
			TokensCommitted :tokensCommitted ,
			TokensStaked :tokensStaked ,
			TokensUnstaking  :tokensUnstaking , 
			TokensRewarded :tokensRewarded ,
			TokensUnstaked  :tokensUnstaked  ,
			TokensRequestedToUnstake:tokensRequestedToUnstake , 
		}
}

func DelegatorNodeInfoFromCadence(value cadence.Value)(DelegatorNodeInfo,error){
	arrayValue := value.(cadence.Array)

	fields:=arrayValue.Values[0].(cadence.Struct).Fields
	
	id:=fields[0].ToGoValue()
	
	nodeID:=fields[1].ToGoValue()
	nodeID,ok:=nodeID.(string)
	if !ok{
		fmt.Errorf("nodeID is not a string!!")
	}
	tokenCommited:=fields[2].ToGoValue()
	tokenStaked:=fields[3].ToGoValue()
	tokensUnstaking:=fields[4].ToGoValue()
	tokenRewarded:=fields[5].ToGoValue()
	tokenUnstaked:=fields[6].ToGoValue()
	tokenRequestedToUnstake:=fields[7].ToGoValue()
	
	return NewDelegatorNodeInfo(id,nodeID,tokenCommited,tokenStaked,tokensUnstaking,
		tokenRewarded,tokenUnstaked,tokenRequestedToUnstake)


}