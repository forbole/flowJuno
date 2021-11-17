package utils

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/forbole/flowJuno/client"
	"github.com/forbole/flowJuno/modules/utils"
	"github.com/onflow/cadence"
	"github.com/onflow/flow-go-sdk"

	"github.com/forbole/flowJuno/types"
)

// GetLockedTokenAccounts return information of an array of locked token accounts
// if the account do not have associated locked account, it would ignore the account
func GetDelegatorAccounts(addresses []string, height int64, client client.Proxy) ([]types.DelegatorAccount, error) {
	catchError := `Could not borrow a reference to public LockedAccountInfo`
	catchError2 := `unexpectedly found nil while forcing an Optional value`
	var delegatorAccount []types.DelegatorAccount

	for _, address := range addresses {
		if address == "" {
			continue
		}

		delegatorId, err := getDelegatorID(address, height, client)
		if err != nil {
			if strings.Contains(err.Error(), catchError) || strings.Contains(err.Error(), catchError2) {
				continue
			}
			return nil, err
		}

		delegatorNodeId, err := getDelegatorNodeID(address, height, client)
		if err != nil {
			if strings.Contains(err.Error(), catchError) || strings.Contains(err.Error(), catchError2) {
				continue
			}
			return nil, err
		}

		delegatorAccount = append(delegatorAccount, types.NewDelegatorAccount(address, delegatorId, delegatorNodeId))

	}
	return delegatorAccount, nil
}

// getDelegatorID return the delegator who staked in a locked account
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

// getDelegatorNodeID get locked account delegator node id
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

	nodeId, err := utils.CadanceConvertString(value)
	if err != nil {
		return "", err
	}

	return nodeId, nil
}
