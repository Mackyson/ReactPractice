package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Task struct {
	gorm.Model
	TaskOwnerId uint
	Content     string
	Deadline    time.Time
	Status      int
}
