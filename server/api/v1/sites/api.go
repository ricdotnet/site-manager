package sites

import (
	"gorm.io/gorm"
	"ricr.dev/site-manager/repository"
	"ricr.dev/site-manager/services"
)

type SitesAPI struct {
	db           *gorm.DB
	sitesRepo    *repository.SitesRepo
	sitesService *services.SitesService
}

func New(db *gorm.DB) *SitesAPI {
	return &SitesAPI{
		db:           db,
		sitesService: services.NewSitesService(),
		sitesRepo:    repository.NewSitesRepo(db),
	}
}
