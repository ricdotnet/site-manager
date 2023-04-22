package sites

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"ricr.dev/site-manager/api/middlewares"
	"ricr.dev/site-manager/config"
)

type Sites struct {
	Sites []string `json:"sites"`
}

// for the user id we are assuming it will always be present on the authorisation token

func Routes(v1 *echo.Group, db *gorm.DB, cfg *config.Config) {
	api := New(db, cfg)

	sites := v1.Group("/sites", middlewares.AuthMiddleware())

	sites.GET("/all/:type", api.all)
	sites.GET("/single/get/:id", api.single)
	sites.POST("/single/create", api.create)
	sites.PATCH("/single/update/:id", api.update)
	sites.PATCH("/single/enable/:id", api.enable)
	sites.DELETE("/single/delete/:id", api.delete)
}
