package api

import "net/http"

// OtaksNotImplementedYet represents a resource that is WIP
func OtaksNotImplementedYet(w http.ResponseWriter, r *http.Request) {
	// TODO: actualy do something

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Not Implemented Yet"))
}
