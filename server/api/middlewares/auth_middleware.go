package middlewares

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/ricdotnet/goenvironmental"
	"net/http"
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

func CookieMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		cookieName, _ := goenvironmental.Get("COOKIE_NAME")

		cookie, err := ctx.Cookie(cookieName)

		if cookie == nil || err != nil {
			return ctx.JSON(http.StatusUnauthorized, map[string]interface{}{})
		}

		ctx.Set("user", &cfg.Session{
			UserID: 1,
			Role:   "Admin",
		})

		return next(ctx)
	}
}
