package settings

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"ricr.dev/site-manager/api/middlewares"
)

func Routes(v1 *echo.Group, db *gorm.DB) {
	settingsApi := New(db)

	settings := v1.Group("/settings", middlewares.CookieMiddleware(db))

	settings.GET("", settingsApi.getAllApiKeys)
	settings.GET("/:key", settingsApi.getApiKey)
	settings.PUT("", settingsApi.createOrUpdateApiKey)
	settings.DELETE("/:key", settingsApi.deleteApiKey)
}
