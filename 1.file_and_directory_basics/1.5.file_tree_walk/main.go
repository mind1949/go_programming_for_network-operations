package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	scan := func(path string, info os.FileInfo, _ error) error {
		fmt.Println(info.IsDir(), path)
		return nil
	}

	_ = filepath.Walk(".", scan)
}
