package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/ricdotnet/goenvironmental"
	"regexp"
	"ricr.dev/site-manager/config"
	"ricr.dev/site-manager/models"
	"time"
)

func MakeToken(user *models.User) string {
	claims := &config.JwtCustomClaims{
		UserID:   user.ID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	// Create token with claims
	secret, _ := goenvironmental.Get("APP_KEY")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, _ := token.SignedString([]byte(secret))

	return t
}

func GetTokenClaims(ctx echo.Context) *config.JwtCustomClaims {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(*config.JwtCustomClaims)

	return claims
}

func IsValidFilename(filename string) bool {
	pattern := "^([A-z0-9-])(.{1})+$"
	regex := regexp.MustCompile(pattern)

	return regex.MatchString(filename)
}
