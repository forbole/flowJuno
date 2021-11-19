package utils

import (
	"fmt"

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

func (suite *ProxyTestSuite)TestProxy_getLockedAccountNodeInfo(){
	proxy:=*suite.Proxy
	height,err:=proxy.LatestHeight()
	suite.Require().NoError(err)

	nodeInfo,err:=getLockedAccountNodeInfo("f1830cb81484659a",height,proxy)
	suite.Require().NoError(err)

	fmt.Println(nodeInfo)
}
