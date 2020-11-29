package main

import (
	"fmt"
	"os"
)

func main() {
	file, _ := os.Stat("main.go")

	fl := fmt.Println
	fl("File name: ", file.Name())
	fl("Size in byte: ", file.Size())
	fl("Last modified: ", file.ModTime())
	fl("Is a directory: ", file.IsDir())

	ff := fmt.Printf
	ff("Permission 9-bit: %s\n", file.Mode())
	ff("Permission 3-digit: %o\n", file.Mode())
	ff("Permission 4-digit: %04o\n", file.Mode())
}
