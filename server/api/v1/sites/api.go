package sites

import (
	"gorm.io/gorm"
	"ricr.dev/site-manager/repository"
	"ricr.dev/site-manager/services"
)

type SitesAPI struct {
	db              *gorm.DB
	repo            repository.SitesRepository
	sitesService    *services.SitesService
	commandsService *services.CommandsService
}

func New(db *gorm.DB) *SitesAPI {
	return &SitesAPI{
		db:              db,
		repo:            &repository.SitesRepo{Db: db},
		sitesService:    services.NewSitesService(),
		commandsService: services.NewCommandsService(db),
	}
}
