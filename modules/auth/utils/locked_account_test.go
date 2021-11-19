package utils

import(
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
