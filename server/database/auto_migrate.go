package database

import (
	"github.com/charmbracelet/log"
	"gorm.io/gorm"
	"os"
	"ricr.dev/site-manager/models"
)

func runMigrations(db *gorm.DB) {
	if err := db.AutoMigrate(dbModels()...); err != nil {
		log.Errorf("Failed to run migrations: %s", err.Error())
		os.Exit(1)
	}
	if err := db.Migrator().DropIndex(&models.Site{}, "config_name"); err != nil {
		log.Warnf("Failed to drop index: %s", err.Error())
	}
	if err := db.Migrator().CreateConstraint(&models.LoginCode{}, "fk_login_code_email_user_email"); err != nil {
		log.Warnf("Failed to add foreign key: %s", err.Error())
	}
}

func dbModels() []interface{} {
	return []interface{}{
		&models.User{},
		&models.Site{},
		&models.SiteData{},
		&models.Settings{},
		&models.Session{},
		&models.LoginCode{},
	}
}
