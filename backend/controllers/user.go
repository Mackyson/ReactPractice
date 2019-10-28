package controllers

import (
	"github.com/Mackyson/ReactPractice/backend/dbUtils"
	"github.com/Mackyson/ReactPractice/backend/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

func IssueSessionID(db *gorm.DB, name string) string { //SessionIDの発行と登録
	sessionID := "test"
	userBefore := models.User{Name: name}
	db.First(&userBefore)
	log.Printf("%+v", userBefore)
	db.Model(&userBefore).Update("session_id", sessionID)
	return sessionID
}

func HashPassword(raw string) string { //パスワードのハッシュ化
	return raw
}

func AddNewUser(c *gin.Context) {

	err := c.Request.ParseForm()
	if err != nil {
		panic(err) //TODO まともなハンドリングをする
	}
	name := c.PostForm("name")
	rawPassword := c.PostForm("password")
	db := dbUtils.GetDB()
	defer dbUtils.CloseDB(db)
	user := models.User{Name: name, Password: HashPassword(rawPassword)}
	result := db.Create(&user)
	log.Printf("%+v", result)
	if result.Error == nil {
		c.JSON(http.StatusOK, gin.H{"sessionID": IssueSessionID(db, name)})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"Error": result.Error})
	}
}
