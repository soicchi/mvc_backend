package models

import (
	"github.com/go-ozzo/ozzo-validation"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Content string
	UserID  uint
}

type PostInput struct {
	Content string `json:"content" binding:"required"`
}

func (post *Post) Create(db *gorm.DB) (Post, error) {
	newPost := Post{
		Content: post.Content,
		UserID:  post.UserID,
	}
	result := db.Create(&newPost)

	return newPost, result.Error
}

func FindAllPosts(db *gorm.DB) ([]Post, error) {
	var posts []Post
	result := db.Find(&posts)

	return posts, result.Error
}

func FindPostById(db *gorm.DB, postId uint) (Post, error) {
	var post Post
	result := db.First(&post, postId)

	return post, result.Error
}

func (post *Post) Update(db *gorm.DB, postInput PostInput) (*Post, error) {
	post.Content = postInput.Content
	result := db.Save(&post)

	return post, result.Error
}

func (post *Post) Validate() error {
	err := validation.ValidateStruct(post,
		validation.Field(&post.Content,
			validation.Required.Error("Content is required"),
			validation.Length(1, 255).Error("Content must be between 1 and 255 characters"),
		),
	)
	return err
}
