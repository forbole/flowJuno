package utils

import (
	"testing"

	"github.com/stretchr/testify/suite"

	testutils "github.com/forbole/flowJuno/modules/utils"
	_ "github.com/proullon/ramsql/driver"
)

func TestAuthProxyTestSuite(t *testing.T) {
	suite.Run(t, new(AuthProxyTestSuite))
}

// AuthProxyTestSuite the base test case class for Auth module
type AuthProxyTestSuite struct {
	testutils.ProxyTestSuite
}
