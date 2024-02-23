package sites

import (
	"gorm.io/gorm"
	"ricr.dev/site-manager/services"
)

type API struct {
	repository   *Repository
	sitesService *services.SitesService
}

func New(db *gorm.DB) *API {
	return &API{
		sitesService: services.NewSitesService(),
		repository:   NewRepository(db),
	}
}
