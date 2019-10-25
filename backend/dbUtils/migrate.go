package dbUtils

import (
	"github.com/Mackyson/ReactPractice/backend/models"
	// "github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

func Migrate() {
	db := GetDB()
	if !db.HasTable(&models.User{}) {
		err := db.CreateTable(&models.User{})
		if err != nil {
			log.Printf("User table exists")
		}
	}
	if !db.HasTable(&models.Task{}) {
		err := db.CreateTable(&models.Task{})
		if err != nil {
			log.Printf("Task table exists")
		}
	}
	db.LogMode(true)
	db.AutoMigrate(&models.User{}, &models.Task{})
}
