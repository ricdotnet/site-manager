package user

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"ricr.dev/site-manager/api/middlewares"
)

func Routes(v1 *echo.Group, db *gorm.DB) {
	api := New(db)

	user := v1.Group("/user")

	user.POST("/login", api.login)
	user.POST("/register", api.register)
	user.PATCH("/update", api.update)
	user.GET("/auth", api.auth, middlewares.AuthMiddleware())
}
