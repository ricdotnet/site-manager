package main

import (
	"net/http"

	"github.com/charmbracelet/log"
)

func reloadNginx(w http.ResponseWriter, r *http.Request) {
	_, err := execute("nginx", []string{"-s", "reload"})

	if err != nil {
		log.Error(err.Error())
		send(w, &response{
			Code:    http.StatusOK,
			Message: err.Error(),
			Failed:  true,
		})
		return
	}

	send(w, &response{
		Code:    http.StatusOK,
		Message: "Reloaded nginx configuration",
		Failed:  false,
	})
}

func certificates(w http.ResponseWriter, r *http.Request) {
	out, err := execute("certbot", []string{"certificates"})

	if err != nil {
		log.Error(err.Error())
		send(w, &response{
			Code:    http.StatusOK,
			Message: err.Error(),
			Failed:  true,
		})
		return
	}

	send(w, &response{
		Code:    http.StatusOK,
		Message: out,
		Failed:  false,
	})
}

func customCommand(w http.ResponseWriter, r *http.Request) {
	body, err := incomingBody[reloadNginxBody](r)

	if err != nil {
		log.Error(err)
		send(w, &response{
			Code:    http.StatusBadRequest,
			Message: "failed to parse incoming payload",
			Failed:  false,
		})
		return
	}

	out, err := execute(body.Command, body.Args)

	if err != nil {
		log.Error(err)
		send(w, &response{
			Code:    http.StatusOK,
			Message: err.Error(),
			Failed:  true,
		})
		return
	}

	send(w, &response{
		Code:    http.StatusOK,
		Message: out,
		Failed:  false,
	})
}

func dockerContainers(w http.ResponseWriter, r *http.Request) {
	out, err := execute("docker", []string{"ps", "-a", "--format={{json .}}"})

	if err != nil {
		log.Error(err.Error())
		send(w, &response{
			Code:    http.StatusOK,
			Message: err.Error(),
			Failed:  true,
		})
		return
	}

	send(w, &response{
		Code:    http.StatusOK,
		Message: out,
		Failed:  false,
	})
}
