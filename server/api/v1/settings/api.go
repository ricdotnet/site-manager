package settings

import (
	"gorm.io/gorm"
	"ricr.dev/site-manager/services"
)

type API struct {
	db              *gorm.DB
	settingsService *services.SettingsService
}

func New(db *gorm.DB) *API {
	return &API{
		db:              db,
		settingsService: services.NewSettingsService(),
	}
}
