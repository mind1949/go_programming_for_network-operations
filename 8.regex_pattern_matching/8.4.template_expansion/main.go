package main

import (
	"fmt"
	"regexp"
)

func main() {
	src := `
host1.domain.tld some text
host2.domain.tld some text
`
	re := regexp.MustCompile(`(?m)(?P<thir>\w+).(?P<sec>\w+).(?P<fir>\w+)\s(?P<text>.*)$`)
	tpl := "$thir.$sec.$fir\tIN\tTXT\t\"$text\"\n"

	dst := []byte{}
	findidx := re.FindAllStringSubmatchIndex
	for _, matches := range findidx(src, -1) {
		dst = re.ExpandString(dst, tpl, src, matches)
	}
	fmt.Println(string(dst))
}
