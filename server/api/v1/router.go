package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ricdotnet/goenvironmental"
	"gorm.io/gorm"
	"net/http"
	"ricr.dev/site-manager/api/v1/domains"
	"ricr.dev/site-manager/api/v1/settings"
	"ricr.dev/site-manager/api/v1/sites"
	"ricr.dev/site-manager/api/v1/user"
	"ricr.dev/site-manager/config"
)

type HealthResponse struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
}

func NewRouter(db *gorm.DB) *echo.Echo {
	allowedOrigins, _ := goenvironmental.Get("CORS_ORIGINS")

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{allowedOrigins, "https://hoppscotch.io"},
		AllowCredentials: true,
	}))

	// ping endpoint to test if the api is running
	e.GET("/ping", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, config.ApiResponse{
			Code:    http.StatusOK,
			Message: "pong",
		})
	})

	e.GET("", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, config.ApiResponse{
			Code:    http.StatusOK,
			Message: "Welcome to the site manager API",
		})
	})

	api := e.Group("/api") // /api group
	v1 := api.Group("/v1") // /v1 group

	sites.Routes(v1, db)
	user.Routes(v1, db)
	settings.Routes(v1, db)
	domains.Routes(v1, db)

	return e
}
