package postgresql_test

import (
	"fmt"

	dbtypes "github.com/forbole/flowJuno/db/types"
	"github.com/forbole/flowJuno/types"
)

func (suite *DbTestSuite) TestBigDipperDb_TotalStakeByType() {

	// ------------------------------
	// --- Prepare the data
	// ------------------------------

	input := []types.TotalStakeByType{
		types.NewTotalStakeByType(1, 2, 3),
	}

	// ------------------------------
	// --- Save the data
	// ------------------------------

	err := suite.database.SaveTotalStakeByType(input)
	suite.Require().NoError(err)

	// ------------------------------
	// --- Verify the data
	// ------------------------------
	expectedRow := dbtypes.NewTotalStakeByTypeRow(1, 2, 3)
	var outputs []dbtypes.TotalStakeByTypeRow
	err = suite.database.Sqlx.Select(&outputs, `SELECT * FROM total_stake_by_type`)
	suite.Require().NoError(err)
	suite.Require().Len(outputs, 1, "should contain only one row")
	fmt.Println(outputs[0])
	fmt.Println(expectedRow)

	suite.Require().True(expectedRow.Equal(outputs[0]))

}

func (suite *DbTestSuite) TestBigDipperDb_StakeRequirements() {

	// ------------------------------
	// --- Prepare the data
	// ------------------------------

	/*  TODO: Prepare parameter    */

	input := []types.StakeRequirements{
		types.NewStakeRequirements(1, 2, 3),
	}

	// ------------------------------
	// --- Save the data
	// ------------------------------

	err := suite.database.SaveStakeRequirements(input)
	suite.Require().NoError(err)

	// ------------------------------
	// --- Verify the data
	// ------------------------------
	expectedRow := dbtypes.NewStakeRequirementsRow(1, 2, 3)
	var outputs []dbtypes.StakeRequirementsRow
	err = suite.database.Sqlx.Select(&outputs, `SELECT * FROM stake_requirements`)
	suite.Require().NoError(err)
	suite.Require().Len(outputs, 1, "should contain only one row")
	suite.Require().True(expectedRow.Equal(outputs[0]))

}

func (suite *DbTestSuite) TestBigDipperDb_WeeklyPayout() {

	// ------------------------------
	// --- Prepare the data
	// ------------------------------

	/*  TODO: Prepare parameter    */

	input := types.NewWeeklyPayout(1, 2)

	// ------------------------------
	// --- Save the data
	// ------------------------------

	err := suite.database.SaveWeeklyPayout(input)
	suite.Require().NoError(err)

	// ------------------------------
	// --- Verify the data
	// ------------------------------
	expectedRow := dbtypes.NewWeeklyPayoutRow(1, 2)
	var outputs []dbtypes.WeeklyPayoutRow
	err = suite.database.Sqlx.Select(&outputs, `SELECT * FROM weekly_payout`)
	suite.Require().NoError(err)
	suite.Require().Len(outputs, 1, "should contain only one row")
	suite.Require().True(expectedRow.Equal(outputs[0]))

}

func (suite *DbTestSuite) TestBigDipperDb_TotalStake() {

	// ------------------------------
	// --- Prepare the data
	// ------------------------------

	input := types.NewTotalStake(1, 2)

	// ------------------------------
	// --- Save the data
	// ------------------------------

	err := suite.database.SaveTotalStake(input)
	suite.Require().NoError(err)

	// ------------------------------
	// --- Verify the data
	// ------------------------------
	expectedRow := dbtypes.NewTotalStakeRow(1, 2)
	var outputs []dbtypes.TotalStakeRow
	err = suite.database.Sqlx.Select(&outputs, `SELECT * FROM total_stake`)
	suite.Require().NoError(err)
	suite.Require().Len(outputs, 1, "should contain only one row")
	suite.Require().True(expectedRow.Equal(outputs[0]))
}

