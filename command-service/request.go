package main

import (
	"encoding/json"
	"net/http"
)

func incomingBody[T interface{}](r *http.Request) (T, error) {
	body := new(T)

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(body); err != nil {
		return *body, err
	}

	return *body, nil
}
