package v1

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"ricr.dev/site-manager/api/v1/sites"
	"ricr.dev/site-manager/config"
)

func New(cfg *config.Config, db *gorm.DB) *echo.Echo {
	e := echo.New()
	//e.Use(cfg.ConfigMiddleware)

	v1 := e.Group("/api/v1")
	sites.Routes(v1, db)

	return e
}
