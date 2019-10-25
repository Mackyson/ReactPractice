package models

import (
	"github.com/jinzhu/gorm"
)

type Task struct {
	gorm.Model
	TaskOwnerId uint64
	Content     string
	Deadline    string
	Status      int
}
