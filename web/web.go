package web

import (
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/tma5/otaks/state"
)

type Server struct {
	state  *state.State
	assets string
}

func NewServer(state *state.State) *Server {
	return &Server{
		state: state,
	}
}

func (srv *Server) Run() error {
	router := mux.NewRouter()

	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets", http.FileServer(http.Dir(srv.assets))))

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	}).Methods("GET")

	router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {

	}).Methods("GET", "POST")

	log.Trace("Initializing web server on :8888")
	return http.ListenAndServe(":8888", router)
}
