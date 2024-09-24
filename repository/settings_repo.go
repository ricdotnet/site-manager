package repository

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"ricr.dev/site-manager/models"
)

type SettingsRepo struct {
	db *gorm.DB
}

func NewSettingsRepo(db *gorm.DB) *SettingsRepo {
	return &SettingsRepo{
		db: db,
	}
}

func (repo *SettingsRepo) FindAll() {}

func (repo *SettingsRepo) FindFirst(setting *models.Settings, key string) error {
	if err := repo.db.First(setting, "`key` = ?", key).Error; err != nil {
		return err
	}

	return nil
}

func (repo *SettingsRepo) CreateOrUpdateOne(setting *models.Settings) error {
	if err := repo.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "key"}},
		DoUpdates: clause.AssignmentColumns([]string{"value"}),
	}).Create(setting).Error; err != nil {
		return err
	}

	return nil
}

func (repo *SettingsRepo) DeleteOne(key string) error {
	if err := repo.db.Delete(&models.Settings{}, "`key` = ?", key).Error; err != nil {
		return err
	}

	return nil
}
