package api

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// Version is the api version
const Version = "undefined"

func getVersionString() string {
	return fmt.Sprintf("otaks-%s", Version)
}

type versionDetail struct {
	APIVersion string            `json:"version"`
	Type       string            `json:"type"`
	Data       versionDetailData `json:"data"`
	Node       string            `json:"nodeId"`
}

type versionDetailData struct {
	OtaksVersion string `json:"version"`
	APIVersion   string `json:"api"`
	Hostname     string `json:"hostname"`
}

type dataPackage struct {
	UID                string `json:"UID"`
	Name               string `json:"Name"`
	Hash               string `json:"Hash"`
	PrimaryKey         string `json:"PrimaryKey"`
	SubmissionDateTime string `json:"SubmissionDateTime"`
	SubmissionUser     string `json:"SubmissionUser"`
	CreatorUID         string `json:"CreatorUid"`
	Keywords           string `json:"Keywords"`
	MIMEType           string `json:"MIMEType"`
	Size               string `json:"Size"`
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

func getversionDetail() versionDetail {
	v := getVersionString()
	i := getOutboundIP()

	return versionDetail{
		APIVersion: "2",
		Type:       "ServerConfig",
		Data: versionDetailData{
			OtaksVersion: v,
			APIVersion:   "2",
			Hostname:     i,
		},
		Node: "otaks",
	}
}

func getAPIVersionConfig(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	v := getversionDetail()
	json.NewEncoder(w).Encode(v)
}

func getAPIVersion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/text")
	v := getVersionString()
	w.Write([]byte(v))
}

func getClientEndpoints(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	v := getVersionString()
	w.Write([]byte(v))
}

func getVideoLinks(w http.ResponseWriter, r *http.Request) {
	// TODO: actually get feeds

	w.Header().Set("Content-Type", "application/xml")
	data := "<videoConnections></videoConnections>"
	w.Write([]byte(data))
}

func insertVideoLink(w http.ResponseWriter, r *http.Request) {
	// TODO: actualy insert video link

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Okay"))
}

func uploadDataPackage(w http.ResponseWriter, r *http.Request) {
	// TODO: actually handle upload

	log.Tracef("r: %v", r)

	location := "http://foo:8087/Marti/api/sync/metadata/__hash__/tool"
	w.Write([]byte(location))
}

func modifyDataPackage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	hash := vars["hash"]
	log.Tracef("modifying datapackage %v", hash)

	// TODO: actually modify data package
	w.WriteHeader(http.StatusOK)
}

func getDataPackage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	hash := vars["hash"]
	log.Tracef("getting datapackage %v", hash)

	// TODO: actually get data packages
	w.WriteHeader(http.StatusNotFound)
}

func searchDataPackages(w http.ResponseWriter, r *http.Request) {
	// TODO: actually search data packages
	w.Write([]byte(""))
}

func getDataPackageStatus(w http.ResponseWriter, r *http.Request) {
	// TODO: actually get data package status
	w.WriteHeader(http.StatusNotFound)
}
