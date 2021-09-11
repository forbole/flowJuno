package utils

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/forbole/flowJuno/client"
	"github.com/onflow/cadence"
	"github.com/onflow/flow-go-sdk"

	"github.com/forbole/flowJuno/types"
)

// GetLockedTokenAccounts return information of an array of locked token accounts
// if the account do not have associated locked account, it would ignore the account
func GetDelegatorAccounts(addresses []string, height int64, client client.Proxy) ([]types.DelegatorAccount, error) {
	catchError := `Could not borrow a reference to public LockedAccountInfo`
	catchError2:=`unexpectedly found nil while forcing an Optional value`
	var delegatorAccount []types.DelegatorAccount

	for _, address := range addresses {
		if address == "" {
			continue
		}

		delegatorId, err := getDelegatorID(address, height, client)
		if err != nil {
			if (strings.Contains(err.Error(), catchError)||strings.Contains(err.Error(), catchError2)){
				continue
			}
			return nil, err
		}

		delegatorNodeId, err := getDelegatorNodeID(address, height, client)
		if err != nil {
			if (strings.Contains(err.Error(), catchError)||strings.Contains(err.Error(), catchError2)) {
				continue
			}
			return nil, err
		}

		delegatorNodeInfo, err := getDelegatorNodeInfo(address, height, client)
		if err != nil {
			if (strings.Contains(err.Error(), catchError)||strings.Contains(err.Error(), catchError2)) {
				continue
			}
			return nil, err
		}

		delegatorAccount = append(delegatorAccount, types.NewDelegatorAccount(address, delegatorId, delegatorNodeId, delegatorNodeInfo))

	}
	return delegatorAccount, nil
}


func getDelegatorID(address string, height int64, client client.Proxy) (int64, error) {
	script := fmt.Sprintf(`
	import LockedTokens from %s

	pub fun main(account: Address): UInt32 {

		let lockedAccountInfoRef = getAccount(account)
			.getCapability<&LockedTokens.TokenHolder{LockedTokens.LockedAccountInfo}>(
				LockedTokens.LockedAccountInfoPublicPath
			)
			.borrow()
			?? panic("Could not borrow a reference to public LockedAccountInfo")
	
		return lockedAccountInfoRef.getDelegatorID()!
	}
	`, client.Contract().LockedTokens)

	flowAddress := flow.HexToAddress(address)
	candanceAddress := cadence.Address(flowAddress)
	//val,err:=cadence.NewValue(candanceAddress)
	candenceArr := []cadence.Value{candanceAddress}

	value, err := client.Client().ExecuteScriptAtLatestBlock(client.Ctx(), []byte(script), candenceArr)
	if err != nil {
		return 0, err
	}

	fmt.Println("Locked Account" + value.String())
	id, err := strconv.ParseInt(value.String(), 10, 32)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func getDelegatorNodeID(address string, height int64, client client.Proxy) (string, error) {
	script := fmt.Sprintf(`
	import LockedTokens from %s

	pub fun main(account: Address): String {
	
		let lockedAccountInfoRef = getAccount(account)
			.getCapability<&LockedTokens.TokenHolder{LockedTokens.LockedAccountInfo}>(
				LockedTokens.LockedAccountInfoPublicPath
			)
			.borrow()
			?? panic("Could not borrow a reference to public LockedAccountInfo")
	
		return lockedAccountInfoRef.getDelegatorNodeID()!
	}
	`, client.Contract().LockedTokens)

	flowAddress := flow.HexToAddress(address)
	candanceAddress := cadence.Address(flowAddress)
	//val,err:=cadence.NewValue(candanceAddress)
	candenceArr := []cadence.Value{candanceAddress}

	value, err := client.Client().ExecuteScriptAtLatestBlock(client.Ctx(), []byte(script), candenceArr)
	if err != nil {
		return "", err
	}

	fmt.Println("Locked Account" + value.String())
	if err != nil {
		return "", err
	}

	return value.String(), nil
}

func getDelegatorNodeInfo(address string, height int64, client client.Proxy) (string, error) {
	script := fmt.Sprintf(`
	import FlowIDTableStaking from %s
	import LockedTokens from %s
	// Returns an array of DelegatorInfo objects that the account controls
	// in its normal account and shared account
	pub fun main(account: Address): [FlowIDTableStaking.DelegatorInfo] {
		let delegatorInfoArray: [FlowIDTableStaking.DelegatorInfo] = []
		let pubAccount = getAccount(account)
		let delegator = pubAccount.getCapability<&{FlowIDTableStaking.NodeDelegatorPublic}>(/public/flowStakingDelegator)!
			.borrow()
		if let delegatorRef = delegator {
			delegatorInfoArray.append(FlowIDTableStaking.DelegatorInfo(nodeID: delegatorRef.nodeID, delegatorID: delegatorRef.id))
		}
		let lockedAccountInfoCap = pubAccount
			.getCapability
			<&LockedTokens.TokenHolder{LockedTokens.LockedAccountInfo}>
			(LockedTokens.LockedAccountInfoPublicPath)
		if lockedAccountInfoCap == nil || !(lockedAccountInfoCap!.check()) {
			return delegatorInfoArray
		}
		let lockedAccountInfo = lockedAccountInfoCap!.borrow()
		if let lockedAccountInfoRef = lockedAccountInfo {
			let nodeID = lockedAccountInfoRef.getDelegatorNodeID()
			let delegatorID = lockedAccountInfoRef.getDelegatorID()
			if (nodeID == nil || delegatorID == nil) {
				return delegatorInfoArray
			}
			delegatorInfoArray.append(FlowIDTableStaking.DelegatorInfo(nodeID: nodeID!, delegatorID: delegatorID!))
		}
		return delegatorInfoArray
	}`, client.Contract().StakingTable, client.Contract().LockedTokens)

	flowAddress := flow.HexToAddress(address)
	candanceAddress := cadence.Address(flowAddress)
	//val,err:=cadence.NewValue(candanceAddress)
	candenceArr := []cadence.Value{candanceAddress}

	value, err := client.Client().ExecuteScriptAtLatestBlock(client.Ctx(), []byte(script), candenceArr)
	if err != nil {
		return "", err
	}

	fmt.Println("Locked Account" + value.String())

	return value.String(), nil
}