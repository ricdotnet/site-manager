package user

import (
	"gorm.io/gorm"
	"ricr.dev/site-manager/repository"
)

type UserAPI struct {
	db            *gorm.DB
	repo          repository.UserRepository
	sessionRepo   repository.SessionRepository
	loginCodeRepo repository.LoginCodeRepository
}

func New(db *gorm.DB) *UserAPI {
	userRepository := &repository.UserRepo{Db: db}
	sessionRepo := &repository.SessionRepo{Db: db}
	loginCodeRepo := &repository.LoginCodeRepo{Db: db}

	return &UserAPI{
		repo:          userRepository,
		sessionRepo:   sessionRepo,
		loginCodeRepo: loginCodeRepo,
	}
}
