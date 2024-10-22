package sites

import (
	"gorm.io/gorm"
	"ricr.dev/site-manager/repository"
	"ricr.dev/site-manager/services"
)

type SitesAPI struct {
	db           *gorm.DB
	repo         repository.SitesRepository
	sitesService *services.SitesService
}

func New(db *gorm.DB) *SitesAPI {
	return &SitesAPI{
		repo:         &repository.SitesRepo{Db: db},
		sitesService: services.NewSitesService(),
	}
}
