package models

import (
	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	Name string
	Posts []Post `gorm:"foreignKey:RoomID"`
}

func FindAllRooms(db *gorm.DB) ([]Room, error) {
	var rooms []Room
	result := db.Find(&rooms)

	return rooms, result.Error
}