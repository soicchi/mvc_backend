package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	content string `gorm:"size:255;not null"`
	UserID  uint   `gorm:"not null"`
}
