package main

import (
	"os"
	"path/filepath"
)

func main() {
	dname := "./dir"
	fname := "file.txt"
	_ = os.MkdirAll(dname, 0777)
	fpath := filepath.Join(dname, fname)
	file, _ := os.Create(fpath)
	file.Close()
	_ = os.Remove(fpath)
	_ = os.RemoveAll(fpath)
}
