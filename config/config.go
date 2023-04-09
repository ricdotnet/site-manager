package config

import (
	"github.com/labstack/echo/v4"
	"github.com/op/go-logging"
	"os"
)

type Config struct {
	Logger *logging.Logger
}

func NewConfig() *Config {
	return &Config{
		Logger: NewLogger(),
	}
}

func (cfg *Config) ConfigMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return next(ctx)
	}
}

func NewLogger() *logging.Logger {
	logger := logging.MustGetLogger("site-manager")

	formatter := logging.MustStringFormatter(
		"%{color}%{time:15:04:05.000} %{longfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}")
	loggerBackend := logging.NewLogBackend(os.Stderr, "", 0)
	backendFormatter := logging.NewBackendFormatter(loggerBackend, formatter)

	backendLeveled := logging.AddModuleLevel(loggerBackend)
	backendLeveled.SetLevel(logging.ERROR, "")

	logging.SetBackend(backendLeveled, backendFormatter)

	return logger
}
