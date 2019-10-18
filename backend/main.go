package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router := gin.Default()
	router.Use(gin.Logger())

	router.POST("/signup", signup)
	router.GET("/:uid/", getTasks)
	router.GET("/:uid/:taskid", getTask)
}

//TODO 以下のハンドラは別のところに行く
func signup(c *gin.Context) {
	c.String(http.StatusOK, "<h1>登録完了！(大嘘)</h1>")
}
func getTasks(c *gin.Context) {
	uid := c.Param("uid")
	c.JSON(http.StatusOK, gin.H{"tasks": uid + "のタスク一覧"})
}
func getTask(c *gin.Context) {
	uid := c.Param("uid")
	taskid := c.Param("taskid")
	c.JSON(http.StatusOK, gin.H{"task": uid + "のタスク" + taskid})
}
