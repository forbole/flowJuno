package pricefeed

import "github.com/forbole/flowJuno/types"

type HistoryModule interface {
	UpdatePricesHistory([]types.TokenPrice) error
}
