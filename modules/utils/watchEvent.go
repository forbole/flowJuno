package utils

import (
	"fmt"

	"github.com/forbole/flowJuno/client"
	flowclient "github.com/onflow/flow-go-sdk/client"
)

func CheckRewardPaidEvent(flowClient client.Proxy) (bool, error) {
	height, err := flowClient.LatestHeight()
	if err != nil {
		return false, err
	}

	// With this to indicate how close to the end of epoch
	events, err := flowClient.Client().GetEventsForHeightRange(flowClient.Ctx(),
		flowclient.EventRangeQuery{
			Type:        fmt.Sprintf("A.%s.FlowIDTableStaking.RewardsPaid", flowClient.Contract().StakingTable[2:]),
			StartHeight: uint64(height-249),
			EndHeight:   uint64(height),
		})
	if err != nil {
		return false, err
	}

	for _, block := range events {
		if len(block.Events) > 0 {
			return true, nil
		}
	}

	return false, nil
}
