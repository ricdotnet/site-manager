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

	sites := v1.Group("/sites", middlewares.AuthMiddleware())

	sites.GET("/all/:type", api.all)
	sites.GET("/single/:id", api.single)
	sites.POST("/single", api.create)
	sites.PATCH("/single/:id", api.update)
	sites.PATCH("/single/:id/status", api.status)
	sites.DELETE("/single/:id", api.delete)
}
