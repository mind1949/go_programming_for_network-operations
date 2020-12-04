package main

import (
	"fmt"
	"net"
)

func main() {
	hosts, _ := net.LookupAddr("8.8.8.8")
	for _, host := range hosts {
		fmt.Println(host)
	}
}
