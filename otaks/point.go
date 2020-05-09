package otaks

import "encoding/xml"

type Point struct {
	XMLName   xml.Name `xml:"point"`
	Latitude  float64  `xml:"lat"`
	Longitude float64  `xml:"long"`
	Height    float64  `xml:"hae"`
	Size      float64  `xml:"ce"`
	Error     float64  `xml:"li"`
}
