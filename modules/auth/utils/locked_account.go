package utils

import (
	"fmt"
	"strings"

	"github.com/forbole/flowJuno/client"
	"github.com/onflow/cadence"
	"github.com/onflow/flow-go-sdk"

	"github.com/forbole/flowJuno/types"
)

// GetLockedTokenAccounts return information of an array of locked token accounts
// if the account do not have associated locked account, it would ignore the account
func GetLockedTokenAccounts(addresses []string, height int64, client client.Proxy) ([]types.LockedAccount, error) {
	catchError := `Could not borrow a reference to public LockedAccountInfo`
	var lockedAccounts []types.LockedAccount

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

		lockedAccounts = append(lockedAccounts, types.NewLockedAccount(address, lockedAddress, balance, unlockLimit))

	}
	return lockedAccounts, nil
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
