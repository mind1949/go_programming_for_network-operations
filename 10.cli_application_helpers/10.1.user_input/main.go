package main

import (
	"bufio"
	"fmt"
	"os"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter username:")
	user, _ := reader.ReadString('\n')
	fmt.Print("Enter password:")
	pass, _ := terminal.ReadPassword(int(syscall.Stdin))
	fmt.Printf("\n\nUsername: %sPassword: %s\n", user, pass)
}
