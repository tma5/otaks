package app

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"io"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/tma5/otaks/cot"
	"github.com/tma5/otaks/state"

	"net"
	"sync"
)

type connectionData struct {
}

// Server provides the app server state
type Server struct {
	Addr string

	state *state.State

	// Shutdown handling
	lock     sync.RWMutex
	started  bool
	shutdown chan struct{}

	connections map[net.Conn]connectionData
	//events      chan cot.Event
}

// NewServer provides a new instance of the app server
func NewServer(state *state.State) *Server {
	srv := Server{
		state: state,
	}
	srv.init()
	return &srv
}

func (srv *Server) init() {
	srv.shutdown = make(chan struct{})
	srv.connections = make(map[net.Conn]connectionData)
	//srv.events = make(chan cot.Event)
}

// Run starts the app server
func (srv *Server) Run() error {
	return srv.listenAndServe()
}

// IsRunning provides the state of the server
func (srv *Server) IsRunning() bool {
	return srv.started
}

func (srv *Server) listenAndServe() error {
	log.Trace("Starting app server on :8087")
	srv.started = true
	ln, err := net.Listen("tcp", ":8087")
	if err != nil {
		return err
	}
	return srv.listen(ln)
}

func (srv *Server) listen(ln net.Listener) error {
	defer ln.Close()

	wg := sync.WaitGroup{}

	go srv.retransmitEvents()

	for srv.IsRunning() {
		c, err := ln.Accept()
		if err != nil {
			if !srv.IsRunning() {
				// handle server closing state
				return nil
			}

			// handle transient issues
			if neterr, ok := err.(net.Error); ok && neterr.Temporary() {
				log.Tracef("network hiccup %v", err)
				continue
			}

			return err
		}

		log.Tracef("handling app connection for %s", c.RemoteAddr().String())

		srv.connections[c] = connectionData{}
		wg.Add(1)
		go srv.handleConnection(&wg, c)
	}

	log.Trace("App server on :8087 died")
	return fmt.Errorf("App server on :8087 died")
}

func (srv *Server) retransmitEvents() {
	log.Trace("beginning to transmit events")
	for srv.IsRunning() {
		event := <-srv.state.Events
		for r := range srv.connections {
			if r.RemoteAddr().String() == event.Origin {
				continue
			}
			log.Tracef("retransmitting event from %s to %s", event.Origin, r.RemoteAddr().String())
			e, err := xml.Marshal(event)
			if err != nil {
				log.Error("problem repackaging event: ", err)
				continue
			}
			log.Tracef("transmitting event: %+v", string(e))
			_, err = r.Write(e)
			if err != nil {
				log.Error("problem retransmitting event: ", err)
				break
			}
		}

	}
}

func (srv *Server) handleConnection(wg *sync.WaitGroup, c net.Conn) {
	defer func() {
		log.Tracef("closing connection from %s", c.RemoteAddr().String())
		delete(srv.connections, c)
		wg.Done()
		c.Close()
	}()

	// handle keepalives
	go func() {
		time.Sleep(time.Second * 2)
		for {
			log.Tracef("sending ping event to %s", c.RemoteAddr().String())
			pingEvent := cot.NewPingEvent()
			e, err := xml.Marshal(pingEvent)
			if err != nil {
				log.Error("problem creating ping event: ", err)
				continue
			}
			_, err = c.Write(e)
			if err != nil {
				log.Error("problem sending ping event: ", err)
				break
			}

			time.Sleep(time.Second * 15)
		}
	}()

	r := bufio.NewReader(c)
	for {
		log.Tracef("reading from", c.RemoteAddr())
		b, err := r.ReadBytes(byte('\n'))
		if err != nil {
			if err != io.EOF {
				log.Tracef("issue reading from conn. %v", err)
			}
		}

		var event cot.Event
		if xml.Unmarshal(b, &event) != nil {
			if err == io.EOF {
				break
			}
			log.Error("failed to interpret event", err)
			continue
		}
		event.Origin = c.RemoteAddr().String()

		// // handle event based on type
		// switch event.Type {
		// case cot.PingEvent:
		// 	c.Write(buf.Bytes())
		// }

		log.WithFields(log.Fields{
			"size": len(b),
			"raw":  string(b),
		}).Trace()
		log.WithFields(log.Fields{
			"origin":   event.Origin,
			"device":   event.UID,
			"callsign": event.Detail.UID.Droid,
			"type":     event.Type,
			"how":      event.How,
			"lat":      event.Point.Latitude,
			"lon":      event.Point.Longitude,
		}).Debug()

		// send the event to the queue
		srv.state.Events <- event

		//wtf
		//c.Write(b)
	}
}
