package utils

import (
	"fmt"
	"strings"

	"github.com/forbole/flowJuno/client"
	"github.com/forbole/flowJuno/modules/utils"
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

		lockedAccountBalances = append(lockedAccountBalances, types.NewLockedAccountBalance(lockedAddress, balance, unlockLimit, uint64(height)))

	}
	return lockedAccountBalances, nil
}

// GetLockedAccount return an array of locked account limit which is network constant
func GetLockedAccount(addresses []string, height int64, client client.Proxy) ([]types.LockedAccount, error) {
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
			return nil, fmt.Errorf("fail to get locked account address :%s", err)
		}

		lockedAccount = append(lockedAccount, types.NewLockedAccount(address, lockedAddress))

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

	value, err := client.Client().ExecuteScriptAtLatestBlock(client.Ctx(), []byte(script), candenceArr)
	if err != nil {
		return 0, err
	}
	balance, err := utils.CadenceConvertUint64(value)
	if err != nil {
		return 0, err
	}
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
	return limit, nil

}

// getLockedTokenAccountAddress get the locked account address associated with the input address
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

	val, ok := value.(cadence.Address)
	if !ok {
		return "", fmt.Errorf("Not a cadence address")
	}

	return val.String(), nil
}

// getDelegatorNodeInfo get delegator info associated to the address
func getDelegatorNodeInfo(address string, height int64, client client.Proxy) ([]types.DelegatorNodeInfo, error) {
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
		return nil, err
	}

	nodeInfos, err := types.DelegatorNodeInfoArrayFromCadence(value)
	if err != nil {
		return nil, err
	}

	return nodeInfos, nil
}
