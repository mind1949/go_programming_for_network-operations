package main

import (
	"encoding/csv"
	"os"
)

func main() {
	lines := [][]string{
		{"HOST", "IP ADDR", "LOCATION"},
		{"RTR1", "1.1.1.1", "Seattle WA"},
		{"RTR2", "2.2.2.2", "Nevada NV"},
	}

	file, _ := os.Create("file.csv")
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()
	for _, line := range lines {
		_ = writer.Write(line)
	}
}
