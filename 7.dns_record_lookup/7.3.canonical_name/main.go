package main

import (
	"fmt"
	"net"
)

func main() {
	cname, _ := net.LookupCNAME("research.swtch.com")
	fmt.Println(cname)
}
