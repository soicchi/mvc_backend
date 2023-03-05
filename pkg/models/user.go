package models

import (
	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"gorm.io/gorm"

	"github.com/soicchi/chatapp_backend/pkg/utils"
)

type User struct {
	gorm.Model
	Name     string `gorm:"size:255;not null"`
	Email    string `gorm:"size:255;not null;unique"`
	Password string `gorm:"size:255;not null"`
	Posts    []Post
}

type SignUpInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateUserInput struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}

func (u *User) Create(db *gorm.DB) (User, error) {
	user := User{
		Name:     u.Name,
		Email:    u.Email,
		Password: utils.Encrypt(u.Password),
	}
	result := db.Create(&user)

	return user, result.Error
}

func FindUserByEmail(db *gorm.DB, email string) (User, error) {
	var user User
	result := db.Where("email = ?", email).First(&user)

	return user, result.Error
}

func FindAllUsers(db *gorm.DB) ([]User, error) {
	var users []User
	result := db.Find(&users)

	return users, result.Error
}

func FindUserById(db *gorm.DB, id uint) (User, error) {
	var user User
	result := db.Where("id = ?", id).First(&user)

	return user, result.Error
}

func (user *User) Update(db *gorm.DB, updateInput UpdateUserInput) (*User, error) {
	user.Name = updateInput.Name
	user.Email = updateInput.Email
	result := db.Save(&user)

	return user, result.Error
}

func (user *User) Delete(db *gorm.DB) error {
	result := db.Delete(&user)

	return result.Error
}

// TODO: 下記メソッド名をSignUpValidateに変更
func (user *User) Validate() error {
	err := validation.ValidateStruct(user,
		validation.Field(&user.Name,
			validation.Required.Error("Name is required"),
			validation.Length(1, 255).Error("Name is too long"),
		),
		validation.Field(&user.Email,
			validation.Required.Error("Email is required"),
			is.Email.Error("invalid format"),
		),
		validation.Field(&user.Password,
			validation.Required.Error("Password is required"),
			validation.Length(8, 255).Error("less than 7 chars or more than 256 chars"),
		),
	)
	return err
}

func (u *User) VerifyPassword(inputPassword string) bool {
	return u.Password == utils.Encrypt(inputPassword)
}
