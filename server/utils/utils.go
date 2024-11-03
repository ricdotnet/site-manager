package utils

import (
	"crypto/rand"
	"encoding/base32"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/ricdotnet/goenvironmental"
	"net/http"
	"regexp"
	"ricr.dev/site-manager/config"
	"time"
)

func MakeToken() string {
	bytes := make([]byte, 15)
	_, _ = rand.Read(bytes)
	return base32.StdEncoding.EncodeToString(bytes)
}

func MakeEmptyCookie() *http.Cookie {
	now := time.Now()

	cookieName, _ := goenvironmental.Get("COOKIE_NAME")
	cookieDomain, _ := goenvironmental.Get("COOKIE_DOMAIN")

	return &http.Cookie{
		Name:    cookieName,
		Domain:  cookieDomain,
		Path:    "/",
		Expires: now.AddDate(0, 0, -1),
	}
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

func DomainsApiUrl(url, userId, apiKey string) string {
	environment, _ := goenvironmental.Get("ENV")

	if environment == "production" {
		return fmt.Sprintf("https://httpapi.com/api/%s?auth-userid=%s&api-key=%s", url, userId, apiKey)
	}

	return fmt.Sprintf("https://test.httpapi.com/api/%s?auth-userid=%s&api-key=%s", url, userId, apiKey)
}
