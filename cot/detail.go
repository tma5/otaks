package cot

// Detail provides the CoT detail element
type Detail struct {
	Text string `xml:",chardata"`
	Takv struct {
		Text     string `xml:",chardata,omit"`
		Os       string `xml:"os,attr,omitempty"`
		Version  string `xml:"version,attr,omitempty"`
		Device   string `xml:"device,attr,omitempty"`
		Platform string `xml:"platform,attr,omitempty"`
	} `xml:"takv,omitempty"`
	Contact struct {
		Text     string `xml:",chardata,omit"`
		Endpoint string `xml:"endpoint,attr,omitempty"`
		Callsign string `xml:"callsign,attr,omitempty"`
	} `xml:"contact,omitempty"`
	UID struct {
		Text  string `xml:",chardata,omit"`
		Droid string `xml:"Droid,attr,omitempty"`
	} `xml:"uid,omitempty"`
	Precisionlocation struct {
		Text        string `xml:",chardata,omit"`
		Altsrc      string `xml:"altsrc,attr,omitempty"`
		Geopointsrc string `xml:"geopointsrc,attr,omitempty"`
	} `xml:"precisionlocation,omitempty"`
	Group struct {
		Text string `xml:",chardata,omit"`
		Role string `xml:"role,attr,omitempty"`
		Name string `xml:"name,attr,omitempty"`
	} `xml:"__group,omitempty"`
	Status struct {
		Text    string `xml:",chardata,omit"`
		Battery string `xml:"battery,attr,omitempty"`
	} `xml:"status,omitempty"`
	Track struct {
		Text   string `xml:",chardata,omit"`
		Course string `xml:"course,attr,omitempty"`
		Speed  string `xml:"speed,attr,omitempty"`
	} `xml:"track,omitempty"`
}
