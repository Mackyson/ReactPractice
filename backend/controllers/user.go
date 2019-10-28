package controllers

import (
	"github.com/Mackyson/ReactPractice/backend/dbUtils"
	"github.com/Mackyson/ReactPractice/backend/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	// "log"
	"net/http"
)

func IssueSessionID(db *gorm.DB, user models.User) string { //SessionIDの発行と登録
	sessionID := "test"
	db.Model(&user).Update("session_id", sessionID)
	return sessionID
}

func HashPassword(raw string) string { //パスワードのハッシュ化
	return raw
}

func UpdateSessionID(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		panic(err) //TODO まともなハンドリングをする
	}
	name := c.PostForm("name")
	rawPassword := c.PostForm("password")
	db := dbUtils.GetDB()
	defer dbUtils.CloseDB(db)
	hashedPassword := HashPassword(rawPassword)
	user := models.User{}
	result := db.Where("name = ? AND name !=\"\"", name).First(&user)
	//存在するか->パスは一致しているか
	if result.Error == nil {
		if user.Password != hashedPassword {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Wrong password"})
		} else {
			c.JSON(http.StatusOK, gin.H{"sessionID": IssueSessionID(db, user)})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not exists such user"})
	}
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
	if result.Error == nil {
		c.JSON(http.StatusOK, gin.H{"sessionID": IssueSessionID(db, user)})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The user with the same name already exists"})
	}
}
