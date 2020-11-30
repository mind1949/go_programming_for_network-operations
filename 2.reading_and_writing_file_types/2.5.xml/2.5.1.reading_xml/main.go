package main

import (
	"encoding/xml"
	"fmt"
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
	data, _ := ioutil.ReadFile("file.xml")
	campus := Campus{}
	_ = xml.Unmarshal(data, &campus)
	fl := fmt.Println
	fl("campus name: ", campus.Name)
	fl("Building name: ", campus.Building.Name)
	fl("Building comment: ", campus.Building.Comment)
	fl("Device name: ", campus.Building.Device.Name)
	fl("Device comment: ", campus.Building.Device.Comment)
	fl("Device type: ", campus.Building.Device.Type)
}
