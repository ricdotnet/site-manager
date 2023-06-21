package sites

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"ricr.dev/site-manager/api/middlewares"
	"ricr.dev/site-manager/config"
)

// for the user id we are assuming it will always be present on the authorisation token

func Routes(v1 *echo.Group, db *gorm.DB, cfg *config.Config) {
	api := New(db, cfg)

	sites := v1.Group("/site", middlewares.AuthMiddleware())

	sites.GET("/all", api.all)
	sites.GET("/:id", api.single)
	sites.POST("/", api.create)
	sites.PATCH("/:id", api.update)
	sites.PATCH("/:id/status", api.status)
	sites.DELETE("/", api.delete)
}
