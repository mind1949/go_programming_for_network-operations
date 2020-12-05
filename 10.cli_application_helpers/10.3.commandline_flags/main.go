package main

import (
	"flag"
	"fmt"
)

func main() {
	port := flag.Int("port", 8000, "specify port")
	enable := flag.Bool("enable", false, "specify enable, default is false")
	name := flag.String("n", "blank", "specify name")
	flag.Parse()

	fmt.Println("port=", *port)
	fmt.Println("enable=", *enable)
	fmt.Println("name=", *name)
}
