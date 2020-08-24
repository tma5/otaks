package api

import "net/http"

func OtaksNotImplementedYet(w http.ResponseWriter, r *http.Request) {
	// TODO: actualy do something

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Okay"))
}
