package main

import (
	"encoding/xml"
	"io/ioutil"
)

type Campus struct {
	XMLName  xml.Name `xml:"campus"`
	Name     string   `xml:"name,attr"`
	Comment  string   `xml:",comment"`
	Building Building
}

type Building struct {
	XMLName xml.Name `xml:"building"`
	Name    string   `xml:"name,attr"`
	Comment string   `xml:",comment"`
	Device  Device
}

type Device struct {
	XMLName xml.Name `xml:"device"`
	Name    string   `xml:"name,attr"`
	Comment string   `xml:",comment"`
	Type    string   `xml:"type,attr"`
}

func main() {
	campus := Campus{
		Name:    "campus",
		Comment: "building-comment",
		Building: Building{
			Name:    "blg1",
			Comment: "device-comment",
			Device: Device{
				Name: "rtr1",
				Type: "router",
			},
		},
	}

	data, _ := xml.MarshalIndent(campus, " ", "")
	_ = ioutil.WriteFile("file.xml", data, 0644)

}
