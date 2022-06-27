package utils

import (
	"fmt"
	"strings"

	"github.com/forbole/flowJuno/client"
	flowclient "github.com/onflow/flow-go-sdk/client"
	"github.com/rs/zerolog/log"
)

func CheckRewardPaidEvent(flowClient client.Proxy) (bool, error) {
	log.Trace().Str("module", "util").Msg("CheckRewardPaidEvent")

	h, err := flowClient.LatestHeight()
	if err != nil {
		return false, err
	}
	
	height:=uint64(h)

	// With this to indicate how close to the end of epoch
	t:=fmt.Sprintf("A.%s.FlowIDTableStaking.RewardsPaid", flowClient.Contract().StakingTable[2:])
		blockevent,err:=flowClient.Client().GetEventsForHeightRange(flowClient.Ctx(),flowclient.EventRangeQuery{
			Type: t,
			StartHeight: 	height-249,
			EndHeight: height,
		})
		if err!=nil{
			return false,err
		}
			for _,b:=range blockevent{
				for _,e:=range b.Events{
					fmt.Println(e.String())
					if strings.Contains(e.Type,t){
						log.Debug().Str("module", "util").Msg("CheckRewardPaidEvent is true")
						return true,nil
					}
				}
			}
			//fmt.Println("GetEventForHeight")

	return false, nil
}
