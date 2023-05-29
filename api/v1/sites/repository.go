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

func (r *Repository) GetAll(user uint) (*[]Site, error) {
	sites := new([]Site)
	if err := r.db.Find(&sites, "user_id = ?", user).Error; err != nil {
		return nil, err
	}

	r.logger.Infof("Found %d site records for user with id %d", len(*sites), user)
	return sites, nil
}

func (r *Repository) GetOne(id int, user uint) (*Site, error) {
	site := &Site{}
	if err := r.db.First(site, "id = ?", id).Error; err != nil {
		return nil, err
	}
	if site != nil {
		r.logger.Infof("Found 1 site record with id of %d for user %d", id, user)
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

func (r *Repository) Delete(id int) error {
	if err := r.db.Delete(&Site{}, id).Error; err != nil {
		return err
	}

	return nil
}
