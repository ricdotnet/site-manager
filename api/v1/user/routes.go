package user

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"ricr.dev/site-manager/api/middlewares"
)

func Routes(v1 *echo.Group, db *gorm.DB) {
	api := New(db)

	user := v1.Group("/user")

	user.POST("/login", api.loginUser)
	user.POST("/register", api.registerUser)
	user.PATCH("/update", api.updateUser, middlewares.AuthMiddleware())
	user.GET("/auth", api.authUser, middlewares.AuthMiddleware())
}
