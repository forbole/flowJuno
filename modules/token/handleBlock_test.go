package token

func (suite *TokenProxyTestSuite) TestTokenProxy_GetLockedAccountAddress() {
	proxy := *suite.Proxy
	supply, err := getCurrentSupply(proxy)
	suite.Require().NoError(err)
	suite.Require().True(supply > uint64(1))
}
