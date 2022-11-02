package database

import (
	"log"

	"github.com/joaotavioos/cms-server/entity"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var dbError error

func Connect(connectionString string) {

	if connectionString == "" {
		dbUri := ":memory:"

		Instance, dbError = gorm.Open(sqlite.Open(dbUri), &gorm.Config{})

	} else {
		Instance, dbError = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	}

	if dbError != nil {
		log.Fatal(dbError)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database!")
}

func Migrate() {
	Instance.AutoMigrate(&entity.UserMysql{}, &entity.PostMysql{}, &entity.PageMysql{})
	log.Println("Database Migration Completed!")
}
