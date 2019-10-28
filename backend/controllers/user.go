package controllers

import (
	"github.com/Mackyson/ReactPractice/backend/dbUtils"
	"github.com/Mackyson/ReactPractice/backend/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func IssueSessionID(uid uint64) {} //SessionIDの発行と登録

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
		c.JSON(http.StatusOK, gin.H{"sessionID": "test"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"Error": result.Error})
	}
}
