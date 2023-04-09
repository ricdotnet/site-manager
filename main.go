package main

import (
	"github.com/ricdotnet/goenvironmental"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	router "ricr.dev/site-manager/api/v1"
	"ricr.dev/site-manager/api/v1/sites"
	"ricr.dev/site-manager/api/v1/user"
	"ricr.dev/site-manager/config"
)

func main() {
	goenvironmental.ParseEnv()

	cfg := config.NewConfig()

	dbString, _ := goenvironmental.Get("DB_STRING")
	db, err := gorm.Open(mysql.Open(dbString), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}
	err = db.AutoMigrate(&user.User{}, &sites.Site{})
	if err != nil {
		return
	}

	v1 := router.New(cfg, db)
	v1.Logger.Fatal(v1.Start(":4000"))
}
