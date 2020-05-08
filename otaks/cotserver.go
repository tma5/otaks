package otaks

import (
	"bytes"
	"net"
	"strings"
	"fmt" 

	//"github.com/BurntSushi/toml"
	log "github.com/sirupsen/logrus"
)

type CotServer struct {
	Addr    string
	Otaks *Server
}

func (c *CotServer) ListenAndServe() error {
	log.Debugf("Starting COT interface...")
	l, err := net.Listen("tcp", c.Addr)
	if err != nil {
		return err
	}
	log.Infof(fmt.Sprintf("Started COT interface on %s...", c.Addr))

	defer l.Close()
	for {
		err := func() error {
			conn, err := l.Accept()
			if err != nil {
				log.Fatal(err)
			}

			go c.handleCotRequest(conn)

			return nil
		}()
		if err != nil {
			return err
		}
	}
}

func (c *CotServer) Shutdown() {}

func (c *CotServer) handleCotRequest(conn net.Conn) {
	b := make([]byte, 1024)
	_, err := conn.Read(b)
	if err != nil {
		log.Error(err)
		conn.Close()
		return
	}
	t := strings.TrimSpace(string(bytes.Trim(b, "\x00")))
	
	// TODO: actually handle COT messages

	conn.Write([]byte(t))
	conn.Close()
}
