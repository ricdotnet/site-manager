package user

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"ricr.dev/site-manager/api/middlewares"
	"ricr.dev/site-manager/config"
)

func Routes(v1 *echo.Group, db *gorm.DB) {
	userApi := New(db)

	user := v1.Group("/user")

	user.POST("/login", userApi.loginUser)
	user.POST("/register" /*userApi.registerUser*/, func(ctx echo.Context) error {
		return ctx.JSON(http.StatusNotImplemented, &config.ApiResponse{
			Code:        http.StatusNotImplemented,
			MessageCode: "not_implemented",
		})
	})
	user.PATCH("/update", userApi.updateUser, middlewares.CookieMiddleware(db))
	user.GET("/auth", userApi.authUser, middlewares.CookieMiddleware(db))
	user.GET("/sessions", userApi.getSessions, middlewares.CookieMiddleware(db))
	user.DELETE("/sessions/:id", userApi.deleteSession, middlewares.CookieMiddleware(db))
	user.GET("/logout", userApi.logout, middlewares.CookieMiddleware(db))
}
