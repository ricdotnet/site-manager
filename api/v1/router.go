package v1

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"ricr.dev/site-manager/api/v1/sites"
	"ricr.dev/site-manager/api/v1/user"
	"ricr.dev/site-manager/config"
)

func New(cfg *config.Config, db *gorm.DB) *echo.Echo {
	e := echo.New()
	//e.Use(cfg.ConfigMiddleware)

	// ping endpoint to test if the api is running
	e.GET("/ping", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, config.ApiResponse{
			Code:    200,
			Message: "pong",
		})
	})

	api := e.Group("/api") // /api group

	v1 := api.Group("/v1") // /v1 group
	sites.Routes(v1, db, cfg)
	user.Routes(v1, db, cfg)

	return e
}
