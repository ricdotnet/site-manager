package domains

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"ricr.dev/site-manager/api/middlewares"
	"ricr.dev/site-manager/config"
)

func Routes(v1 *echo.Group, db *gorm.DB, cfg *config.Config) {
	//api := New(db, cfg)

	domains := v1.Group("/domains", middlewares.AuthMiddleware())

	domains.GET("/all", func(c echo.Context) error {
		return c.JSON(http.StatusOK, config.ApiResponse{
			Code:    200,
			Message: "all domains",
		})
	})
}