func (suite *DbTestSuite) TestBigDipperDb_StakingTable() {

	// ------------------------------
	// --- Prepare the data
	// ------------------------------

	input := types.NewStakingTable(10, []string{"abc", "efg"})

	// ------------------------------
	// --- Save the data
	// ------------------------------

	err := suite.database.SaveStakingTable(input)
	suite.Require().NoError(err)

	// ------------------------------
	// --- Verify the data
	// ------------------------------
	expectedRow := dbtypes.NewStakingTableRow(10, `["abc","efg"]`)
	var outputs []dbtypes.StakingTableRow
	err = suite.database.Sqlx.Select(&outputs, `SELECT * FROM staking_table`)
	suite.Require().NoError(err)
	suite.Require().Len(outputs, 1, "should contain only one row")
	suite.Require().True(expectedRow.Equal(outputs[0]))

}

func (suite *DbTestSuite) TestBigDipperDb_ProposedTable() {

	// ------------------------------
	// --- Prepare the data
	// ------------------------------

	input := types.NewProposedTable(10, []string{"abc", "efg"})

	// ------------------------------
	// --- Save the data
	// ------------------------------

	err := suite.database.SaveProposedTable(input)
	suite.Require().NoError(err)

	// ------------------------------
	// --- Verify the data
	// ------------------------------
	expectedRow := dbtypes.NewProposedTableRow(10, `["abc","efg"]`)
	var outputs []dbtypes.ProposedTableRow
	err = suite.database.Sqlx.Select(&outputs, `SELECT * FROM proposed_table`)
	suite.Require().NoError(err)
	suite.Require().Len(outputs, 1, "should contain only one row")
	suite.Require().True(expectedRow.Equal(outputs[0]))

}

func (suite *DbTestSuite) TestBigDipperDb_CurrentTable() {

	// ------------------------------
	// --- Prepare the data
	// ------------------------------

	input := types.NewCurrentTable(10, []string{"abc", "efg"})

	// ------------------------------
	// --- Save the data
	// ------------------------------

	err := suite.database.SaveCurrentTable(input)
	suite.Require().NoError(err)

	// ------------------------------
	// --- Verify the data
	// ------------------------------
	expectedRow := dbtypes.NewCurrentTableRow(10, `["abc","efg"]`)
	var outputs []dbtypes.CurrentTableRow
	err = suite.database.Sqlx.Select(&outputs, `SELECT * FROM current_table`)
	suite.Require().NoError(err)
	suite.Require().Len(outputs, 1, "should contain only one row")
	suite.Require().True(expectedRow.Equal(outputs[0]))

}

func (suite *DbTestSuite) TestBigDipperDb_NodeUnstakingTokens() {

	// ------------------------------
	// --- Prepare the data
	// ------------------------------

	/*  TODO: Prepare parameter    */

	input := []types.NodeUnstakingTokens{
		types.NewNodeUnstakingTokens("0x1", 100000008, 1),
	}

	// ------------------------------
	// --- Save the data
	// ------------------------------

	err := suite.database.SaveNodeUnstakingTokens(input)
	suite.Require().NoError(err)

	// ------------------------------
	// --- Verify the data
	// ------------------------------
	expectedRow := dbtypes.NewNodeUnstakingTokensRow("0x1", 100000008, 1)
	var outputs []dbtypes.NodeUnstakingTokensRow
	err = suite.database.Sqlx.Select(&outputs, `SELECT * FROM node_unstaking_tokens`)
	suite.Require().NoError(err)
	suite.Require().Len(outputs, 1, "should contain only one row")
	suite.Require().True(expectedRow.Equal(outputs[0]))

}

