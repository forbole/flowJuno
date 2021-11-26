package utils

func (suite *AuthProxyTestSuite) TestProxy_getStakerNodeId() {
	proxy := *suite.Proxy
	height, err := proxy.LatestHeight()
	suite.Require().NoError(err)
	nodeIds, err := getStakerNodeId("f1830cb81484659a",height, proxy)
	suite.Require().NoError(err)

	suite.Require().Equal(nodeIds[0],"e7df1454826425251716a703e907981672a43208ef3eabfc95d593673da778f6") 
} 
