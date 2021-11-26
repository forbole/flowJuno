package stakingutils

func (suite *StakingProxyTestSuite) TestProxy_getDelegatorID() {
	proxy := *suite.Proxy

	table,err:=getCurrentTable(1,proxy)
	suite.Require().NoError(err)

	suite.Require().True(len(table.Table)>1)
	suite.Require().Equal(int64(1),table.Height)
}

