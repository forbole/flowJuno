package utils

import (
	"testing"

	"github.com/stretchr/testify/suite"

	testutils "github.com/forbole/flowJuno/modules/utils"
	_ "github.com/proullon/ramsql/driver"
)

func TestConsensusProxyTestSuite(t *testing.T) {
	suite.Run(t, new(ConsensusProxyTestSuite))
}

type ConsensusProxyTestSuite struct {
	testutils.DbProxyTestSuite
}
