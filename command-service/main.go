package main

import (
	"github.com/charmbracelet/log"
	"github.com/ricdotnet/goenvironmental"
	"net/http"
)

type reloadNginxBody struct {
	Command string   `json:"command"`
	Args    []string `json:"args"`
}

func main() {
	goenvironmental.ParseEnv()

	log.SetReportCaller(true)

	var port string
	port, _ = goenvironmental.Get("PORT")

	if port == "" {
		port = "8080"
	}

	http.HandleFunc("POST /reload-nginx", auth(reloadNginx))
	http.HandleFunc("POST /certificates", auth(certificates))

	http.HandleFunc("POST /custom-command", auth(customCommand))

	log.Infof("Server listening on port %s", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
