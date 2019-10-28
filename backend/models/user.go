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
	Name      string `json:"name" gorm:"unique;not null"`
	Password  string `json:"password" gorm:"not null"`
	SessionID string `json:"session_id" gorm:"index"`
}
