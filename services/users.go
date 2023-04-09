package services

import (
	"github.com/op/go-logging"
	"ricr.dev/site-manager/config"
)

type UsersService struct {
	logger *logging.Logger
}

func NewUsersService(cfg *config.Config) *UsersService {
	return &UsersService{
		logger: cfg.Logger,
	}
}
