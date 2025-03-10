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

func (repo *SettingsRepo) GetAll(item ...interface{}) error {
	settings := item[0].(*[]Setting)
	userId := item[1].(uint)

	if err := repo.Db.Find(settings, "user_id = ?", userId).Error; err != nil {
		return err
	}

	return nil
}

func (repo *SettingsRepo) GetOne(key string, item ...interface{}) error {
	setting := item[0].(*Setting)
	userId := item[1].(uint)

	if err := repo.Db.First(setting, "`key` = ? and `user_id` = ?", key, userId).Error; err != nil {
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

func (repo *SettingsRepo) DeleteOne(opts ...interface{}) error {
	key := opts[0].(string)
	userId := opts[1].(uint)

	return repo.Db.Delete(&Setting{}, "`key` = ? and `user_id` = ?", key, userId).Error
}
