// TODO: I will test this type of migrations, if I don't like I will convert all this into normal sql files and run on every launch

package db

import (
	"gorm.io/gorm"
	"ricr.dev/site-manager/models"
)

func RunMigrations(db *gorm.DB) {
	if err := db.AutoMigrate(dbModels()...); err != nil {
		panic(err.Error())
	}
}

func dbModels() []interface{} {
	return []interface{}{
		&models.User{},
		&models.Site{},
		&models.SiteData{},
	}
}
