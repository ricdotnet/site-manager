package database

import (
	"gorm.io/gorm"
	"ricr.dev/site-manager/models"
)

func runMigrations(db *gorm.DB) {
	if err := db.AutoMigrate(dbModels()...); err != nil {
		panic(err.Error())
	}
}

func dbModels() []interface{} {
	return []interface{}{
		&models.User{},
		&models.Site{},
		&models.SiteData{},
		&models.Settings{},
	}
}
