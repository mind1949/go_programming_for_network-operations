package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Host  string `yaml:"host"`
	Ports string `yaml:"ports"`
}

func main() {
	data, _ := ioutil.ReadFile("file.yaml")
	config := Config{}
	_ = yaml.Unmarshal(data, &config)
	fmt.Printf("%s: %v\n", config.Host, config.Ports)
}
