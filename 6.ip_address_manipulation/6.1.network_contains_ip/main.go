package main

import (
	"fmt"
	"net"
)

func main() {
	netw := "1.1.1.0/24"
	ips := []string{"1.1.1.1", "2.2.2.2"}
	_, cidrnet, _ := net.ParseCIDR(netw)
	for _, ip := range ips {
		addr := net.ParseIP(ip)
		result := cidrnet.Contains(addr)
		fmt.Printf("%s contains %s: %t\n", netw, ip, result)
	}
}
