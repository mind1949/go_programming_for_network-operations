package main

import (
	"fmt"
	"regexp"
)

func main() {
	dat := `#host.domain.tld#`
	pat := regexp.MustCompile(`(?P<third>(\w+)).(?P<sec>(\w+)).(?P<first>(\w+))`)
	mch := pat.FindStringSubmatch(dat)

	res := make(map[string]string)
	for i, name := range pat.SubexpNames() {
		if i != 0 && name != "" {
			res[name] = mch[i]
		}
	}

	ff := fmt.Printf
	ff("\n%q\n", mch)
	ff("res[\"first\"]: %s\n", res["first"])
	ff("res[\"sec\"]: %s\n", res["sec"])
	ff("res[\"third\"]: %s\n", res["third"])
}
