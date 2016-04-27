package ipCommons

import (
	"fmt"
	"net"
	"testing"
)

func TestUint32ToIP(t *testing.T) {
	fmt.Printf("%#v\n", Uint32ToIp4(0))
	fmt.Printf("%#v\n", Uint32ToIp4(167903491))
	fmt.Println(Ip4ToUint32(net.IP{0, 0, 0, 3}))
	fmt.Println(Ip4ToUint32(net.IP{10, 2, 1, 3}))
	fmt.Println(Ip4StringToUint32("10.2.1.3"))
	fmt.Println(Uint32ToIp4String(167903491))
	fmt.Println(Ip4ToUint32(net.IP{255, 255, 255, 255}))
	fmt.Println(Ip4ToUint32(net.IP{255, 255, 255, 255}))
	fmt.Println(IpRangeToIpStringArray("1.1.1.1-1.1.2.3"))
}
