package stakingutils

func (suite *StakingProxyTestSuite) TestStakingProxy_getDelegatorID() {
	proxy := *suite.Proxy

	table,err:=getCurrentTable(1,proxy)
	suite.Require().NoError(err)

	suite.Require().True(len(table.Table)>0)
	suite.Require().Equal(int64(1),table.Height)
}

func (suite *StakingProxyTestSuite) TestStakingProxy_getWeeklyPayout() {
	proxy := *suite.Proxy

	table,err:=getWeeklyPayout(1,proxy)
	suite.Require().NoError(err)

	suite.Require().True(table.Payout>1)
	suite.Require().Equal(int64(1),table.Height)
}

func (suite *StakingProxyTestSuite) TestStakingProxy_getTotalStake() {
	proxy := *suite.Proxy

	table,err:=getTotalStake(1,proxy)
	suite.Require().NoError(err)

	suite.Require().True(table.TotalStake>1)
	suite.Require().Equal(int64(1),table.Height)
}

func (suite *StakingProxyTestSuite) TestStakingProxy_getTotalStakeByType() {
	proxy := *suite.Proxy

	table,err:=getTotalStakeByType(1,proxy)
	suite.Require().NoError(err)

	suite.Require().Len(table,5)
	for i,t:=range table{
		suite.Require().Equal(int64(1),t.Height)
		suite.Require().Equal(int8(i+1),t.Role)
	}
}

func (suite *StakingProxyTestSuite) TestStakingProxy_getStakeRequirements() {
	proxy := *suite.Proxy

	table,err:=getStakeRequirements(1,proxy)
	suite.Require().NoError(err)

	suite.Require().Len(table,5)
	for i,t:=range table{
		suite.Require().Equal(int64(1),t.Height)
		suite.Require().Equal(int8(i+1),t.Role)
	}
}

func (suite *StakingProxyTestSuite) TestStakingProxy_getProposedTable() {
	proxy := *suite.Proxy

	table,err:=getProposedTable(1,proxy)
	suite.Require().NoError(err)

	suite.Require().True(len(table.ProposedTable)>1)
	suite.Require().Equal(int64(1),table.Height)
}

func (suite *StakingProxyTestSuite) TestStakingProxy_getCutPercentage() {
	proxy := *suite.Proxy

	table,err:=getCutPercentage(1,proxy)
	suite.Require().NoError(err)

	suite.Require().True(table.CutPercentage>1)
	suite.Require().Equal(int64(1),table.Height)
}

func (suite *StakingProxyTestSuite) TestStakingProxy_GetTable() {
	proxy := *suite.Proxy

	table,err:=GetTable(1,proxy)
	suite.Require().NoError(err)

	suite.Require().True(len(table.StakingTable)>1)
	suite.Require().Equal(int64(1),table.Height)
}

func (suite *StakingProxyTestSuite) TestStakingProxy_GetNodeInfosFromTable() {
	proxy := *suite.Proxy

	table,err:=GetNodeInfosFromTable(1,proxy)
	suite.Require().NoError(err)

	suite.Require().True(len(table)>1)
}

