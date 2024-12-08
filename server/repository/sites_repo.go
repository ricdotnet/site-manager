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

func (repo *SitesRepo) GetAll(items ...interface{}) error {
	sites := items[0].(*[]models.Site)
	user := items[1].(*config.Session)

	if err := repo.Db.Order("domain").Find(&sites, "user_id = ?", user.UserID).Error; err != nil {
		return err
	}

	log.Infof("Found %d site records for user %s", len(*sites), user.Username)

	return nil
}

func (repo *SitesRepo) GetOneByID(id uint, items ...interface{}) error {
	site := items[0].(*models.Site)
	user := items[1].(*config.Session)

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
	site := item.(*models.Site)

	return repo.Db.Create(site).Error
}

func (repo *SitesRepo) UpdateOne(item interface{}) error {
	site := item.(*models.Site)

	return repo.Db.Updates(site).Error
}

func (repo *SitesRepo) DeleteManyByID(deleteIds []uint, items ...interface{}) error {
	sites := items[0].(*[]models.Site)

	repo.Db.Find(&sites, `ID in ?`, deleteIds)

	return repo.Db.Where(`ID in ?`, deleteIds).Delete(&models.Site{}).Error
}
