package main

import (
	"github.com/charmbracelet/log"
	"github.com/ricdotnet/goenvironmental"
	"net/http"
)

func auth(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Infof("Incoming request %s", r.URL.Path)

		authKey, _ := goenvironmental.Get("AUTHORIZATION_KEY")
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" || authHeader != authKey {
			log.Error("Invalid authorization header")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		handlerFunc(w, r)
	}
}