func (suite *DbTestSuite) TestBigDipperDb_NodeTotalCommitment() {

	// ------------------------------
	// --- Prepare the data
	// ------------------------------

	/*  TODO: Prepare parameter    */

	input := []types.NodeTotalCommitment{
		types.NewNodeTotalCommitment("0x1", 100000008, 1),
	}

	// ------------------------------
	// --- Save the data
	// ------------------------------

	err := suite.database.SaveNodeTotalCommitment(input)
	suite.Require().NoError(err)

	// ------------------------------
	// --- Verify the data
	// ------------------------------
	expectedRow := dbtypes.NewNodeTotalCommitmentRow("0x1", 100000008, 1)
	var outputs []dbtypes.NodeTotalCommitmentRow
	err = suite.database.Sqlx.Select(&outputs, `SELECT * FROM node_total_commitment`)
	suite.Require().NoError(err)
	suite.Require().Len(outputs, 1, "should contain only one row")
	suite.Require().True(expectedRow.Equal(outputs[0]))

}

func (suite *DbTestSuite) TestBigDipperDb_NodeTotalCommitmentWithoutDelegators() {

	// ------------------------------
	// --- Prepare the data
	// ------------------------------

	/*  TODO: Prepare parameter    */

	input := []types.NodeTotalCommitmentWithoutDelegators{
		types.NewNodeTotalCommitmentWithoutDelegators("0x1", 100000008, 1),
	}

	// ------------------------------
	// --- Save the data
	// ------------------------------

	err := suite.database.SaveNodeTotalCommitmentWithoutDelegators(input)
	suite.Require().NoError(err)

	// ------------------------------
	// --- Verify the data
	// ------------------------------
	expectedRow := dbtypes.NewNodeTotalCommitmentWithoutDelegatorsRow("0x1", 100000008, 1)
	var outputs []dbtypes.NodeTotalCommitmentWithoutDelegatorsRow
	err = suite.database.Sqlx.Select(&outputs, `SELECT * FROM node_total_commitment_without_delegators`)
	suite.Require().NoError(err)
	suite.Require().Len(outputs, 1, "should contain only one row")
	suite.Require().True(expectedRow.Equal(outputs[0]))

}

func (suite *DbTestSuite) TestBigDipperDb_NodeStakingKey() {

	// ------------------------------
	// --- Prepare the data
	// ------------------------------

	/*  TODO: Prepare parameter    */

	input := []types.NodeStakingKey{
		types.NewNodeStakingKey("0x1", "0x2", 1),
	}

	// ------------------------------
	// --- Save the data
	// ------------------------------

	err := suite.database.SaveNodeStakingKey(input)
	suite.Require().NoError(err)

	// ------------------------------
	// --- Verify the data
	// ------------------------------
	expectedRow := dbtypes.NewNodeStakingKeyRow("0x1", "0x2", 1)
	var outputs []dbtypes.NodeStakingKeyRow
	err = suite.database.Sqlx.Select(&outputs, `SELECT * FROM node_staking_key`)
	suite.Require().NoError(err)
	suite.Require().Len(outputs, 1, "should contain only one row")
	suite.Require().True(expectedRow.Equal(outputs[0]))

}

func (suite *DbTestSuite) TestBigDipperDb_NodeStakedTokens() {

	// ------------------------------
	// --- Prepare the data
	// ------------------------------

	/*  TODO: Prepare parameter    */

	input := []types.NodeStakedTokens{
		types.NewNodeStakedTokens("0x1", 2, 1),
	}

	// ------------------------------
	// --- Save the data
	// ------------------------------

	err := suite.database.SaveNodeStakedTokens(input)
	suite.Require().NoError(err)

	// ------------------------------
	// --- Verify the data
	// ------------------------------
	expectedRow := dbtypes.NewNodeStakedTokensRow("0x1", 2, 1)
	var outputs []dbtypes.NodeStakedTokensRow
	err = suite.database.Sqlx.Select(&outputs, `SELECT * FROM node_staked_tokens`)
	suite.Require().NoError(err)
	suite.Require().Len(outputs, 1, "should contain only one row")
	suite.Require().True(expectedRow.Equal(outputs[0]))

}

