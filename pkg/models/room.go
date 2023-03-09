package models

import (
	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	Name string
	Post []Post `gom:"foreignKey:RoomID"`
}