package utils

import (
	"testing"

	"github.com/forbole/flowJuno/client"
	"github.com/forbole/flowJuno/types"
	"github.com/stretchr/testify/suite"

	_ "github.com/proullon/ramsql/driver"
)

func TestProxyTestSuite(t *testing.T) {
	suite.Run(t, new(ProxyTestSuite))
}



type ProxyTestSuite struct {
	suite.Suite

	Proxy *client.Proxy
}

func (suite *ProxyTestSuite) SetupTest() {
	rpcConfig:=types.NewRPCConfig("","access.mainnet.nodes.onflow.org:9000","Mainnet")
	modules:=[]string{
		"auth","messages","staking","consensus","token"}
	cosmosConfig:=types.NewCosmosConfig("",modules,19050753)

	config:=types.NewConfig(rpcConfig,nil,cosmosConfig,nil,nil,nil,nil)

	proxy,err := client.NewClientProxy(config,nil)
	suite.Require().NoError(err)

	suite.Proxy = proxy
}

