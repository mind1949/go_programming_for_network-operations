package main

import (
	"fmt"
	"net"
)

func main() {
	cname, srvs, _ := net.LookupSRV("xmpp-server", "tcp", "google.com")
	fmt.Printf("cname: %s\n\n", cname)
	for _, srv := range srvs {
		fmt.Printf("%v:%v:%d:%d\n",
			srv.Target, srv.Port, srv.Priority, srv.Weight)
	}
}
