package middlewares

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"ricr.dev/site-manager/config"
	"ricr.dev/site-manager/models"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		authHeader := ctx.Request().Header["Authorization"]

		if authHeader == nil || authHeader[0] == "" {
			return ctx.JSON(http.StatusBadRequest, config.ApiResponse{
				Code:    400,
				Message: "Auth header not present",
			})
		}

		userCtx := models.User{}
		userCtx.Username = authHeader[0]
		ctx.Set("user", userCtx)

		return next(ctx)
	}
}
