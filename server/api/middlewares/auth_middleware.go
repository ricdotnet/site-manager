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

			if cookie == nil || err != nil {
				return ctx.JSON(http.StatusUnauthorized, map[string]interface{}{})
			}

			sessionRepo := &repository.SessionRepo{Db: db}
			userRepo := &repository.UserRepo{Db: db}

			session := &models.Session{}
			_ = sessionRepo.GetOne(cookie.Value, session)

			now := time.Now()

			if session.ExpiresAt.Before(now) {
				log.Warnf("Session %s for user %s is expired", session.Token, session.User.Username)

				_ = sessionRepo.DeleteOne(session.Token)
				ctx.SetCookie(utils.MakeEmptyCookie())

				return ctx.JSON(http.StatusUnauthorized, map[string]interface{}{})
			}

			user := &models.User{}
			_ = userRepo.GetOneByID(session.UserId, user)

			ctx.Set("token", session.Token)
			ctx.Set("user", &cfg.Session{
				UserID:   session.UserId,
				Username: user.Username,
				Email:    user.Email,
				Role:     "Admin",
			})

			return next(ctx)
		}
	}
}
