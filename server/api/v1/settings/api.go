package settings

import (
	"gorm.io/gorm"
	"ricr.dev/site-manager/repository"
	"ricr.dev/site-manager/services"
)

type SettingsAPI struct {
	db              *gorm.DB
	repository      *repository.Repository
	settingsService *services.SettingsService
}

func New(db *gorm.DB) *SettingsAPI {
	return &SettingsAPI{
		repository:      repository.NewRepository(db),
		settingsService: services.NewSettingsService(),
	}
}
