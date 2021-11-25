package utils

func (suite *AuthProxyTestSuite) TestProxy_getStakerNodeId() {
	proxy := *suite.Proxy
	height, err := proxy.LatestHeight()
	suite.Require().NoError(err)
	_, err = getStakerNodeId("f1830cb81484659a",height, proxy)
	suite.Require().NoError(err)
} 
