package settings

import (
	"gorm.io/gorm/clause"
)

func (api *API) findAll() {}

func (api *API) findFirst(setting *Setting, key string) error {
	if err := api.db.First(setting, "`key` = ?", key).Error; err != nil {
		return err
	}

	return nil
}

func (api *API) insertOrUpdate(setting *Setting) error {
	if err := api.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "key"}},
		DoUpdates: clause.AssignmentColumns([]string{"value"}),
	}).Create(setting).Error; err != nil {
		return err
	}

	return nil
}

func (api *API) delete(key string) error {
	if err := api.db.Delete(&Setting{}, "`key` = ?", key).Error; err != nil {
		return err
	}

	return nil
}
