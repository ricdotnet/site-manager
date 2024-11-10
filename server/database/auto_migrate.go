package database

import (
	"github.com/charmbracelet/log"
	"gorm.io/gorm"
	"ricr.dev/site-manager/models"
)

func runMigrations(db *gorm.DB) {
	if err := db.AutoMigrate(dbModels()...); err != nil {
		panic(err.Error())
	}
	if err := db.Migrator().DropIndex(&models.Site{}, "config_name"); err != nil {
		log.Warnf("Failed to drop index: %s", err.Error())
	}
}

func dbModels() []interface{} {
	return []interface{}{
		&models.User{},
		&models.Site{},
		&models.SiteData{},
		&models.Settings{},
		&models.Session{},
	}
}
