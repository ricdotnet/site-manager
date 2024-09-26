package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/log"
	"github.com/ricdotnet/goenvironmental"
	router "ricr.dev/site-manager/api/v1"
	"ricr.dev/site-manager/database"
	"ricr.dev/site-manager/scripts"
)

func main() {
	sa := flag.Bool("sa", false, "map config files to their domains")
	run := flag.Bool("run", false, "start the app")
	envFile := flag.String("env-file", "", "environment to run the app in")
	flag.Parse()

	// parse the env file if it was provided... goenvironmental will still provide env vars from the system through os.Getenv
	if *envFile != "" {
		goenvironmental.ParseEnv(*envFile)
	} else {
		goenvironmental.ParseEnv()
	}

	if *sa {
		sitesAvailable()
		return
	}

	if *run {
		conn, err := database.Connect()
		if err != nil {
			panic(fmt.Sprintf("Could not connect to the database: %s", err.Error()))
		}

		port, err := goenvironmental.Get("PORT")
		if err != nil {
			panic("A port for the api was not defined")
		}

		// define the echo router and run
		v1 := router.NewRouter(conn)
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
