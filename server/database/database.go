package database

import (
	"log"

	"github.com/joaotavioos/cms-server/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var dbError error

func Connect(connectionString string) {

	Instance, dbError = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
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
