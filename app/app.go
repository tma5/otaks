package app

import (
	"encoding/xml"
	"fmt"
	"io"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/tma5/otaks/cot"

	"net"
	"sync"
)

type AppServer struct {
	Addr string

	// Shutdown handling
	lock        sync.RWMutex
	started     bool
	shutdown    chan struct{}
	connections map[net.Conn]struct{}
}

func NewAppServer() *AppServer {
	a := AppServer{}
	a.init()
	return &a
}

func (a *AppServer) init() {
	a.shutdown = make(chan struct{})
	a.connections = make(map[net.Conn]struct{})
}

func (a *AppServer) Run() error {
	return a.ListenAndServe()
}

func (a *AppServer) IsRunning() bool {
	return a.started
}

func (a *AppServer) ListenAndServe() error {
	log.Trace("Starting app server on :8087")
	a.started = true
	ln, err := net.Listen("tcp", ":8087")
	if err != nil {
		return err
	}
	return a.ServeTCP(ln)
}

func (a *AppServer) ServeTCP(ln net.Listener) error {
	defer ln.Close()

	wg := sync.WaitGroup{}

	for a.IsRunning() {
		c, err := ln.Accept()
		if err != nil {
			if !a.IsRunning() {
				return nil
			}

			// handle transient issues
			if neterr, ok := err.(net.Error); ok && neterr.Temporary() {
				continue
			}

			return err
		}

		log.Tracef("handling app connection for %s", c.RemoteAddr().String())

		a.connections[c] = struct{}{}
		wg.Add(1)
		go a.ServeTCPConn(&wg, c)
	}

	log.Trace("App server on :8087 died")
	return fmt.Errorf("App server on :8087 died")
}

func (a *AppServer) ServeTCPConn(wg *sync.WaitGroup, c net.Conn) {
	defer func() {
		log.Tracef("closing connection from %s", c.RemoteAddr().String())
		delete(a.connections, c)
		wg.Done()
		c.Close()
	}()

	timeoutDuration := 30 * time.Second

	go func() {
		for {
			time.Sleep(time.Second * 15)

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
		}
	}()

	for {
		c.SetDeadline(time.Now().Add(timeoutDuration))
		buf := make([]byte, 0, 1024) // big buffer
		_, err := c.Read(buf)
		if err != nil {
			if err == io.EOF {
				log.Tracef("%s reached EOF", c.RemoteAddr().String())
				continue
			}
			log.Error("read error:", err)
			break
		}

		if len(buf) > 0 {
			var event cot.Event
			xml.Unmarshal(buf, &event)
			log.WithFields(log.Fields{
				"size": len(buf),
				"raw":  string(buf),
			}).Trace()
			log.WithFields(log.Fields{
				"device":   event.UID,
				"callsign": event.Detail.UID.Droid,
				"lat":      event.Point.Latitude,
				"lon":      event.Point.Longitude,
			}).Debug()

			//wtf
			c.Write(buf)
		}
	}
}
