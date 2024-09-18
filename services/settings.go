package services

import (
	"github.com/op/go-logging"
	"ricr.dev/site-manager/config"
)

type SettingsService struct {
	logger *logging.Logger
}

func NewSettingsService(cfg *config.Config) *SettingsService {
	return &SettingsService{
		logger: cfg.Logger,
	}
}
