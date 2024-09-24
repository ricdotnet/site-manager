package user

import (
	"gorm.io/gorm"
	"ricr.dev/site-manager/repository"
)

type UserAPI struct {
	db       *gorm.DB
	userRepo *repository.UserRepo
}

func New(db *gorm.DB) *UserAPI {
	return &UserAPI{
		db:       db,
		userRepo: repository.NewUserRepo(db),
	}
}
