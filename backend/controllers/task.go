package controllers

import (
	"app/db"
	"app/models"
)

func GetOwnTasks(id int) models.Task {
	db.Init()
	defer db.CloseDB()
	db := db.GetDB()
	table := models.Task{}
	db.First(&table, id)
	return table
}
