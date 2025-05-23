package user

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"ricr.dev/site-manager/api/middlewares"
)

func Routes(v1 *echo.Group, db *gorm.DB) {
	userApi := New(db)

	user := v1.Group("/user")

	user.PATCH("/update", userApi.updateUser, middlewares.CookieMiddleware(db))
	user.GET("/auth", userApi.authUser, middlewares.CookieMiddleware(db))
	user.GET("/sessions", userApi.getSessions, middlewares.CookieMiddleware(db))
	user.DELETE("/sessions/:id", userApi.deleteSession, middlewares.CookieMiddleware(db))
	user.GET("/logout", userApi.logout, middlewares.CookieMiddleware(db))

	user.POST("/sign-in", userApi.signIn)
	user.POST("/verify-code", userApi.verifyCode)
}
