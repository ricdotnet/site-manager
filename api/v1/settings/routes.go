package settings

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"ricr.dev/site-manager/api/middlewares"
	"ricr.dev/site-manager/config"
)

func Routes(v1 *echo.Group, db *gorm.DB, cfg *config.Config) {
	api := New(db, cfg)

	settings := v1.Group("/settings", middlewares.AuthMiddleware())

	settings.GET("/:key", api.getApiKey)
	settings.PUT("", api.createOrUpdateApiKey)
	settings.DELETE("/:key", api.deleteApiKey)
}
