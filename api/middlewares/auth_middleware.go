package middlewares

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/ricdotnet/goenvironmental"
	cfg "ricr.dev/site-manager/config"
)

func AuthMiddleware() echo.MiddlewareFunc {
	secret, _ := goenvironmental.Get("APP_KEY")

	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(cfg.JwtCustomClaims)
		},
		SigningKey: []byte(secret),
	}

	return echojwt.WithConfig(config)
}
