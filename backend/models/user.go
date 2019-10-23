package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	// Id        uint      `json:"id" gorm:"primary_key"`
	// CreatedAt time.Time `json:"createdAt" gorm:"index"`
	// UpdatedAt time.Time `json:"updatedAt" gorm:"index"`
	// DeletedAt time.Time `json:"deletedAt" gorm:"index"`
	Name     string `json:"name" gorm:"index"`
	Password string `json:"password" gorm:"index"`
	Token    string `json:"token" gorm:"index"`
}
