package cot

// Point describe a location
type Point struct {
	Latitude  string `xml:"lat,attr"`
	Longitude string `xml:"long,attr"`
	Height    string `xml:"hae,attr"`
	Size      string `xml:"ce,attr"`
	Error     string `xml:"le,attr"`
}

// NewUndefinedPoint gets a point defining an undefined location
func NewUndefinedPoint() Point {
	return Point{
		Latitude:  "00.00000000",
		Longitude: "00.00000000",
		Height:    "00.00000000",
		Size:      "9999999.0",
		Error:     "9999999.0",
	}
}
