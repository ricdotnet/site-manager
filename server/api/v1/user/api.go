package user

import (
	"gorm.io/gorm"
	"ricr.dev/site-manager/repository"
)

type UserAPI struct {
	db   *gorm.DB
	repo repository.UserRepository
}

func New(db *gorm.DB) *UserAPI {
	userRepositoy := &repository.UserRepo{Db: db}

	return &UserAPI{
		db:   db,
		repo: userRepositoy,
	}
}
