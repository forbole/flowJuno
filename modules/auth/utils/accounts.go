package utils

import (
	"fmt"

	"github.com/forbole/flowJuno/client"
	"github.com/forbole/flowJuno/types"
	"github.com/onflow/flow-go-sdk"

	"github.com/rs/zerolog/log"

	db "github.com/forbole/flowJuno/db/postgresql"
)

// GetAccounts returns the account data for the given addresses
func GetAccounts(addresses []string, height int64, client client.Proxy) ([]types.Account, error) {
	log.Debug().Str("module", "auth").Str("operation", "accounts").Int("height", int(height)).Msg("getting accounts data")
	var accounts []types.Account

	for _, address := range addresses {
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

		newAccount, err := types.NewAccount(*account)
		if err != nil {
			return nil, fmt.Errorf("Cannot Get Account: %s", err)
		}

		accounts = append(accounts, newAccount)

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

	err = db.SaveAccounts(accounts, uint64(height))
	if err != nil {
		return err
	}

	err = UpdateLockedAccount(addresses, height, client, db)
	if err != nil {
		return err
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

// UpdateLockedAccount get all the details that need for locked account
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

	var delegatorsAccounts []types.DelegatorAccount
	for _, address := range addresses {
		accountdelegators, err := getDelegatorNodeInfo(address, height, client)
		if err != nil {
			return fmt.Errorf("cannot get delegators from address: %s", err)
		}
		if accountdelegators == nil {
			continue
		}

		for _, delegator := range accountdelegators {
			delegatorsAccounts = append(delegatorsAccounts, types.NewDelegatorAccount(address, int64(delegator.Id), delegator.NodeID))
		}
	}

	err = db.SaveDelegatorAccounts(delegatorsAccounts)
	if err != nil {
		return fmt.Errorf("cannot save delegators from address: %s", err)
	}

	return nil
}
