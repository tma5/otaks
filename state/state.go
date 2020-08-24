package state

import (
	"encoding/xml"
	"net/http"
	"sync"

	log "github.com/sirupsen/logrus"
	"github.com/tma5/otaks/config"
	"github.com/tma5/otaks/cot"
)

type State struct {
	Config *config.Config
	mux    sync.Mutex
	Events chan cot.Event
}

func NewState(config *config.Config) *State {
	state := &State{
		Config: config,
		Events: make(chan cot.Event),
	}

	return state
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

	s.Events <- event
	w.WriteHeader(http.StatusCreated)
}
