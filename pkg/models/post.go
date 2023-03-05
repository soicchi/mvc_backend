package models

import (
	"github.com/go-ozzo/ozzo-validation"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	content string `gorm:"size:255;not null"`
	userID  uint   `gorm:"not null"`
}
