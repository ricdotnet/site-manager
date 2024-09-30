package repository

import (
	"github.com/charmbracelet/log"
	"gorm.io/gorm"
	"ricr.dev/site-manager/config"
	"ricr.dev/site-manager/models"
)

type SitesRepo struct {
	Db *gorm.DB
}

type Site = models.Site

func (repo *SitesRepo) GetAll(items ...interface{}) error {
	sites := items[0].(*[]Site)
	user := items[1].(*config.JwtCustomClaims)

	if err := repo.Db.Find(&sites, "user_id = ?", user.UserID).Error; err != nil {
		return err
	}

	log.Infof("Found %d site records for user %s", len(*sites), user.Username)

	return nil
}

func (repo *SitesRepo) GetOneByID(id uint, items ...interface{}) error {
	site := items[0].(*Site)
	user := items[1].(*config.JwtCustomClaims)

	if err := repo.Db.First(site, "id = ? AND user_id = ?", id, user.UserID).Error; err != nil {
		return err
	}

	if site == nil {
		return gorm.ErrRecordNotFound
	}

	log.Infof("Found 1 site record with id %d for user %s", id, user.Username)

	return nil
}

func (repo *SitesRepo) CreateOne(item interface{}) error {
	site := item.(*Site)

	if err := repo.Db.Create(site).Error; err != nil {
		return err
	}

	return nil
}

func (repo *SitesRepo) UpdateOne(item interface{}) error {
	site := item.(*Site)

	if err := repo.Db.Updates(site).Error; err != nil {
		return err
	}

	return nil
}

// TODO: maybe merge this with above?
func (repo *SitesRepo) EnableOne(item interface{}) error {
	site := item.(*Site)

	if err := repo.Db.Model(site).
		Where("id = ?", site.ID).
		Update("enabled", site.Enabled).Error; err != nil {
		return err
	}

	return nil
}

func (repo *SitesRepo) DeleteManyByID(sites []uint) error {
	err := repo.Db.Transaction(func(tx *gorm.DB) error {
		for _, site := range sites {
			log.Infof("Deleting site with id %d", site)

			if err := repo.Db.Delete(Site{}, site).Error; err != nil {
				return err
			}
		}

		return nil
	})

	return err
}
