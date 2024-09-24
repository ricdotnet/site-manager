package settings

import (
	"gorm.io/gorm"
	"ricr.dev/site-manager/repository"
	"ricr.dev/site-manager/services"
)

type SettingsAPI struct {
	db              *gorm.DB
	settingsRepo    *repository.SettingsRepo
	settingsService *services.SettingsService
}

func New(db *gorm.DB) *SettingsAPI {
	return &SettingsAPI{
		db:              db,
		settingsRepo:    repository.NewSettingsRepo(db),
		settingsService: services.NewSettingsService(),
	}
}
