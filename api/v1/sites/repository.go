package sites

import (
	"github.com/op/go-logging"
	"gorm.io/gorm"
	"ricr.dev/site-manager/config"
)

type Repository struct {
	db     *gorm.DB
	logger *logging.Logger
}

func NewRepository(db *gorm.DB, cfg *config.Config) *Repository {
	return &Repository{
		db:     db,
		logger: cfg.Logger,
	}
}

func (r *Repository) GetAll(user int) *[]Site {
	sites := new([]Site)
	r.db.Find(&sites, "user = ?", user)
	r.logger.Infof("Found %d site records for user with id %d", len(*sites), user)

	return sites
}

func (r *Repository) GetOne(id int, user int) *Site {
	site := &Site{}
	r.db.First(site, "id = ?", id)
	if site != nil {
		r.logger.Info("Found 1 site record with id of %d for user %d", id, user)
	}

	return site
}

func (r *Repository) Create(site *Site) *Site {
	r.db.Create(site)
	return site
}

func (r *Repository) Update(site *Site) *Site {
	r.db.Updates(site)
	return site
}

func (r *Repository) Delete(id int) int {
	r.db.Delete(&Site{}, id)
	return id
}
