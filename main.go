package main

import (
	"flag"
	"fmt"
	"github.com/charmbracelet/log"
	"os"
	"time"

	"github.com/ricdotnet/goenvironmental"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	router "ricr.dev/site-manager/api/v1"
	"ricr.dev/site-manager/db"
	"ricr.dev/site-manager/scripts"
)

func main() {
	sa := flag.Bool("sa", false, "map config files to their domains")
	run := flag.Bool("run", false, "start the app")
	envFile := flag.String("env-file", ".env", "environment to run the app in")
	flag.Parse()

	goenvironmental.ParseEnv(*envFile)
	cfg := config.NewConfig()

	if *sa {
		sitesAvailable()
		return
	}

	if *run {
		// TODO: extract db related stuff maybe into the /db dir
		dbHost, _ := goenvironmental.Get("DB_HOST")
		dbUser, _ := goenvironmental.Get("DB_USER")
		dbPass, _ := goenvironmental.Get("DB_PASSWORD")
		dbName, _ := goenvironmental.Get("DB_NAME")

		dns := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
		dbConn, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
		if err != nil {
			panic(err.Error())
		}
		db.RunMigrations(dbConn)

		port, err := goenvironmental.Get("PORT")
		if err != nil {
			panic("A port for the api was not defined")
		}

		// define the echo router and run
		v1 := router.NewRouter(dbConn)
		v1.Logger.Fatal(v1.Start(fmt.Sprintf(":%s", port)))
	}

	println("Nothing to run. -h to see the flags you can run")
}

// TODO: extract this into its own package
func sitesAvailable() {
	start := time.Now()

	sitesMap := scripts.Init().MapSites()
	file, _ := os.Create("db_inserts.sql")

	for k, v := range sitesMap {
		line := fmt.Sprintf("INSERT INTO sites (domain, config_name, user) VALUES ('%s', '%s', %d);\n", v, k, 1)
		// ignore the errors :-)
		_, err := file.WriteString(line)
		if err != nil {
			log.Fatalf("Writing failed while writing sql insert for %s", v)
		}
	}

	// ignore the errors :-)
	_ = file.Close()

	finish := time.Now()
	seconds := float64(finish.UnixMilli()-start.UnixMilli()) / 1000

	fmt.Printf("Finished parsing all config files in %v seconds", seconds)
}
