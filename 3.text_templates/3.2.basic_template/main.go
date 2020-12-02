package main

import (
	"encoding/csv"
	"io"
	"os"
	"text/template"
)

type Vlans struct {
	Id, Name string
}

func main() {
	data, _ := os.Open("file.csv")
	reader := csv.NewReader(data)

	var vlans []Vlans
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}

		vlans = append(vlans, Vlans{
			Id:   row[0],
			Name: row[1],
		})
	}

	tmpl, _ := template.ParseFiles("file.tmpl")
	_ = tmpl.Execute(os.Stdout, vlans)
}
