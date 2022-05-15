package ip_test

import (
	"github.com/ilijamt/blacklist_checker/internal/ip"
	"github.com/stretchr/testify/assert"
	"net"
	"testing"
)

func TestIsPrivateIP(t *testing.T) {
	assert.True(t, ip.IsPrivate(net.ParseIP("127.0.0.1")), "should be private")
	assert.True(t, ip.IsPrivate(net.ParseIP("192.168.254.254")), "should be private")
	assert.True(t, ip.IsPrivate(net.ParseIP("10.255.0.3")), "should be private")
	assert.True(t, ip.IsPrivate(net.ParseIP("172.16.255.255")), "should be private")
	assert.True(t, ip.IsPrivate(net.ParseIP("172.31.255.255")), "should be private")
	assert.False(t, ip.IsPrivate(net.ParseIP("128.0.0.1")), "should not be private")
	assert.False(t, ip.IsPrivate(net.ParseIP("192.169.255.255")), "should not be private")
	assert.False(t, ip.IsPrivate(net.ParseIP("9.255.0.255")), "should not be private")
	assert.False(t, ip.IsPrivate(net.ParseIP("172.32.255.255")), "should not be private")

	assert.True(t, ip.IsPrivate(net.ParseIP("::0")), "should be private")
	assert.True(t, ip.IsPrivate(net.ParseIP("::1")), "should be private")
	assert.False(t, ip.IsPrivate(net.ParseIP("::2")), "should not be private")

	assert.True(t, ip.IsPrivate(net.ParseIP("fe80::1")), "should be private")
	assert.True(t, ip.IsPrivate(net.ParseIP("febf::1")), "should be private")
	assert.False(t, ip.IsPrivate(net.ParseIP("fec0::1")), "should not be private")
	assert.False(t, ip.IsPrivate(net.ParseIP("feff::1")), "should not be private")

	assert.True(t, ip.IsPrivate(net.ParseIP("ff00::1")), "should be private")
	assert.True(t, ip.IsPrivate(net.ParseIP("ff10::1")), "should be private")
	assert.True(t, ip.IsPrivate(net.ParseIP("ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff")), "should be private")

	assert.True(t, ip.IsPrivate(net.ParseIP("2002::")), "should be private")
	assert.True(t, ip.IsPrivate(net.ParseIP("2002:ffff:ffff:ffff:ffff:ffff:ffff:ffff")), "should be private")
	assert.True(t, ip.IsPrivate(net.ParseIP("0100::")), "should be private")
	assert.True(t, ip.IsPrivate(net.ParseIP("0100::0000:ffff:ffff:ffff:ffff")), "should be private")
	assert.False(t, ip.IsPrivate(net.ParseIP("0100::0001:0000:0000:0000:0000")), "should be private")
}
