package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	config := `
	desc v1
!
int Vlan2
	desc2
`
	a := regexp.MustCompile("(?m)(\n^!$\n)")
	m := a.Split(config, -1)
	arr := []string{}
	for _, s := range m {
		s := strings.TrimSuffix(s, "\n")
		arr = append(arr, s)
	}

	fmt.Printf("%q\n", arr)
}
