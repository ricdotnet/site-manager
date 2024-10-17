package settings

import (
	"gorm.io/gorm"
	"ricr.dev/site-manager/repository"
	"ricr.dev/site-manager/services"
)

type SettingsAPI struct {
	db              *gorm.DB
	repo            repository.SettingsRepository
	settingsService *services.SettingsService
}

func New(db *gorm.DB) *SettingsAPI {
	settingsRepo := &repository.SettingsRepo{Db: db}

	return &SettingsAPI{
		db:              db,
		repo:            settingsRepo,
		settingsService: services.NewSettingsService(),
	}
}
