package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/soicchi/chatapp_backend/pkg/database"
	"github.com/soicchi/chatapp_backend/pkg/models"
	"github.com/soicchi/chatapp_backend/pkg/utils"
)

func SignUpHandler(context *gin.Context) {
	var signUpInput models.SignUpInput
	err := context.ShouldBind(&signUpInput)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Invalid request body",
		})
		return
	}

	newUser := &models.User{
		Name:     signUpInput.Name,
		Email:    signUpInput.Email,
		Password: signUpInput.Password,
	}
	err = newUser.Validate()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Invalid request body",
		})
		return
	}

	db := database.GetDB()
	user, err := newUser.Create(db)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Failed to create user",
		})
		return
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Failed to generate token",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"user_id": user.ID,
		"token":   token,
		"message": "Successfully created user",
	})
}

func LoginHandler(context *gin.Context) {
	var loginInput models.LoginInput
	err := context.ShouldBind(&loginInput)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Invalid request body",
		})
		return
	}

	db := database.GetDB()
	user, err := models.FindByEmail(db, loginInput.Email)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Failed to find user",
		})
		return
	}

	if !user.VerifyPassword(loginInput.Password) {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Password is invalid",
		})
		return
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Failed to generate token",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Successfully logged in",
		"token":   token,
	})
}
