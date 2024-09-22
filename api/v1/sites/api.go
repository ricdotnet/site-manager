package sites

import (
	"gorm.io/gorm"
	"ricr.dev/site-manager/services"
)

type API struct {
	db           *gorm.DB
	sitesService *services.SitesService
}

func New(db *gorm.DB) *API {
	return &API{
		db:           db,
		sitesService: services.NewSitesService(),
	}
}
