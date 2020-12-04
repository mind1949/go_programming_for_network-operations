package main

import (
	"fmt"
	"net"
)

func main() {
	netw := "1.1.1.0/30"
	ip, ipnet, _ := net.ParseCIDR(netw)

	inc := func(ip net.IP) {
		for i := len(ip) - 1; i >= 0; i-- {
			ip[i]++
			if ip[i] >= 1 {
				break
			}
		}
	}

	ipmask := ip.Mask(ipnet.Mask)
	for ip := ipmask; ipnet.Contains(ip); inc(ip) {
		fmt.Println(ip)
	}

}
