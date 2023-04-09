package sites

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"ricr.dev/site-manager/config"
)

type Sites struct {
	Sites []string `json:"sites"`
}

// for the user id we are assuming it will always be present on the authorisation token

func Routes(g *echo.Group, db *gorm.DB, cfg *config.Config) {
	api := New(db, cfg)

	sites := g.Group("/sites")

	sites.GET("/all/:type", api.all)
	sites.GET("/single/:id", api.single)
	sites.POST("/single", api.create)
	sites.PATCH("/single/:id", api.update)
	sites.DELETE("/single/:id", api.delete)
}
