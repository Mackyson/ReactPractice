package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

//下の構造体は別ファイルに
type Task struct {
	Content string `json:"content"`
}

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
	router.Run(":" + port)
}

//TODO 下のハンドラは別ファイルに

func setCORS(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	c.Header("Access-Control-Max-Age", "86400")
	c.Header("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func signup(c *gin.Context) {
	c.String(http.StatusOK, "<h1>登録完了！(大嘘)</h1>")
}
func getTasks(c *gin.Context) {
	setCORS(c)

	uid := c.Param("uid")
	var tasks []Task
	tasks = make([]Task, 0)
	tasks = append(tasks, Task{uid + "'s task1"}, Task{uid + "'s task2"})
	c.JSON(http.StatusOK, tasks)
}
func getTask(c *gin.Context) {
	setCORS(c)

	uid := c.Param("uid")
	taskid := c.Param("taskid")
	c.JSON(http.StatusOK, gin.H{"task": uid + "のタスク" + taskid})
}
