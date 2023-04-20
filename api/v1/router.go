package v1

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"ricr.dev/site-manager/api/v1/sites"
	"ricr.dev/site-manager/api/v1/user"
	"ricr.dev/site-manager/config"
)

func New(cfg *config.Config, db *gorm.DB) *echo.Echo {
	e := echo.New()
	//e.Use(cfg.ConfigMiddleware)

	api := e.Group("/api") // /api group

	v1 := api.Group("/v1") // /v1 group
	sites.Routes(v1, db, cfg)
	user.Routes(v1, db, cfg)

	return e
}
