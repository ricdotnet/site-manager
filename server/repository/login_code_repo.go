package repository

import (
	"gorm.io/gorm"
	"ricr.dev/site-manager/models"
)

type LoginCodeRepo struct {
	Db *gorm.DB
}

func (repo *LoginCodeRepo) CreateOne(code interface{}) error {
	return repo.Db.Create(code).Error
}

func (repo *LoginCodeRepo) GetOne(code string, items ...interface{}) error {
	loginCode := items[0].(*models.LoginCode)
	return repo.Db.Where(&models.LoginCode{Code: code}).First(loginCode).Error
}

func (repo *LoginCodeRepo) DeleteOne(opts ...interface{}) error {
	code := opts[0].(string)
	return repo.Db.Delete(&models.LoginCode{}, "`code` = ?", code).Error
}
