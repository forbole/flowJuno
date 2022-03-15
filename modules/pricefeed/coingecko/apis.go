package coingecko

import (
	"math"
	"net/http"
	coingecko "github.com/superoo7/go-gecko/v3"
	"time"
	"github.com/forbole/flowJuno/types"
)

type CoingeckoClient struct{
	client *coingecko.Client
}

func NewCoingeckoClient(timeout int)CoingeckoClient{
	httpClient := &http.Client{
		Timeout: time.Second * time.Duration(10),
	}
	client:=coingecko.NewClient(httpClient)
	return CoingeckoClient{
		client:client,
	}
}

// GetCoinsList allows to fetch from the remote APIs the list of all the supported tokens
func (c CoingeckoClient) GetCoinsList() (coins Tokens, err error) {
	//err = queryCoinGecko("/coins/list", &coins)
	coinlist,err:=c.client.CoinsList()
	if err!=nil{
		return nil,err
	}
	token:=make(Tokens,len(*coinlist))
	for i,coin:=range *coinlist{
		token[i]=NewToken(coin.ID,coin.Symbol,coin.Name)
	}
	return coins, err
}

// GetTokensPrices queries the remote APIs to get the token prices of all the tokens having the given ids
func (c CoingeckoClient) GetTokensPrices(ids []string) ([]types.TokenPrice, error) {
	//query := "/coins/markets?vs_currency=usd&ids=flow"
	//err := queryCoinGecko(query, &prices)
	prices,err:=c.client.CoinsMarket("usd",[]string{"flow"},"",0,0,false,nil)
	if err != nil {
		return nil, err
	}
	tokenPrices := make([]types.TokenPrice, len(*prices))
	for i, price := range *prices {
		timestamp,err:=time.Parse("2021-07-30 00:00:00 +0000 UTC",price.LastUpdated)
		if err!=nil{
			return nil,err
		}
		tokenPrices[i] = types.NewTokenPrice(
			price.Symbol,
			price.CurrentPrice,
			int64(math.Trunc(price.MarketCap)),
			timestamp,
		)
	}


	return tokenPrices, nil
}