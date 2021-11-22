package utils

import (
	"github.com/forbole/flowJuno/types"
)

func (suite *ProxyTestSuite)TestProxy_GetLockedAccountAddress(){
	proxy:=*suite.Proxy
	height,err:=proxy.LatestHeight()
	suite.Require().NoError(err)
	addresses:=[]string{
		"f1830cb81484659a",
	}
	lockedAccount,err:=GetLockedAccount(addresses,height,proxy)
	suite.Require().NoError(err)

	expected:=[]types.LockedAccount{
		types.NewLockedAccount("f1830cb81484659a","0x15e4b565057e6545"),
	}

	suite.Require().Equal(lockedAccount[0],expected[0])
}

func (suite *ProxyTestSuite)TestProxy_GetLockedAccountBalance(){
	proxy:=*suite.Proxy
	height,err:=proxy.LatestHeight()
	suite.Require().NoError(err)

	balance,err:=getLockedTokenAccountBalance("f1830cb81484659a",height,proxy)
	suite.Require().NoError(err)

	suite.Require().Equal(uint64(0),balance)
}

func (suite *ProxyTestSuite)TestProxy_getLockedTokenAccountUnlockLimit(){
	proxy:=*suite.Proxy
	height,err:=proxy.LatestHeight()
	suite.Require().NoError(err)

	balance,err:=getLockedTokenAccountUnlockLimit("f1830cb81484659a",height,proxy)
	suite.Require().NoError(err)

	suite.Require().Equal(uint64(100000),balance)
}

func (suite *ProxyTestSuite)TestProxy_getDelegatorNodeInfo(){
	proxy:=*suite.Proxy
	height,err:=proxy.LatestHeight()
	suite.Require().NoError(err)

	nodeInfo,err:=getDelegatorNodeInfo("808b03495a0408bb",height,proxy)
	suite.Require().NoError(err)

	expected:=[]types.DelegatorNodeInfo{
		types.NewDelegatorNodeInfo(
		 3905,
		"2cfab7e9163475282f67186b06ce6eea7fa0687d25dd9c7a84532f2016bc2e5e",
		0,0, 0, 0, 0, 0,
	  )}

	suite.Require().Equal(expected,nodeInfo)
	}
