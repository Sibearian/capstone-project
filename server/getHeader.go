package server

import (
	"errors"
	"fmt"
	"net/http"
)

func getHeader(key string, r *http.Request) (string, error) {
	res := r.Header.Get(key)
	if res == "" {
		return "", errors.New(fmt.Sprintf("Missing %s in the header", key))
	}

	return res, nil
}
