package state

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"sync"

	log "github.com/sirupsen/logrus"
	"github.com/tma5/otaks/config"
	"github.com/tma5/otaks/cot"
)

type State struct {
	Config *config.Config
	mux    sync.Mutex
	Events []cot.Event
}

func NewState(config *config.Config) *State {
	state := &State{
		Config: config,
		Events: make([]cot.Event, 0),
	}

	return state
}

func (s *State) QueueEvent(event cot.Event) {
	s.mux.Lock()
	s.Events = append(s.Events, event)
	s.mux.Unlock()
}

func (s *State) NextEvent() (*cot.Event, error) {
	if len(s.Events) > 0 {

		e := s.Events[0]
		s.Events = s.Events[1:]

		return &e, nil
	}

	return nil, fmt.Errorf("No event to dequeue")
}

func (s *State) QueueEventFromHttpRequest(w http.ResponseWriter, r *http.Request) {
	var event cot.Event
	if err := xml.NewDecoder(r.Body).Decode(&event); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	event.Text = ""
	event.Detail.Text = ""

	event.Origin = "api"

	log.WithFields(log.Fields{
		"origin":   event.Origin,
		"device":   event.UID,
		"callsign": event.Detail.UID.Droid,
		"type":     event.Type,
		"how":      event.How,
		"lat":      event.Point.Latitude,
		"lon":      event.Point.Longitude,
	}).Debug()

	s.QueueEvent(event)
	w.WriteHeader(http.StatusCreated)
}
