package domains

import (
	"gorm.io/gorm"
	"ricr.dev/site-manager/repository"
	"ricr.dev/site-manager/services"
)

type DomainsAPI struct {
	repository     *repository.Repository
	domainsService *services.DomainsService
}

func New(db *gorm.DB) *DomainsAPI {
	return &DomainsAPI{
		repository:     repository.NewRepository(db),
		domainsService: services.NewDomainsService(db),
	}
}
