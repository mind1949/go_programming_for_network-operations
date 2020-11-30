package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Inventory struct {
	Total   int      `json:"total"`
	Devices []string `json:"devices"`
}

func main() {
	data, _ := ioutil.ReadFile("file.json")
	inventory := Inventory{}
	_ = json.Unmarshal(data, &inventory)
	fmt.Printf("Total: %d\nDevices: %v\n", inventory.Total, inventory.Devices)
}
