package models

import (
	"regexp"

	"gorm.io/gorm"
	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"

	"github.com/soicchi/chatapp_backend/pkg/utils"
)

type User struct {
	gorm.Model
	Name     string `gorm:"size:255;not null"`
	Email    string `gorm:"size:255;not null;unique"`
	Password string `gorm:"size:255;not null"`
}

type SignUpInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// type LoginInput struct {
// 	Email	string `json:"email" binding:"required"`
// 	Password string `json:"password" binding:"required"`
// }

func (u *User) Create(db *gorm.DB) (User, error) {
	user := User{
		Name:     u.Name,
		Email:    u.Email,
		Password: utils.Encrypt(u.Password),
	}
	result := db.Create(&user)

	return user, result.Error
}

// func FindByEmail(db *gorm.DB, password string) (User, error) {
// 	var user User
// 	result := db.Where("email = ?", password).First(&user)

// 	return user, result.Error
// }

// func (u *User) CheckPassword(password string) bool {
// 	return u.Password == utils.Encrypt(password)
// }

func (u *User) Validate() error {
	err := validation.ValidateStruct(u,
		validation.Field(&u.Name,
			validation.Required.Error("Name is required"),
			validation.Length(1, 255).Error("Name is too long"),
		),
		validation.Field(&u.Email,
			validation.Required.Error("Email is required"),
			is.Email.Error("invalid format"),
		),
		validation.Field(&u.Password,
			validation.Required.Error("Password is required"),
			validation.Length(8, 255).Error("less than 7 chars or more than 256 chars"),
		),
	)
	return err
}

func (u *User) Validate() error {
	passwordRegex := "^[A-Za-z0-9]*[A-Z][A-Za-z0-9]*[a-z][A-Za-z0-9]*\\d[A-Za-z0-9]*$"
	err := validation.ValidateStruct(u,
		validation.Field(&u.Name,
			validation.Required.Error("Name is required"),
			validation.Length(1, 255).Error("Name is too long"),
		),
		validation.Field(&u.Email,
			validation.Required.Error("Email is required"),
			is.Email.Error("invalid format"),
		),
		validation.Field(&u.Password,
			validation.Required.Error("Password is required"),
			validation.Length(8, 255).Error("less than 7 chars or more than 256 chars"),
			validation.Match(regexp.MustCompile(passwordRegex)).Error("invalid format"),
		),
	)
	return err
}

func (u *User) Create(db *gorm.DB) (User, error) {
	user := User{
		Name:     u.Name,
		Email:    u.Email,
		Password: encrypt.Encrypt(u.Password),
	}
	result := db.Create(&user)

	return user, result.Error
}

func (u *User) Validate() error {
	passwordRegex := "^[A-Za-z0-9]*[A-Z][A-Za-z0-9]*[a-z][A-Za-z0-9]*\\d[A-Za-z0-9]*$"
	err := validation.ValidateStruct(u,
		validation.Field(&u.Name,
			validation.Required.Error("Name is required"),
			validation.Length(1, 255).Error("Name is too long"),
		),
		validation.Field(&u.Email,
			validation.Required.Error("Email is required"),
			is.Email.Error("invalid format"),
		),
		validation.Field(&u.Password,
			validation.Required.Error("Password is required"),
			validation.Length(8, 255).Error("less than 7 chars or more than 256 chars"),
			validation.Match(regexp.MustCompile(passwordRegex)).Error("invalid format"),
		),
	)
	return err
}