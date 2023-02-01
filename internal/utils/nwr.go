package utils

import "net"

func IpInCidr(cidr, ip string) bool {
	_, ipnetA, err := net.ParseCIDR(cidr)
	if err != nil {
		return false
	}

	ipB := net.ParseIP(ip)

	return ipnetA.Contains(ipB)
}
