package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func main() {
	cmd := exec.Command("cat", "r.txt", "w.txt")
	stdoutpipe, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	stderrpipe, err := cmd.StderrPipe()
	if err != nil {
		panic(err)
	}
	cmd.Start()
	stdout, _ := ioutil.ReadAll(stdoutpipe)
	stderr, _ := ioutil.ReadAll(stderrpipe)
	fmt.Printf("stdout: %s\nstderr: %s\n", stdout, stderr)
	cmd.Wait()
}
