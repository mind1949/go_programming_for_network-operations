package main

import (
	"encoding/json"
	"fmt"
)

type Inventory struct {
	Total   int      `json:"total"`
	Devices []string `json:"devices"`
}

func main() {
	inventory := Inventory{
		Total:   3,
		Devices: []string{"SW1", "SW2", "SW3"},
	}

	data, _ := json.MarshalIndent(inventory, "", " ")

	fmt.Println(string(data))
}
