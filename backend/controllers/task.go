package controllers

import (
	"github.com/Mackyson/ReactPractice/backend/dbUtils"
	"github.com/Mackyson/ReactPractice/backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetOwnTasks(id int) models.Task {
	db := dbUtils.GetDB()
	defer dbUtils.CloseDB(db)
	table := models.Task{}
	db.First(&table, id)
	return table
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
