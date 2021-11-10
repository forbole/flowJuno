package utils

import (
	"fmt"
	"strings"

	"github.com/forbole/flowJuno/client"
	"github.com/onflow/cadence"
	"github.com/onflow/flow-go-sdk"

	"github.com/forbole/flowJuno/types"
)



// GetLockedAccountBalance return information of an array of locked token accounts
// contain balance unlock limit which is variable of locked account
func GetLockedAccountBalance(addresses []string, height int64, client client.Proxy) ([]types.LockedAccountBalance, error) {
	catchError := `Could not borrow a reference to public LockedAccountInfo`
	var lockedAccountBalances []types.LockedAccountBalance

	for _, address := range addresses {
		if address == "" {
			continue
		}

		lockedAddress, err := getLockedTokenAccountAddress(address, height, client)
		if err != nil {
			if strings.Contains(err.Error(), catchError) {
				continue
			}
			return nil, err
		}

		balance, err := getLockedTokenAccountBalance(address, height, client)
		if err != nil {
			if strings.Contains(err.Error(), catchError) {
				continue
			}
			return nil, err
		}

		unlockLimit, err := getLockedTokenAccountUnlockLimit(address, height, client)
		if err != nil {
			if strings.Contains(err.Error(), catchError) {
				continue
			}
			return nil, err
		}

		lockedAccountBalances = append(lockedAccountBalances, types.NewLockedAccountBalance(lockedAddress, balance, unlockLimit,uint64(height)))

	}
	return lockedAccountBalances, nil
}


// GetLockedAccount return an array of locked account limit which is network constant
func GetLockedAccount(addresses []string, height int64, client client.Proxy)([]types.LockedAccount,error){
	catchError := `Could not borrow a reference to public LockedAccountInfo`
	var lockedAccount []types.LockedAccount

	for _, address := range addresses {
		if address == "" {
			continue
		}

		lockedAddress, err := getLockedTokenAccountAddress(address, height, client)
		if err != nil {
			if strings.Contains(err.Error(), catchError) {
				continue
			}
			return nil, err
		}

		nodeInfo, err := getLockedAccountNodeInfo(address, height, client)
		if err != nil {
			if strings.Contains(err.Error(), catchError) {
				continue
			}
			return nil, err
		}

		for _,node:=range nodeInfo{
			lockedAccount = append(lockedAccount, types.NewLockedAccount(address,lockedAddress,node.NodeID,uint64(node.Id)))

		}

	}
	return lockedAccount, nil

}

// getLockedTokenAccountBalance get the account balance by address
func getLockedTokenAccountBalance(address string, height int64, client client.Proxy) (uint64, error) {
	script := fmt.Sprintf(`
	import LockedTokens from %s

	pub fun main(account: Address): UFix64 {
	
		let lockedAccountInfoRef = getAccount(account)
			.getCapability<&LockedTokens.TokenHolder{LockedTokens.LockedAccountInfo}>(
				LockedTokens.LockedAccountInfoPublicPath
			)
			.borrow()
			?? panic("Could not borrow a reference to public LockedAccountInfo")
	
		return lockedAccountInfoRef.getLockedAccountBalance()
	}
	`, client.Contract().LockedTokens)

	flowAddress := flow.HexToAddress(address)
	candanceAddress := cadence.Address(flowAddress)
	//val,err:=cadence.NewValue(candanceAddress)
	candenceArr := []cadence.Value{candanceAddress}

	var balance uint64
	value, err := client.Client().ExecuteScriptAtLatestBlock(client.Ctx(), []byte(script), candenceArr)
	if err != nil {
		return 0, err
	}
	balance, ok := value.ToGoValue().(uint64)
	if !ok {
		return 0, fmt.Errorf("cadence script does not return a uint64 value")
	}
	fmt.Printf("Locked Account Unlock Limit! %d", balance)

	return balance, nil
}

// getLockedTokenAccountUnlockLimit get the unlock limit by address
func getLockedTokenAccountUnlockLimit(address string, height int64, client client.Proxy) (uint64, error) {
	script := fmt.Sprintf(`
	import LockedTokens from %s

	pub fun main(account: Address): UFix64 {
	
		let lockedAccountInfoRef = getAccount(account)
			.getCapability<&LockedTokens.TokenHolder{LockedTokens.LockedAccountInfo}>(
				LockedTokens.LockedAccountInfoPublicPath
			)
			.borrow()
			?? panic("Could not borrow a reference to public LockedAccountInfo")
	
		return lockedAccountInfoRef.getUnlockLimit()
	}
	`, client.Contract().LockedTokens)

	flowAddress := flow.HexToAddress(address)
	candanceAddress := cadence.Address(flowAddress)
	//val,err:=cadence.NewValue(candanceAddress)
	candenceArr := []cadence.Value{candanceAddress}

	var limit uint64
	value, err := client.Client().ExecuteScriptAtLatestBlock(client.Ctx(), []byte(script), candenceArr)
	if err != nil {
		return 0, err
	}

	limit, ok := value.ToGoValue().(uint64)
	if !ok {
		return 0, fmt.Errorf("cadence script does not return a uint64 value")
	}
	fmt.Printf("Locked Account Unlock Limit! %d", limit)
	return limit, nil

}

func getLockedTokenAccountAddress(address string, height int64, client client.Proxy) (string, error) {
	script := fmt.Sprintf(`
	import LockedTokens from %s

	pub fun main(account: Address): Address {
	
		let lockedAccountInfoRef = getAccount(account)
			.getCapability<&LockedTokens.TokenHolder{LockedTokens.LockedAccountInfo}>(
				LockedTokens.LockedAccountInfoPublicPath
			)
			.borrow()
			?? panic("Could not borrow a reference to public LockedAccountInfo")
	
		return lockedAccountInfoRef.getLockedAccountAddress()
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

	return value.String(), nil
}


func getLockedAccountNodeInfo(address string, height int64, client client.Proxy) ([]types.DelegatorNodeInfo, error) {
	script := fmt.Sprintf(`
	import FlowIDTableStaking from %s
	import LockedTokens from %s
	// Returns locked account's delegator info objects that the account controls
	pub fun main(account: Address): FlowIDTableStaking.DelegatorInfo {
		let delegatorInfo:FlowIDTableStaking.DelegatorInfo
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
				return nil
			}
			delegatorInfo = infoFlowIDTableStaking.DelegatorInfo(nodeID: nodeID!, delegatorID: delegatorID!))
		}
		return delegatorInfo
	}`, client.Contract().StakingTable, client.Contract().LockedTokens)

	flowAddress := flow.HexToAddress(address)
	candanceAddress := cadence.Address(flowAddress)
	//val,err:=cadence.NewValue(candanceAddress)
	candenceArr := []cadence.Value{candanceAddress}

	value, err := client.Client().ExecuteScriptAtLatestBlock(client.Ctx(), []byte(script), candenceArr)
	if err != nil {
		return nil, err
	}

	nodeInfos, err := types.DelegatorNodeInfoArrayFromCadence(value)
	if err != nil {
		return nil, err
	}

	
	return nodeInfos, nil
}
