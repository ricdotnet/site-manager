package config

import (
	"github.com/labstack/echo/v4"
)

type Config struct {
	//Logger *logging.Logger
}

func (cfg *Config) ConfigMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return next(ctx)
	}
}
