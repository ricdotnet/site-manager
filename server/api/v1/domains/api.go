package domains

import (
	"gorm.io/gorm"
	"ricr.dev/site-manager/services"
)

type DomainsAPI struct {
	domainsService *services.DomainsService
}

func New(db *gorm.DB) *DomainsAPI {
	return &DomainsAPI{
		domainsService: services.NewDomainsService(db),
	}
}
