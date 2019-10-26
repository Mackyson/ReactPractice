package dbUtils

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"os"
)

func GetDB() *gorm.DB {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	config := fmt.Sprintf("%s:%s@(mysql:3306)/%s?parseTime=true",
		user,
		password,
		dbname,
	)
	db, err := gorm.Open("mysql", config) //OpenしまくってもどうせプールされるからOK
	if err != nil {
		log.Printf("failed to connect DB")
		log.Printf(config)
		log.Fatal(err)
	}
	log.Printf("DB connection is opened")

	return db
}

func CloseDB(db *gorm.DB) {
	db.Close()
}
