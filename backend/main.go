package main

import (
	"net/http"
	"os"
	"strconv"

	"app/controllers"
	"app/models"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router := gin.Default()
	router.Use(gin.Logger())

	// APIのハンドルを定義
	router.POST("/signup", signUp)
	// router.POST("/signin", signIn)
	router.GET("/:uid/", getTasks)
	router.GET("/:uid/:taskid", getTask)

	router.Run(":" + port)
}

//TODO 下のハンドラは別ファイルに

func setCORS(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	c.Header("Access-Control-Max-Age", "86400")
	c.Header("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func signUp(c *gin.Context) {
	c.String(http.StatusOK, "<h1>登録完了！(大嘘)</h1>")
}
func getTasks(c *gin.Context) {
	setCORS(c)

	uid := c.Param("uid")
	uuid, _ := strconv.Atoi(uid)
	tasks := controllers.GetOwnTasks(uuid)
	fmt.Printf("%+v", tasks)
	// var tmptasks []models.Task
	tmptasks := make([]models.Task, 0)
	tmptasks = append(tmptasks, models.Task{Content: uid + "'s task1"}, models.Task{Content: uid + "'s task2"})
	c.JSON(http.StatusOK, tmptasks)
}
func getTask(c *gin.Context) {
	setCORS(c)

	uid := c.Param("uid")
	taskid := c.Param("taskid")
	c.JSON(http.StatusOK, gin.H{"task": uid + "のタスク" + taskid})
}
