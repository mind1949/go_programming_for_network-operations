package main

import (
	"fmt"
	"net"
)

func main() {
	netw := "1.1.1.1/27"
	ipv4IP, ipv4Net, _ := net.ParseCIDR(netw)

	m := ipv4Net.Mask
	dotmask := fmt.Sprintf("%d.%d.%d.%d", m[0], m[1], m[2], m[3])
	fmt.Printf("dot-decimal notation: %s %s\n", ipv4IP, dotmask)

	cidrmask := net.IPMask(net.ParseIP(dotmask).To4())
	length, _ := cidrmask.Size()
	slash := fmt.Sprintf("%s/%d", ipv4IP, length)
	fmt.Println("CIDR notation: ", slash)
}
