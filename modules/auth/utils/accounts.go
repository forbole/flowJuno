package utils

import (
	"fmt"

	"github.com/forbole/flowJuno/client"
	"github.com/onflow/flow-go-sdk"

	"github.com/rs/zerolog/log"

	db "github.com/forbole/flowJuno/db/postgresql"
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
func GetAccounts(addresses []string, height int64, client client.Proxy) ([]flow.Account, error) {
	log.Debug().Str("module", "auth").Str("operation", "accounts").Int("height", int(height)).Msg("getting accounts data")
	var accounts []flow.Account

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

		accounts = append(accounts, *account)

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

	err = db.SaveAccounts(accounts)
	if err != nil {
		return err
	}

	err = UpdateLockedAccount(addresses, height, client, db)
	if err != nil {
		return err
	}

	delegatorAccount, err := GetDelegatorAccounts(addresses, height, client)
	if err != nil {
		return err
	}

	if len(delegatorAccount) != 0 {
		err = db.SaveDelegatorAccounts(delegatorAccount)
		if err != nil {
			return err
		}
	}

	stakerAccount, err := GetStakerAccounts(addresses, height, client)
	if err != nil {
		return err
	}
	if len(stakerAccount) != 0 {
		err = db.SaveStakerNodeId(stakerAccount)
	}

	return err
}

func UpdateLockedAccount(addresses []string, height int64, client client.Proxy, db *db.Db) error {
	lockedAccount, err := GetLockedAccount(addresses, height, client)
	if err != nil {
		return err
	}

	if len(lockedAccount) == 0 {
		return nil
	}

	err = db.SaveLockedAccount(lockedAccount)
	if err != nil {
		return err
	}

	LockedAccountBalance, err := GetLockedAccountBalance(addresses, height, client)
	if err != nil {
		return err
	}

	err = db.SaveLockedAccountBalance(LockedAccountBalance)
	if err != nil {
		return err
	}

	return nil
}
