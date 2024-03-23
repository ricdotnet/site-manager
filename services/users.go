package services

import (
	"ricr.dev/site-manager/config"
)

type UsersService struct {
}

func NewUsersService(cfg *config.Config) *UsersService {
	return &UsersService{}
}
