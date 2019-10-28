package main

import (
	"net/http"
	"os"

	"github.com/Mackyson/ReactPractice/backend/controllers"
	"github.com/Mackyson/ReactPractice/backend/dbUtils"
	// "github.com/Mackyson/ReactPractice/backend/models"
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
	router.Use(static.Serve("/", static.LocalFile("./frontend/public/", true)))
	router.NoRoute(func(c *gin.Context) { c.File("./frontend/public/index.html") })

	dbUtils.Migrate()
	// APIのハンドルを定義
	//TODO API handler Grouping
	//TODO ハンドラの設定を別ファイルに切り出す
	router.POST("/signup", controllers.AddNewUser)
	router.POST("/signin", controllers.UpdateSessionID)
	router.GET("todo/:uid/", controllers.GetOwnTasks)
	router.POST("todo/:uid/", controllers.AddNewTask)
	router.GET("todo/:uid/:id", getTask)

	router.Run(":" + port)
}

//TODO 下のハンドラは別ファイルに

func signUp(c *gin.Context) {
	c.String(http.StatusOK, "<h1>登録完了！(大嘘)</h1>")
}
func getTask(c *gin.Context) {
	uid := c.Param("uid")
	taskid := c.Param("taskid")
	c.JSON(http.StatusOK, gin.H{"task": uid + "のタスク" + taskid})
}
