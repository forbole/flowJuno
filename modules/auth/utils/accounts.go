package utils

import (
	"fmt"

	"github.com/forbole/flowJuno/client"
	"github.com/onflow/flow-go-sdk"

	"github.com/rs/zerolog/log"

	db "github.com/forbole/flowJuno/db/postgresql"


	"github.com/forbole/flowJuno/types"
)

/*
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
	log.Debug().Str("module", "auth").Str("operation", "accounts").Msg("getting accounts data")
	var accounts []types.Account

	for _, address := range addresses {
		account,err:=client.Client().GetAccount(client.Ctx(),flow.HexToAddress(address))
		if err != nil {
			log.Error().Str("module", "auth").Err(err).Int64("height", height).
				Str("Auth", "Get Account").Msg("error while getting accounts")
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
	fmt.Println(accounts)

	return db.SaveAccounts(accounts)
}
