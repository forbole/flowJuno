package stakingutils

func (suite *StakingProxyTestSuite) TestProxy_getDelegatorID() {
	proxy := *suite.Proxy
	height, err := proxy.LatestHeight()
	suite.Require().NoError(err)

	

}