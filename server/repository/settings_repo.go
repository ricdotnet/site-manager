package repository

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"ricr.dev/site-manager/models"
)

type SettingsRepo struct {
	Db *gorm.DB
}

type Setting = models.Settings

func (repo *SettingsRepo) GetOne(key string, item ...interface{}) error {
	setting := item[0].(*Setting)

	if err := repo.Db.First(setting, "`key` = ?", key).Error; err != nil {
		return err
	}

	return nil
}

func (repo *SettingsRepo) CreateOrUpdateOne(item interface{}) (interface{}, error) {
	setting := item.(*Setting)

	if err := repo.Db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "key"}},
		DoUpdates: clause.AssignmentColumns([]string{"value"}),
	}).Create(setting).Error; err != nil {
		return nil, err
	}

	return setting, nil
}

func (repo *SettingsRepo) DeleteOne(key string) error {
	if err := repo.Db.Delete(&Setting{}, "`key` = ?", key).Error; err != nil {
		return err
	}

	return nil
}
