package docker

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"ricr.dev/site-manager/api/middlewares"
	"ricr.dev/site-manager/services"
)

type DockerAPI struct {
	commandsService *services.CommandsService
}

func new(db *gorm.DB) *DockerAPI {
	return &DockerAPI{
		commandsService: services.NewCommandsService(db),
	}
}

func Routes(v1 *echo.Group, db *gorm.DB) {
	dockerApi := new(db)

	docker := v1.Group("/docker", middlewares.CookieMiddleware(db))

	docker.GET("/containers", dockerApi.getContainers)
	// docker.GET("/images", dockerApi.getImages)
	// docker.GET("/volumes", dockerApi.getVolumes)
	// docker.GET("/networks", dockerApi.getNetworks)
}
