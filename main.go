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
	"time"
)

func main() {
	goenvironmental.ParseEnv()
	cfg := config.NewConfig()

	if len(os.Args) > 1 {
		command := os.Args[1]
		if command == "sa" {
			sitesAvailable(cfg.Logger)
			return
		}
	}

	dbString, _ := goenvironmental.Get("DB_STRING")
	db, err := gorm.Open(mysql.Open(dbString), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}
	err = db.AutoMigrate(dbModels()...)
	if err != nil {
		return
	}

	v1 := router.New(cfg, db)
	v1.Logger.Fatal(v1.Start(":4000"))
}

func dbModels() []interface{} {
	return []interface{}{
		&user.User{},
		&sites.Site{},
	}
}

func sitesAvailable(l *logging.Logger) {
	start := time.Now()

	sitesMap := scripts.Init(l).MapSites()
	file, _ := os.Create("db_inserts.sql")

	for k, v := range sitesMap {
		line := fmt.Sprintf("INSERT INTO sites (domain, config_name, user) VALUES ('%s', '%s', %d);\n", v, k, 1)
		// ignore the errors :-)
		_, err := file.WriteString(line)
		if err != nil {
			l.Fatalf("Writing failed while writing sql insert for %s", v)
		}
	}

	// ignore the errors :-)
	_ = file.Close()

	finish := time.Now()
	seconds := float64(finish.UnixMilli()-start.UnixMilli()) / 1000

	fmt.Printf("Finished parsing all config files in %v seconds", seconds)
}
