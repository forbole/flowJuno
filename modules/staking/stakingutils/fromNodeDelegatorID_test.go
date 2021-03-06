package stakingutils

import "github.com/forbole/flowJuno/types"

func (suite *StakingProxyTestSuite) TestStakingProxy_getDelegatorInfo() {
	proxy := *suite.Proxy
	nodeId := "28c3544d52f7f733b23269b601a301b10c85c8359382454bfd7a80430e37827b"

	delegator := []uint32{1, 2}
	nodeInfo := types.NewStakerNodeInfo(nodeId, 2, "flow--mainnet--consensus-31.figment.io:3569",
		"f91474dc1be6e62c9284057e7727718d7e3c337671d3dfb1538935da645923e1eab46ec668e4d3da541c8ddeea257accacd0d2f188f42db0251804fc57d63bca",
		"ada64f100d7f02d56a4f195ff7c844456178130c841ee79022d3d29705a76916ed323ed6b1a758d77ef68c32a816c65316084a6859bc645166c12e78cdf321ac52010afdf24573a3988ae6f050557bb2a16d3d9573ec8ea4175862b205ec3db9",
		333448072243, 0, 0, 0, 46886804123689, delegator, 2, 0, 100)
	height := int64(1)

	outputs, err := getDelegatorInfo(nodeInfo, height, proxy)
	suite.Require().NoError(err)

	suite.Require().Len(outputs, 2)
	for i, row := range outputs {
		suite.Require().Equal(nodeId, row.NodeID)
		suite.Require().Equal(uint32(i+1), row.Id)
	}
}
