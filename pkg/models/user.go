package models

import (
	"gorm.io/gorm"

	"github.com/soicchi/chatapp_backend/pkg/encrypt"
)

type User struct {
	gorm.Model
	Name string `gorm:"size:255;not null;`
	Email string `gorm:"size:255;not null;unique`
	Password string `gorm:"size:255;not null;`
}

type SignUpInput struct {
	Name string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u *User) Create(db *gorm.DB) (User, error) {
	user := User{
		Name: u.Name,
		Email: u.Email,
		Password: encrypt.Encrypt(u.Password),
	}
	result := db.Create(&user)

	return user, result.Error
}
