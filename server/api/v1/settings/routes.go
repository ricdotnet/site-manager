package settings

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"ricr.dev/site-manager/api/middlewares"
)

func Routes(v1 *echo.Group, db *gorm.DB) {
	api := New(db)

	settings := v1.Group("/settings", middlewares.AuthMiddleware())

	settings.GET("/:key", api.getApiKey)
	settings.PUT("", api.createOrUpdateApiKey)
	settings.DELETE("/:key", api.deleteApiKey)
}
