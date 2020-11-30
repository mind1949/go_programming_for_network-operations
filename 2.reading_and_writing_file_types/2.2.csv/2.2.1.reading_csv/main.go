package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

type Row struct {
	ColA, ColB, ColC string
}

func main() {
	file, _ := os.Open("file.csv")

	var rows []Row

	reader := csv.NewReader(file)
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}

		rows = append(rows, Row{row[0], row[1], row[2]})
	}

	for _, row := range rows {
		fmt.Printf("%s--%s--%s\n", row.ColA, row.ColB, row.ColC)
	}
}
