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

// GetLockedTokenAccounts return information of an array of locked token accounts
// if the account do not have associated locked account, it would ignore the account
func GetStakerAccounts(addresses []string, height int64, client client.Proxy) ([]types.StakerNodeId, error) {
	catchError := `Could not borrow a reference to public LockedAccountInfo`
	catchError2 := `unexpectedly found nil while forcing an Optional value`

	var stakerAccounts []types.StakerNodeId

	for _, address := range addresses {
		if address == "" {
			continue
		}

		stakerNodeInfos, err := getStakerNodeId(address, height, client)
		if err != nil {
			if strings.Contains(err.Error(), catchError) || strings.Contains(err.Error(), catchError2) {
				continue
			}
			return nil, err
		}

		for _,nodeinfo:=range stakerNodeInfos{
			stakerAccounts = append(stakerAccounts,types.NewStakerNodeId(address,nodeinfo))
		}
	}
	return stakerAccounts, nil
}
// Danger zone

func getStakerNodeId(address string, height int64, client client.Proxy) ([]string, error) {
	script := fmt.Sprintf(`
	import FlowIDTableStaking from %s
	import LockedTokens from %s
	// Returns an array of node_info_id string that the account controls
	// in its normal account and shared account
	
	pub fun main(account: Address): [String] {
	
		let nodeInfoArray: [String] = []
	
		let pubAccount = getAccount(account)
	
		let nodeStaker = pubAccount.getCapability<&{FlowIDTableStaking.NodeStakerPublic}>(FlowIDTableStaking.NodeStakerPublicPath)!
			.borrow()
	
		if let nodeRef = nodeStaker {
			nodeInfoArray.append(nodeRef.id!)
		}
	
		let lockedAccountInfoCap = pubAccount
			.getCapability
			<&LockedTokens.TokenHolder{LockedTokens.LockedAccountInfo}>
			(LockedTokens.LockedAccountInfoPublicPath)
	
		if lockedAccountInfoCap == nil || !(lockedAccountInfoCap!.check()) {
			return nodeInfoArray
		}
	
		if let lockedAccountInfoRef = lockedAccountInfoCap!.borrow() {
		
			if (lockedAccountInfoRef.getNodeID() == nil) {
				return nodeInfoArray
			}
	
			nodeInfoArray.append(FlowIDTableStaking.NodeInfo(lockedAccountInfoRef.getNodeID()!)
		}
		
		log(nodeInfoArray)
		
		return nodeInfoArray
     }`, client.Contract().StakingTable, client.Contract().LockedTokens)

	flowAddress := flow.HexToAddress(address)
	candanceAddress := cadence.Address(flowAddress)
	//val,err:=cadence.NewValue(candanceAddress)
	candenceArr := []cadence.Value{candanceAddress}

	value, err := client.Client().ExecuteScriptAtLatestBlock(client.Ctx(), []byte(script), candenceArr)
	if err != nil {
		return nil, err
	}

	stakerNodeInfo, err := utils.CadenceConvertStringArray(value)
	if err != nil {
		return nil, err
	}


	return stakerNodeInfo, nil
}