func (suite *DbTestSuite) TestBigDipperDb_NodeRole() {

	// ------------------------------
	// --- Prepare the data
	// ------------------------------

	/*  TODO: Prepare parameter    */

	input := []types.NodeRole{
		types.NewNodeRole("0x1", 2, 1),
	}

	// ------------------------------
	// --- Save the data
	// ------------------------------

	err := suite.database.SaveNodeRole(input)
	suite.Require().NoError(err)

	// ------------------------------
	// --- Verify the data
	// ------------------------------
	expectedRow := dbtypes.NewNodeRoleRow("0x1", 2, 1)
	var outputs []dbtypes.NodeRoleRow
	err = suite.database.Sqlx.Select(&outputs, `SELECT * FROM node_role`)
	suite.Require().NoError(err)
	suite.Require().Len(outputs, 1, "should contain only one row")
	suite.Require().True(expectedRow.Equal(outputs[0]))

}

func (suite *DbTestSuite) TestBigDipperDb_NodeRewardedTokens() {

	// ------------------------------
	// --- Prepare the data
	// ------------------------------

	/*  TODO: Prepare parameter    */

	input := []types.NodeRewardedTokens{
		types.NewNodeRewardedTokens("0x1", 2, 1),
	}

	// ------------------------------
	// --- Save the data
	// ------------------------------

	err := suite.database.SaveNodeRewardedTokens(input)
	suite.Require().NoError(err)

	// ------------------------------
	// --- Verify the data
	// ------------------------------
	expectedRow := dbtypes.NewNodeRewardedTokensRow("0x1", 2, 1)
	var outputs []dbtypes.NodeRewardedTokensRow
	err = suite.database.Sqlx.Select(&outputs, `SELECT * FROM node_rewarded_tokens`)
	suite.Require().NoError(err)
	suite.Require().Len(outputs, 1, "should contain only one row")
	suite.Require().True(expectedRow.Equal(outputs[0]))

}

func (suite *DbTestSuite) TestBigDipperDb_NodeNetworkingKey() {

	// ------------------------------
	// --- Prepare the data
	// ------------------------------

	/*  TODO: Prepare parameter    */

	input := []types.NodeNetworkingKey{
		types.NewNodeNetworkingKey("0x1", "0x2", 1),
	}

	// ------------------------------
	// --- Save the data
	// ------------------------------

	err := suite.database.SaveNodeNetworkingKey(input)
	suite.Require().NoError(err)

	// ------------------------------
	// --- Verify the data
	// ------------------------------
	expectedRow := dbtypes.NewNodeNetworkingKeyRow("0x1", "0x2", 1)
	var outputs []dbtypes.NodeNetworkingKeyRow
	err = suite.database.Sqlx.Select(&outputs, `SELECT * FROM node_networking_key`)
	suite.Require().NoError(err)
	suite.Require().Len(outputs, 1, "should contain only one row")
	suite.Require().True(expectedRow.Equal(outputs[0]))

}

func (suite *DbTestSuite) TestBigDipperDb_NodeNetworkingAddress() {

	// ------------------------------
	// --- Prepare the data
	// ------------------------------

	/*  TODO: Prepare parameter    */

	input := []types.NodeNetworkingAddress{
		types.NewNodeNetworkingAddress("0x1", "0x2", 1),
	}

	// ------------------------------
	// --- Save the data
	// ------------------------------

	err := suite.database.SaveNodeNetworkingAddress(input)
	suite.Require().NoError(err)

	// ------------------------------
	// --- Verify the data
	// ------------------------------
	expectedRow := dbtypes.NewNodeNetworkingAddressRow("0x1", "0x2", 1)
	var outputs []dbtypes.NodeNetworkingAddressRow
	err = suite.database.Sqlx.Select(&outputs, `SELECT * FROM node_networking_address`)
	suite.Require().NoError(err)
	suite.Require().Len(outputs, 1, "should contain only one row")
	suite.Require().True(expectedRow.Equal(outputs[0]))

}

