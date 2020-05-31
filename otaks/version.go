package otaks

import (
	"fmt"
	"net"

	log "github.com/sirupsen/logrus"
)

const Version = "undefined"

func getVersionString() string {
	return fmt.Sprintf("otaks-%s", Version)
}

type VersionDetail struct {
	ApiVersion string            `json:"version"`
	Type       string            `json:"type"`
	Data       VersionDetailData `json:"data"`
	Node       string            `json:"nodeId"`
}

type VersionDetailData struct {
	OtaksVersion string `json:"version"`
	ApiVersion   string `json:"api"`
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
		ApiVersion: "2",
		Type:       "ServerConfig",
		Data: VersionDetailData{
			OtaksVersion: v,
			ApiVersion:   "2",
			Hostname:     i,
		},
		Node: "otaks",
	}
}
