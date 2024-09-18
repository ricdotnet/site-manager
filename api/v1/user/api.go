package user

import (
	"github.com/op/go-logging"
	"gorm.io/gorm"
	"ricr.dev/site-manager/config"
)

type API struct {
	db     *gorm.DB
	logger *logging.Logger
}

func New(db *gorm.DB, cfg *config.Config) *API {
	return &API{
		db:     db,
		logger: cfg.Logger,
	}
}
