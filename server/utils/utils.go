package utils

import (
	"crypto/rand"
	"encoding/base32"
	"fmt"
	"github.com/ricdotnet/goenvironmental"
	"net/http"
	"regexp"
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
