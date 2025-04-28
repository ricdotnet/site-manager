package docker

import (
	"encoding/json"
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/labstack/echo/v4"
	"ricr.dev/site-manager/config"
)

// type Container struct {
// 	ID           string `json:"id"`
// 	Names        string `json:"names"`
// 	Image        string `json:"image"`
// 	Command      string `json:"command"`
// 	CreatedAt    string `json:"created_at"`
// 	State        string `json:"state"`
// 	Status       string `json:"status"`
// 	Ports        string `json:"ports"`
// 	Size         string `json:"size"`
// 	Labels       string `json:"labels"`
// 	Mounts       string `json:"mounts"`
// 	Networks     string `json:"networks"`
// 	LocalVolumes string `json:"local_volumes"`
// }

func (d *DockerAPI) getContainers(ctx echo.Context) error {
	userCtx := ctx.Get("user")

	log.Infof("Getting docker containers for user %d", userCtx.(*config.Session).UserID)

	containers, err := d.commandsService.GetContainers(userCtx)

	if err != nil {
		log.Errorf("Error retrieving containers: %v", err)

		return ctx.JSON(http.StatusInternalServerError, &config.Response[any]{
			Code:    http.StatusInternalServerError,
			Message: "Error retrieving containers",
		})
	}

	log.Infof("Retrieved %d containers", len(containers))

	unmarshalledContainers := make([]config.Container, 0)

	for _, container := range containers {
		unmarshalledContainer := &config.Container{}

		log.Debugf("Container: %s", container)

		err := json.Unmarshal([]byte(container), unmarshalledContainer)

		if err != nil {
			log.Errorf("Error unmarshalling container: %v", err)

			return ctx.JSON(http.StatusInternalServerError, &config.Response[any]{
				Code:    http.StatusInternalServerError,
				Message: "Error unmarshalling container data",
			})
		}

		unmarshalledContainers = append(unmarshalledContainers, *unmarshalledContainer)
	}

	return ctx.JSON(http.StatusOK, &config.Response[[]config.Container]{
		Code:    http.StatusOK,
		Message: "Containers retrieved successfully",
		Data:    unmarshalledContainers,
	})
}
