package middlewares

import (
	"github.com/charmbracelet/log"
	"github.com/labstack/echo/v4"
	"github.com/ricdotnet/goenvironmental"
	"gorm.io/gorm"
	"net/http"
	cfg "ricr.dev/site-manager/config"
	"ricr.dev/site-manager/models"
	"ricr.dev/site-manager/repository"
	"ricr.dev/site-manager/utils"
	"time"
)

func CookieMiddleware(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			cookieName, _ := goenvironmental.Get("COOKIE_NAME")
			cookie, err := ctx.Cookie(cookieName)
			userAgent := ctx.Request().Header.Get("User-Agent")

			if cookie == nil || err != nil {
				return ctx.JSON(http.StatusUnauthorized, map[string]interface{}{})
			}

			sessionRepo := &repository.SessionRepo{Db: db}
			session := &models.Session{}

			_ = sessionRepo.GetOne(cookie.Value, session)

			userRepo := &repository.UserRepo{Db: db}
			user := &models.User{}
			_ = userRepo.GetOneByID(session.UserID, user)

			if session.Token == "" {
				log.Warnf("Session %s was not found", cookie.Value)

				ctx.SetCookie(utils.MakeEmptyCookie())

				return ctx.JSON(http.StatusUnauthorized, map[string]interface{}{})
			}

			if session.UserAgent != userAgent {
				log.Warnf("Session %s for user %s is invalid because user agent does not match our records", session.Token, user.Username)

				_ = sessionRepo.DeleteOne(session.Token)
				ctx.SetCookie(utils.MakeEmptyCookie())

				return ctx.JSON(http.StatusUnauthorized, map[string]interface{}{})
			}

			now := time.Now()
			session.LastActive = now

			if session.ExpiresAt.Before(now) {
				log.Warnf("Session %s for user %s is expired", session.Token, user.Username)

				_ = sessionRepo.DeleteOne(session.Token)
				ctx.SetCookie(utils.MakeEmptyCookie())

				return ctx.JSON(http.StatusUnauthorized, map[string]interface{}{})
			}

			if session.ExpiresAt.Sub(now) < 24*time.Hour {
				log.Warnf("Session %s for user %s is about to expire", session.Token, user.Username)

				session.ExpiresAt = now.AddDate(0, 0, 30)

				log.Warnf("Session %s for user %s is extended to %s", session.Token, user.Username, session.ExpiresAt)
			}

			_ = sessionRepo.UpdateOne(session)

			ctx.Set("token", session.Token)
			ctx.Set("user", &cfg.Session{
				UserID:   session.UserID,
				Username: user.Username,
				Email:    user.Email,
				Role:     "Admin",
			})

			return next(ctx)
		}
	}
}