func (suite *DbTestSuite) TestBigDipperDb_NodeInitialWeight() {

	// ------------------------------
	// --- Prepare the data
	// ------------------------------

	/*  TODO: Prepare parameter    */

	input := []types.NodeInitialWeight{
		types.NewNodeInitialWeight("0x1", 2, 1),
	}

	// ------------------------------
	// --- Save the data
	// ------------------------------

	err := suite.database.SaveNodeInitialWeight(input)
	suite.Require().NoError(err)

	// ------------------------------
	// --- Verify the data
	// ------------------------------
	expectedRow := dbtypes.NewNodeInitialWeightRow("0x1", 2, 1)
	var outputs []dbtypes.NodeInitialWeightRow
	err = suite.database.Sqlx.Select(&outputs, `SELECT * FROM node_initial_weight`)
	suite.Require().NoError(err)
	suite.Require().Len(outputs, 1, "should contain only one row")
	suite.Require().True(expectedRow.Equal(outputs[0]))

}

func (suite *DbTestSuite) TestBigDipperDb_NodeInfoFromAddresses() {

	// ------------------------------
	// --- Prepare the data
	// ------------------------------

	/*  TODO: Prepare parameter    */

	id := "e7df1454826425251716a703e907981672a43208ef3eabfc95d593673da778f6"

	stakerNodeInfo := types.NewStakerNodeInfo(id, 5, "34.211.45.12:3569",
		"3fa19db960a86a1722a2d8ffa9563bd1e7d905c91536860c64f9e808ef88862639112a1e872d7eef93dd91207d07dd7c043e2a1e80077b109290682250429f1f",
		"87d827de3e1b3541c394dcbbb6d76d98e3a7710d6740d28122c468b83f41002625e7a5788cabfcd6ce76b188f7f60de614364d4ab2932dfe0ed6f2d602bd551606ea31045ca2ccde9658a175ccd73da859ab17e56ad81ca4f6ef982c5968a7cb",
		0, 20000000, 0, 0, 0, []uint32{}, 0, 0, 0)

	input := []types.NodeInfoFromAddress{
		types.NewNodeInfoFromAddress("0x1", stakerNodeInfo, 1),
	}

	// ------------------------------
	// --- Save the data
	// ------------------------------

	err := suite.database.SaveNodeInfoFromAddresses(input)
	suite.Require().NoError(err)

	// ------------------------------
	// --- Verify the data
	// ------------------------------
	expectedStakerNodeInfo := `{"Id":"e7df1454826425251716a703e907981672a43208ef3eabfc95d593673da778f6","Role":5,"NetworkingAddress":"34.211.45.12:3569","NetworkingKey":"3fa19db960a86a1722a2d8ffa9563bd1e7d905c91536860c64f9e808ef88862639112a1e872d7eef93dd91207d07dd7c043e2a1e80077b109290682250429f1f","StakingKey":"87d827de3e1b3541c394dcbbb6d76d98e3a7710d6740d28122c468b83f41002625e7a5788cabfcd6ce76b188f7f60de614364d4ab2932dfe0ed6f2d602bd551606ea31045ca2ccde9658a175ccd73da859ab17e56ad81ca4f6ef982c5968a7cb","TokensStaked":0,"TokensCommitted":20000000,"TokensUnstaking":0,"TokensUnstaked":0,"TokensRewarded":0,"Delegators":[],"DelegatorIDCounter":0,"TokensRequestedToUnstake":0,"InitialWeight":0}`

	expectedRow := dbtypes.NewNodeInfoFromAddressRow("0x1", expectedStakerNodeInfo, 1)
	var outputs []dbtypes.NodeInfoFromAddressRow
	err = suite.database.Sqlx.Select(&outputs, `SELECT * FROM node_info_from_address`)
	suite.Require().NoError(err)
	suite.Require().Len(outputs, 1, "should contain only one row")
	suite.Require().True(expectedRow.Equal(outputs[0]))
}

