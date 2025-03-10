package user

import (
	"gorm.io/gorm"
	"ricr.dev/site-manager/repository"
	"ricr.dev/site-manager/services"
)

type UserAPI struct {
	db          *gorm.DB
	repository  *repository.Repository
	userService *services.UsersService
}

func New(db *gorm.DB) *UserAPI {
	return &UserAPI{
		repository:  repository.NewRepository(db),
		userService: &services.UsersService{},
	}
}
