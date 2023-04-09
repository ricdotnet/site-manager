package user

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"ricr.dev/site-manager/config"
)

func Routes(g *echo.Group, db *gorm.DB, cfg *config.Config) {
	api := New(db, cfg)

	user := g.Group("/user")

	user.POST("/login", api.login)
	user.POST("/register", api.register)
	user.PATCH("/update", api.update)
}
