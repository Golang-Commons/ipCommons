package ipCommons

import (
	"encoding/binary"
	"errors"
	"net"
	"strings"
)

// Convert an ipv4 to a uint32
func Ip4ToUint32(ip net.IP) (uint32, error) {
	ip = ip.To4()
	if ip == nil {
		return 0, errors.New("not able to convert ip to ipv4")
	}
	return binary.BigEndian.Uint32([]byte(ip)), nil
}

func Ip4StringToUint32(ipString string) (uint32, error) {
	ip := net.ParseIP(ipString)
	if ip == nil {
		return 0, errors.New("parse ip string error " + ipString)
	}
	return Ip4ToUint32(ip)
}

func Uint32ToIp4String(ip uint32) string {
	return Uint32ToIp4(ip).String()
}

//convert a uint32 to an ipv4
func Uint32ToIp4(ip uint32) net.IP {
	addr := net.IP{0, 0, 0, 0}
	binary.BigEndian.PutUint32(addr, ip)
	return addr
}

//from "1.1.1.1-1.1.2.3" to uint32,uint32,error
func ipRangeToIpStartEndString(ipRangeString string) (a1 uint32, a2 uint32, err error) {
	ipRange := strings.Split(ipRangeString, "-")
	if len(ipRange) != 2 {
		err = errors.New("wrong ip range")
		return
	}
	a1, err = Ip4StringToUint32(ipRange[0])
	if err != nil {
		err = errors.New("wrong start ip")
		return
	}
	a2, err = Ip4StringToUint32(ipRange[1])
	if err != nil {
		err = errors.New("wrong end ip")
		return
	}
	return
}

func IpRangeToIpArray(ipRangeString string) ([]net.IP, error) {
	a1, a2, err := ipRangeToIpStartEndString(ipRangeString)
	if err != nil {
		return nil, err
	}
	var ipUint []net.IP
	for i := a1; i <= a2; i++ {
		ipUint = append(ipUint, Uint32ToIp4(i))
	}
	return ipUint, nil
}

func IpRangeToIpStringArray(ipRangeString string) ([]string, error) {
	a1, a2, err := ipRangeToIpStartEndString(ipRangeString)
	if err != nil {
		return nil, err
	}
	var ipUint []string
	for i := a1; i <= a2; i++ {
		ipUint = append(ipUint, Uint32ToIp4String(i))
	}
	return ipUint, nil
}
