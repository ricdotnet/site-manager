package sites

import (
	"github.com/op/go-logging"
	"gorm.io/gorm"
	"ricr.dev/site-manager/config"
)

type API struct {
	repository *Repository
	logger     *logging.Logger
}

func New(db *gorm.DB, cfg *config.Config) *API {
	return &API{
		repository: NewRepository(db),
		logger:     cfg.Logger,
	}
}
