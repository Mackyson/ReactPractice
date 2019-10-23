package db

import (
	// "github.com/ReactPractice/backend/models"
	"app/models"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"os"
)

var (
	db  *gorm.DB
	err error
)

func Init() {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	config := fmt.Sprintf("%s:%s@(mysql:3306)/%s",
		user,
		password,
		dbname,
	)
	db, err = gorm.Open("mysql", config)
	if err != nil {
		log.Printf("failed to connect DB")
		log.Printf(config)
		log.Fatal(err)
	}
	log.Printf("DB connection is opened")

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

func GetDB() *gorm.DB {
	return db
}
func CloseDB() {
	db.Close()
}
