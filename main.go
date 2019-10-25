package main

import (
	"net/http"
	"os"
	"strconv"

	"github.com/Mackyson/ReactPractice/backend/controllers"
	// "app/dbUtils"
	"fmt"
	"github.com/Mackyson/ReactPractice/backend/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(cors.Default())
	router.Use(static.Serve("/", static.LocalFile("/frontend/public", true)))
	router.NoRoute(func(c *gin.Context) { c.File("/frontend/public/index.html") })

	//dbUtils.Migrate()
	// APIのハンドルを定義
	router.POST("/signup", signUp)
	// router.POST("/signin", signIn)
	router.GET("todo/:uid/", getTasks)
	router.POST("todo/:uid/", controllers.AddNewTask)
	router.GET("todo/:uid/:id", getTask)

	router.Run(":" + port)
}

//TODO 下のハンドラは別ファイルに

func signUp(c *gin.Context) {
	c.String(http.StatusOK, "<h1>登録完了！(大嘘)</h1>")
}
func getTasks(c *gin.Context) {

	uid := c.Param("uid")
	uidInt, _ := strconv.Atoi(uid)
	tasks := controllers.GetOwnTasks(uidInt)
	fmt.Printf("%+v", tasks)
	// var tmptasks []models.Task
	tmptasks := make([]models.Task, 0)
	tmptasks = append(tmptasks, models.Task{Content: uid + "'s task1"}, models.Task{Content: uid + "'s task2"})
	c.JSON(http.StatusOK, tmptasks)
}
func getTask(c *gin.Context) {
	uid := c.Param("uid")
	taskid := c.Param("taskid")
	c.JSON(http.StatusOK, gin.H{"task": uid + "のタスク" + taskid})
}
