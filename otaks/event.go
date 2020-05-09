package otaks

import (
	"encoding/xml"
	"time"
)

//
// <?xml version="1.0" encoding="UTF-8"standalone="yes"?>
// <event
//   version="2.0"
//   uid="Linux-ABC.server-ping"
//   type="b-t-f"
//   time="2020-02-14T20:32:31.444Z"
//   start="2020-02-14T20:32:31.444Z"
//   stale="2020-02-15T20:32:31.444Z"
//   how="h-g-i-g-o"
// />
//

type Event struct {
	XMLName xml.Name  `xml:"event"`
	Version string    `xml:"version,attr"`
	UID     string    `xml:"uid,attr"`
	Type    string    `xml:"type,attr"`
	Time    time.Time `xml:"time,attr"`
	Start   time.Time `xml:"start,attr"`
	Stale   time.Time `xml:"stale,attr"`
	How     string    `xml:"how",attr`
}
