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

func (r *Repository) GetAll(user *config.JwtCustomClaims) (*[]Site, error) {
	sites := new([]Site)
	if err := r.db.Find(&sites, "user_id = ?", user.UserID).Error; err != nil {
		return nil, err
	}

	r.logger.Infof("Found %d site records for user %s", len(*sites), user.Username)
	return sites, nil
}

func (r *Repository) GetOne(id int, user *config.JwtCustomClaims) (*Site, error) {
	site := &Site{}
	if err := r.db.First(site, "id = ? AND user_id = ?", id, user.UserID).Error; err != nil {
		return nil, err
	}
	if site != nil {
		r.logger.Infof("Found 1 site record with id %d for user %s", id, user.Username)
	}

	return site, nil
}

func (r *Repository) Create(site *Site) (*Site, error) {
	if err := r.db.Create(site).Error; err != nil {
		return nil, err
	}
	return site, nil
}

func (r *Repository) Update(site *Site) (*Site, error) {
	updated := &Site{}
	var err error
	if err = r.db.Updates(site).Error; err != nil {
		return nil, err
	}
	if err = r.db.First(updated, "id = ?", site.ID).Error; err != nil {
		return nil, err
	}
	return updated, nil
}

func (r *Repository) Enable(site *Site) error {
	if err := r.db.Model(site).
		Where("id = ?", site.ID).
		Update("enabled", site.Enabled).Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) Delete(sites *[]uint) error {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		for _, site := range *sites {
			r.logger.Debugf("Deleting site with id %d", site)

			if err := r.db.Delete(&Site{}, site).Error; err != nil {
				return err
			}
		}

		return nil
	})

	return err
}
