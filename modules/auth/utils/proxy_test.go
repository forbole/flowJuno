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

type AuthProxyTestSuite struct {
	testutils.ProxyTestSuite
}