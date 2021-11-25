package token

import (
	"testing"

	"github.com/stretchr/testify/suite"

	testutils "github.com/forbole/flowJuno/modules/utils"
	_ "github.com/proullon/ramsql/driver"
)

func TestTokenProxyTestSuite(t *testing.T) {
	suite.Run(t, new(TokenProxyTestSuite))
}

type TokenProxyTestSuite struct {
	testutils.ProxyTestSuite
}