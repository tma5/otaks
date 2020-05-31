package cot

import (
	"encoding/xml"
	"fmt"
)

// EventType describes the CoT Type
type EventType string

const (
	// EmergencySendEvent is a shortcut to the CoT type 'b-a-o-tbl'
	EmergencySendEvent EventType = "b-a-o-tbl"

	// EmergencyCancelEvent is a shortcut to the CoT type 'b-a-o-can'
	EmergencyCancelEvent EventType = "b-a-o-can"

	// MartiDataEvent is a shortcut to the CoT type 'b-f-t-a'
	MartiDataEvent EventType = "b-f-t-a"

	// GeoChatEvent is a shortcut to the CoT type 'b-t-f'
	GeoChatEvent EventType = "b-t-f"

	// PingEvent is a shortcut to the CoT type 't-x-c-t'
	PingEvent EventType = "t-x-c-t"

	// GroundCombatEvent is a shortcut to the CoT type 'a-f-G-U-C'
	GroundCombatEvent EventType = "a-f-G-U-C"

	// TimeoutEvent is a shortcut to the NonCoT type 'timeout'
	// this might not be correct
	TimeoutEvent EventType = "timeout"
)

// HowType describes the CoT How attribute
type HowType string

const (
	// FromGPS is a shortcut for CoT How 'm-g'
	FromGPS HowType = "m-g"

	// NonCoT is a shortcut for CoT how 'h-g-i-g-o'
	NonCoT HowType = "h-g-i-g-o" // is this correct?
)

// Event describes a cursor-on-target event
type Event struct {
	XMLName xml.Name  `xml:"event"`
	Text    string    `xml:",chardata"`
	Version string    `xml:"version,attr"`
	UID     string    `xml:"uid,attr"`
	Type    EventType `xml:"type,attr"`
	Time    string    `xml:"time,attr"`
	Start   string    `xml:"start,attr"`
	Stale   string    `xml:"stale,attr"`
	How     HowType   `xml:"how,attr"`
	Point   Point
	Opex    string `xml:"opex,omitempty"`
	QOS     string `xml:"qos,omitempty"`
	Access  string `xml:"access,omitempty"`
	Detail  Detail `xml:"detail,omitempty"`
}

// NewEvent gets a blank event
func NewEvent() Event {
	p := NewUndefinedPoint()
	e := Event{
		Point: p,
	}
	return e
}

// NewTimeoutEvent provides a timeout event
func NewTimeoutEvent() Event {
	e := NewEvent()
	e.How = NonCoT
	e.Type = TimeoutEvent

	d := Detail{}
	e.Detail = d

	return e
}

// NewGeoChatEvent provides a geochat event
func NewGeoChatEvent() Event {
	e := NewEvent()
	e.UID = fmt.Sprintf("GeoChat.%s", e.UID)
	e.How = NonCoT
	e.Type = GeoChatEvent

	return e
}

// NewPingEvent provides a ping event
func NewPingEvent() Event {
	e := NewEvent()
	e.UID = fmt.Sprintf("%s-ping", e.UID)
	e.How = FromGPS
	e.Type = PingEvent
	d := Detail{}
	e.Detail = d

	return e
}