func (suite *DbTestSuite) TestBigDipperDb_NodeInfoFromNodeID() {

	// ------------------------------
	// --- Prepare the data
	// ------------------------------

	/*  TODO: Prepare parameter    */

	id := "e7df1454826425251716a703e907981672a43208ef3eabfc95d593673da778f6"

	stakerNodeInfo := types.NewStakerNodeInfo(id, 5, "34.211.45.12:3569",
		"3fa19db960a86a1722a2d8ffa9563bd1e7d905c91536860c64f9e808ef88862639112a1e872d7eef93dd91207d07dd7c043e2a1e80077b109290682250429f1f",
		"87d827de3e1b3541c394dcbbb6d76d98e3a7710d6740d28122c468b83f41002625e7a5788cabfcd6ce76b188f7f60de614364d4ab2932dfe0ed6f2d602bd551606ea31045ca2ccde9658a175ccd73da859ab17e56ad81ca4f6ef982c5968a7cb",
		0, 20000000, 0, 0, 0, []uint32{}, 0, 0, 0)

	input := []types.NodeInfoFromNodeID{
		types.NewNodeInfoFromNodeID("0x1", stakerNodeInfo, 1),
	}

	// ------------------------------
	// --- Save the data
	// ------------------------------

	err := suite.database.SaveNodeInfoFromNodeIDs(input)
	suite.Require().NoError(err)

	// ------------------------------
	// --- Verify the data
	// ------------------------------
	expectedStakerNodeInfo := `{"Id":"e7df1454826425251716a703e907981672a43208ef3eabfc95d593673da778f6","Role":5,"NetworkingAddress":"34.211.45.12:3569","NetworkingKey":"3fa19db960a86a1722a2d8ffa9563bd1e7d905c91536860c64f9e808ef88862639112a1e872d7eef93dd91207d07dd7c043e2a1e80077b109290682250429f1f","StakingKey":"87d827de3e1b3541c394dcbbb6d76d98e3a7710d6740d28122c468b83f41002625e7a5788cabfcd6ce76b188f7f60de614364d4ab2932dfe0ed6f2d602bd551606ea31045ca2ccde9658a175ccd73da859ab17e56ad81ca4f6ef982c5968a7cb","TokensStaked":0,"TokensCommitted":20000000,"TokensUnstaking":0,"TokensUnstaked":0,"TokensRewarded":0,"Delegators":[],"DelegatorIDCounter":0,"TokensRequestedToUnstake":0,"InitialWeight":0}`

	expectedRow := dbtypes.NewNodeInfoFromNodeIDRow("0x1", expectedStakerNodeInfo, 1)
	var outputs []dbtypes.NodeInfoFromNodeIDRow
	err = suite.database.Sqlx.Select(&outputs, `SELECT * FROM node_info_from_node_id`)
	suite.Require().NoError(err)
	suite.Require().Len(outputs, 1, "should contain only one row")
	suite.Require().True(expectedRow.Equal(outputs[0]))
}

func (suite *DbTestSuite) TestBigDipperDb_NodeCommittedTokens() {

	// ------------------------------
	// --- Prepare the data
	// ------------------------------

	/*  TODO: Prepare parameter    */

	input := []types.NodeCommittedTokens{
		types.NewNodeCommittedTokens("0x1", 2, 1),
	}

	// ------------------------------
	// --- Save the data
	// ------------------------------

	err := suite.database.SaveNodeCommittedTokens(input)
	suite.Require().NoError(err)

	// ------------------------------
	// --- Verify the data
	// ------------------------------
	expectedRow := dbtypes.NewNodeCommittedTokensRow("0x1", 2, 1)
	var outputs []dbtypes.NodeCommittedTokensRow
	err = suite.database.Sqlx.Select(&outputs, `SELECT * FROM node_committed_tokens`)
	suite.Require().NoError(err)
	suite.Require().Len(outputs, 1, "should contain only one row")
	suite.Require().True(expectedRow.Equal(outputs[0]))

}
