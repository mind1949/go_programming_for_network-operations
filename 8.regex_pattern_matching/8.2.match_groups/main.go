package main

import (
	"fmt"
	"regexp"
)

func main() {
	data := `#host.domain.tld#`
	pattern := regexp.MustCompile(`(\w+).(\w+).(\w+)`)
	groups := pattern.FindAllStringSubmatch(data, -1)
	fmt.Printf("%q\n", groups)
	fmt.Printf("groups[0][0]: %q\n", groups[0][0])
	fmt.Printf("groups[0][1]: %q\n", groups[0][1])
	fmt.Printf("groups[0][2]: %q\n", groups[0][2])
	fmt.Printf("groups[0][3]: %q\n", groups[0][3])
}
