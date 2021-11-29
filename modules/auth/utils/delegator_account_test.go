package utils

func (suite *AuthProxyTestSuite) TestProxy_getDelegatorID() {
	proxy := *suite.Proxy
	height, err := proxy.LatestHeight()
	suite.Require().NoError(err)
	address := "808b03495a0408bb"
	id, err := getDelegatorID(address, height, proxy)
	suite.Require().NoError(err)
	suite.Require().Equal(uint32(3905), id)
}

func (suite *AuthProxyTestSuite) TestProxy_getDelegatorNodeID() {
	proxy := *suite.Proxy
	height, err := proxy.LatestHeight()
	suite.Require().NoError(err)
	address := "808b03495a0408bb"
	id, err := getDelegatorNodeID(address, height, proxy)
	suite.Require().NoError(err)
	suite.Require().Equal("2cfab7e9163475282f67186b06ce6eea7fa0687d25dd9c7a84532f2016bc2e5e", id)
}
