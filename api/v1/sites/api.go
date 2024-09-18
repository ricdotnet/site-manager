package sites

import (
	"github.com/op/go-logging"
	"gorm.io/gorm"
	"ricr.dev/site-manager/config"
	"ricr.dev/site-manager/services"
)

type API struct {
	db           *gorm.DB
	logger       *logging.Logger
	sitesService *services.SitesService
}

func New(db *gorm.DB, cfg *config.Config) *API {
	return &API{
		db:           db,
		logger:       cfg.Logger,
		sitesService: services.NewSitesService(cfg),
	}
}
