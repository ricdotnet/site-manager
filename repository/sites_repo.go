package repository

import (
	"github.com/charmbracelet/log"
	"gorm.io/gorm"
	"ricr.dev/site-manager/config"
	"ricr.dev/site-manager/models"
)

type SitesRepo struct {
	db *gorm.DB
}

func NewSitesRepo(db *gorm.DB) *SitesRepo {
	return &SitesRepo{
		db: db,
	}
}

func (repo *SitesRepo) FindAll(user *config.JwtCustomClaims) (*[]models.Site, error) {
	sites := new([]models.Site)
	if err := repo.db.Find(&sites, "user_id = ?", user.UserID).Error; err != nil {
		return nil, err
	}

	log.Infof("Found %d site records for user %s", len(*sites), user.Username)
	return sites, nil
}

func (repo *SitesRepo) FindFirst(id int, user *config.JwtCustomClaims) (*models.Site, error) {
	site := &models.Site{}
	if err := repo.db.First(site, "id = ? AND user_id = ?", id, user.UserID).Error; err != nil {
		return nil, err
	}
	if site != nil {
		log.Infof("Found 1 site record with id %d for user %s", id, user.Username)
	}

	return site, nil
}

func (repo *SitesRepo) InsertOne(site *models.Site) (*models.Site, error) {
	if err := repo.db.Create(site).Error; err != nil {
		return nil, err
	}
	return site, nil
}

func (repo *SitesRepo) UpdateOne(site *models.Site) (*models.Site, error) {
	updated := &models.Site{}
	var err error
	if err = repo.db.Updates(site).Error; err != nil {
		return nil, err
	}
	if err = repo.db.First(updated, "id = ?", site.ID).Error; err != nil {
		return nil, err
	}
	return updated, nil
}

// TODO: maybe merge this with above?
func (repo *SitesRepo) EnableOne(site *models.Site) error {
	if err := repo.db.Model(site).
		Where("id = ?", site.ID).
		Update("enabled", site.Enabled).Error; err != nil {
		return err
	}

	return nil
}

func (repo *SitesRepo) Delete(sites *[]uint) error {
	err := repo.db.Transaction(func(tx *gorm.DB) error {
		for _, site := range *sites {
			log.Infof("Deleting site with id %d", site)

			if err := repo.db.Delete(&models.Site{}, site).Error; err != nil {
				return err
			}
		}

		return nil
	})

	return err
}
