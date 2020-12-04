package main

import (
	"fmt"
	"net"
)

func main() {
	mxs, _ := net.LookupMX("google.cn")
	for _, mx := range mxs {
		fmt.Println(mx.Host, mx.Pref)
	}
}
