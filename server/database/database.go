package database

import (
	"fmt"

	"github.com/ricdotnet/goenvironmental"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	host, _ := goenvironmental.Get("DB_HOST")
	user, _ := goenvironmental.Get("DB_USER")
	pass, _ := goenvironmental.Get("DB_PASSWORD")
	name, _ := goenvironmental.Get("DB_NAME")

	dns := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, name)
	conn, err := gorm.Open(mysql.Open(dns), &gorm.Config{})

	runMigrations(conn)

	return conn, err
}
