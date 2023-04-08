package user

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Routes(g *echo.Group, db *gorm.DB) {
	api := New(db)

	user := g.Group("/user")

	user.POST("/login", api.login)
	user.POST("/register", api.register)
	user.PATCH("/update", api.update)
}
