package server

import "net/http"

func sendError(err error, errorCode int, w http.ResponseWriter) {
	w.WriteHeader(errorCode)
	w.Write([]byte(err.Error()))

	return
}
