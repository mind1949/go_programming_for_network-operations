package main

import (
	"io/ioutil"
	"os"
	"text/template"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Base Base `yaml:"base"`
	Bgp  Bgp  `yaml:"bgp"`
	Ntp  Ntp  `yaml:"ntp"`
}

type Base struct {
	Hostname string `yaml:"hostname"`
}

type Bgp struct {
	As        string     `yaml:"as"`
	Id        string     `yaml:"id"`
	Neighbors []Neighbor `yaml:"neighbors"`
}

type Neighbor struct {
	Ip string `yaml:"ip"`
	As string `yaml:"as"`
}

type Ntp struct {
	Source    string `yaml:"source"`
	Primary   string `yaml:"primary"`
	Secondary string `yaml:"secondary"`
}

func main() {
	var config Config
	data, _ := ioutil.ReadFile("file.yaml")
	_ = yaml.Unmarshal(data, &config)
	tmpl := make(map[string]*template.Template)
	files := template.Must(template.ParseFiles("base.tmpl", "bgp.tmpl", "ntp.tmpl"))
	tmpl["base.tmpl"] = files
	_ = tmpl["base.tmpl"].ExecuteTemplate(os.Stdout, "base", config)
}
