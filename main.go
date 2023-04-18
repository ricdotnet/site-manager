package main

import (
	"fmt"
	"github.com/op/go-logging"
	"github.com/ricdotnet/goenvironmental"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	router "ricr.dev/site-manager/api/v1"
	"ricr.dev/site-manager/api/v1/sites"
	"ricr.dev/site-manager/api/v1/user"
	"ricr.dev/site-manager/config"
	"ricr.dev/site-manager/scripts"
)

func main() {
	goenvironmental.ParseEnv()
	cfg := config.NewConfig()

	command := os.Args[1]
	if command == "sa" {
		sitesAvailable(cfg.Logger)
		return
	}

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

func sitesAvailable(l *logging.Logger) {
	sitesMap := scripts.Init(l).MapSites()
	file, _ := os.Create("file to write")

	for k, v := range sitesMap {
		line := fmt.Sprintf("INSERT INTO sites (domain, config_name, user) VALUES ('%s', '%s', %d);\n", v, k, 1)
		// ignore the errors :-)
		_, _ = file.WriteString(line)
	}

	// ignore the errors :-)
	_ = file.Close()
}
