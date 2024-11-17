package user

import (
	"gorm.io/gorm"
	"ricr.dev/site-manager/repository"
	"ricr.dev/site-manager/services"
)

type UserAPI struct {
	db            *gorm.DB
	repo          repository.UserRepository
	sessionRepo   repository.SessionRepository
	loginCodeRepo repository.LoginCodeRepository
	userService   *services.UsersService
}

func New(db *gorm.DB) *UserAPI {
	userRepository := &repository.UserRepo{Db: db}
	sessionRepo := &repository.SessionRepo{Db: db}
	loginCodeRepo := &repository.LoginCodeRepo{Db: db}
	userService := &services.UsersService{}

	return &UserAPI{
		repo:          userRepository,
		sessionRepo:   sessionRepo,
		loginCodeRepo: loginCodeRepo,
		userService:   userService,
	}
}
