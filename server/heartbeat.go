package server

import "net/http"

func heartbeat(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("All good"))
}
