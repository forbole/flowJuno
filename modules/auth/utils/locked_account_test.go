package utils_test



func (suite *ProxyTestSuite)TestProxy_GetLockedAccountAddress(){
	proxy:=suite.Proxy
	GetLockedAccount(proxy.Client())
}