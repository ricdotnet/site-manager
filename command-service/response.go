package main

import (
	"encoding/json"
	"github.com/charmbracelet/log"
	"net/http"
)

type response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Failed  bool   `json:"failed"`
}

func send(w http.ResponseWriter, res *response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.Code)

	outgoingMessage, _ := json.Marshal(res)

	_, err := w.Write(outgoingMessage)

	if err != nil {
		log.Fatalf("error writing response: %v", err)
	}
}
