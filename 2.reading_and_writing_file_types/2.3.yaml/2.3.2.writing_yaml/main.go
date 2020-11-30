package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Host  string `yaml:"host"`
	Ports []int  `yaml:"ports"`
}

func main() {
	config := Config{
		Host:  "localhost",
		Ports: []int{80, 443},
	}
	data, _ := yaml.Marshal(config)
	_ = ioutil.WriteFile("file.yaml", data, 0644)
}
