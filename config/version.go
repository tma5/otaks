package config

import (
	"fmt"
	"net"

	log "github.com/sirupsen/logrus"
)

// Version captures the current version
const Version = "0.1.1"

func getVersionString() string {
	return fmt.Sprintf("otaks-%s", Version)
}

// VersionDetail ...
type VersionDetail struct {
	APIVersion string            `json:"version"`
	Type       string            `json:"type"`
	Data       VersionDetailData `json:"data"`
	Node       string            `json:"nodeId"`
}

// VersionDetailData ...
type VersionDetailData struct {
	OtaksVersion string `json:"version"`
	APIVersion   string `json:"api"`
	Hostname     string `json:"hostname"`
}

func getOutboundIP() string {
	// note: this does not actually establish a connection
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String()
}

func getVersionDetail() VersionDetail {
	v := getVersionString()
	i := getOutboundIP()

	return VersionDetail{
		APIVersion: "2",
		Type:       "ServerConfig",
		Data: VersionDetailData{
			OtaksVersion: v,
			APIVersion:   "2",
			Hostname:     i,
		},
		Node: "otaks",
	}
}
