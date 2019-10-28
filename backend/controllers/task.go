package controllers

import (
	"github.com/Mackyson/ReactPractice/backend/dbUtils"
	"github.com/Mackyson/ReactPractice/backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//TODO sessionIDの確認
func GetOwnTasks(c *gin.Context) {
	uidStr := c.Param("uid")
	uid, _ := strconv.ParseUint(uidStr, 10, 32) //32bit unsigned int(decimal)
	db := dbUtils.GetDB()
	defer dbUtils.CloseDB(db)
	tasks := make([]models.Task, 0)
	db.Find(&tasks, models.Task{TaskOwnerId: uid})
	c.JSON(http.StatusOK, tasks)
}
func AddNewTask(c *gin.Context) {
	uidStr := c.Param("uid")
	uid, _ := strconv.ParseUint(uidStr, 10, 32) //32bit unsigned int(decimal)
	db := dbUtils.GetDB()
	defer dbUtils.CloseDB(db)
	var task models.Task
	c.BindJSON(&task)
	task.Status = 1
	task.TaskOwnerId = uid
	db.Create(&task)
	c.String(http.StatusOK, "")
}
