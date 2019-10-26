package models

import (
	"github.com/jinzhu/gorm"
)

type Task struct {
	gorm.Model
	TaskOwnerId uint64 `json:"taskOwnerId"`
	Content     string `json:"content"`
	Deadline    string `json:"deadline"`
	Status      int    `json:"status"`
}
