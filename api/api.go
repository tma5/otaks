package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/tma5/otaks/state"
)

// Server provides the state of the api server
type Server struct {
	state *state.State
}

// NewServer provides a new instance of an api server
func NewServer(state *state.State) *Server {
	return &Server{
		state: state,
	}
}

// Run begins the server
func (srv *Server) Run() error {
	router := mux.NewRouter()

	router.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(struct {
			Status string `json:"status"`
		}{
			Status: "ok",
		})
	}).Methods("GET")

	// TODO: marti routes
	router.HandleFunc("/Marti/api/version/config", getAPIVersionConfig).Methods("GET")
	router.HandleFunc("/Marti/api/clientEndPoints", getClientEndpoints).Methods("GET")
	router.HandleFunc("/Marti/api/sync/metadata/{hash}/tool", modifyDataPackage).Methods("PUT")
	router.HandleFunc("/Marti/api/sync/metadata/{hash}/tool", getDataPackage).Methods("GET")

	router.HandleFunc("/Marti/sync/search", searchDataPackages).Methods("GET")
	router.HandleFunc("/Marti/sync/missionupload", uploadDataPackage).Methods("POST")
	router.HandleFunc("/Marti/sync/content", getDataPackage).Methods("GET")
	router.HandleFunc("/Marti/sync/missionquery", getDataPackageStatus).Methods("GET")

	router.HandleFunc("/Marti/vcm", getVideoLinks).Methods("GET")
	router.HandleFunc("/Marti/vcm", insertVideoLink).Methods("POST")

	router.HandleFunc("/Marti/KmlMasterSA", getKMLManifest).Methods("GET")
	router.HandleFunc("/Marti/LatestKML", getKML).Methods("GET")

	router.HandleFunc("/Marti/webtak", redirectWeb).Methods("GET")

	// otaks rest api
	router.HandleFunc("/otaks/events", srv.state.QueueEventFromHttpRequest).Methods("POST")
	router.HandleFunc("/otaks/devices", OtaksNotImplementedYet).Methods("GET")

	log.Trace("Initializing api server on :8080")
	return http.ListenAndServe(":8080", router)
}
