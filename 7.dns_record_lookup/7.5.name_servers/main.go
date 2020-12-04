package main

import (
	"fmt"
	"net"
)

func main() {
	nss, _ := net.LookupNS("google.cn")
	for _, ns := range nss {
		fmt.Println(ns.Host)
	}
}
