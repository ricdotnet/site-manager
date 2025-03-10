package sites

import (
	"gorm.io/gorm"
	"ricr.dev/site-manager/repository"
	"ricr.dev/site-manager/services"
)

type SitesAPI struct {
	db           *gorm.DB
	repository   *repository.Repository
	sitesService *services.SitesService
}

func New(db *gorm.DB) *SitesAPI {
	return &SitesAPI{
		db:           db,
		repository:   repository.NewRepository(db),
		sitesService: services.NewSitesService(),
	}
}
