package utils

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/forbole/flowJuno/client"
	"github.com/onflow/cadence"
	"github.com/onflow/flow-go-sdk"

	"github.com/rs/zerolog/log"

	db "github.com/forbole/flowJuno/db/postgresql"

	"github.com/forbole/flowJuno/types"
)

/*
//TODO: to replace getGenesisAccounts when genesis function is here
// GetGenesisAccounts parses the given appState and returns the genesis accounts
func GetGenesisAccounts(appState map[string]json.RawMessage, cdc codec.Marshaler) ([]types.Account, error) {
	log.Debug().Str("module", "auth").Msg("parsing genesis")

	var authState authtypes.GenesisState
	if err := cdc.UnmarshalJSON(appState[authtypes.ModuleName], &authState); err != nil {
		return nil, err
	}

	// Store the accounts
	accounts := make([]types.Account, len(authState.Accounts))
	for index, account := range authState.Accounts {
		var accountI authtypes.AccountI
		err := cdc.UnpackAny(account, &accountI)
		if err != nil {
			return nil, err
		}

		accounts[index] = types.NewAccount(accountI.GetAddress().String(), accountI)
	}

	return accounts, nil
} */

// --------------------------------------------------------------------------------------------------------------------

// GetAccounts returns the account data for the given addresses
func GetAccounts(addresses []string, height int64, client client.Proxy) ([]types.Account, error) {
	log.Debug().Str("module", "auth").Str("operation", "accounts").Int("height", int(height)).Msg("getting accounts data")
	var accounts []types.Account

	for _, address := range addresses {
		fmt.Println("GetAccounts:" + address)
		if address == "" {
			continue
		}
		//not working atm because of flow bug
		//account,err:=client.Client().GetAccountAtBlockHeight(client.Ctx(),flow.HexToAddress(address),uint64(height))

		account, err := client.Client().GetAccount(client.Ctx(), flow.HexToAddress(address))

		if err != nil {
			return nil, err
		}

		if account == nil {
			return nil, fmt.Errorf("address is not valid and cannot get details")
		}

		accounts = append(accounts, types.NewAccount(account.Address.String()))

	}

	return accounts, nil
}

// UpdateAccounts takes the given addresses and for each one queries the chain
// retrieving the account data and stores it inside the database.
func UpdateAccounts(addresses []string, db *db.Db, height int64, client client.Proxy) error {
	accounts, err := GetAccounts(addresses, height, client)
	if err != nil {
		return err
	}

	lockedAccount, err := GetLockedTokenAccounts(addresses, height, client)
	if err != nil {
		return err
	}

	delegatorAccount, err := GetDelegatorAccounts(addresses, height, client)
	if err != nil {
		return err
	}

	err = db.SaveAccounts(accounts)
	if err != nil {
		return err
	}

	err = db.SaveLockedTokenAccounts(lockedAccount)
	if err != nil {
		return err
	}

	err = db.SaveDelegatorAccount(delegatorAccount)
	if err != nil {
		return err
	}
	return nil
}

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

// getLockedTokenAccountBalance get the account balance by address
func getLockedTokenAccountBalance(address string, height int64, client client.Proxy) (float64, error) {
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

	var balance float64
	value, err := client.Client().ExecuteScriptAtLatestBlock(client.Ctx(), []byte(script), candenceArr)
	if err != nil {
		return 0, err
	}
	fmt.Println("Locked Account Balance!" + value.String())
	balance, err = strconv.ParseFloat(value.String(), 64)
	if err != nil {
		return 0, err
	}
	return balance, nil
}

// getLockedTokenAccountUnlockLimit get the unlock limit by address
func getLockedTokenAccountUnlockLimit(address string, height int64, client client.Proxy) (float64, error) {
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

	var limit float64
	value, err := client.Client().ExecuteScriptAtLatestBlock(client.Ctx(), []byte(script), candenceArr)
	if err != nil {
		return 0, err
	}
	fmt.Println("Locked Account Unlock Limit!" + value.String())
	limit, err = strconv.ParseFloat(value.String(), 64)
	if err != nil {
		return 0, err
	}
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

func getStakerNodeID(address string, height int64, client client.Proxy) (string, error) {
	script := fmt.Sprintf(`
	import LockedTokens from %s

	import LockedTokens from ${global.contracts.LockedTokens}
	pub fun main(account: Address): String {
		let lockedAccountInfoRef = getAccount(account)
			.getCapability<&LockedTokens.TokenHolder{LockedTokens.LockedAccountInfo}>(LockedTokens.LockedAccountInfoPublicPath)!
			.borrow() ?? panic("Could not borrow a reference to public LockedAccountInfo")
	
		return lockedAccountInfoRef.getNodeID()!
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

// Danger zone

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

func getStakerNodeInfo(address string, height int64, client client.Proxy) (string, error) {
	script := fmt.Sprintf(`
	import FlowIDTableStaking from %s
	import LockedTokens from %s
	// Returns an array of NodeInfo objects that the account controls
	// in its normal account and shared account
	
	pub fun main(account: Address): [FlowIDTableStaking.NodeInfo] {
	
		let nodeInfoArray: [FlowIDTableStaking.NodeInfo] = []
	
		let pubAccount = getAccount(account)
	
		let nodeStaker = pubAccount.getCapability<&{FlowIDTableStaking.NodeStakerPublic}>(FlowIDTableStaking.NodeStakerPublicPath)!
			.borrow()
	
		if let nodeRef = nodeStaker {
			nodeInfoArray.append(FlowIDTableStaking.NodeInfo(nodeID: nodeRef.id))
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
	
			nodeInfoArray.append(FlowIDTableStaking.NodeInfo(nodeID: lockedAccountInfoRef.getNodeID()!))
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
		return "", err
	}

	fmt.Println("Locked Account" + value.String())

	return value.String(), nil
}
