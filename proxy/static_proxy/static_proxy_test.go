package static_proxy

import "testing"

func TestStaticProxy(t *testing.T) {
	h := &HouseRenter{}
	proxy := &Proxy{h}
	proxy.rentHouse()
}
