package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/soicchi/chatapp_backend/pkg/models"
	"github.com/soicchi/chatapp_backend/pkg/utils"
)

var cookieMaxAge int = 60 * 60 * 24 * 30

func (handler *Handler) SignUpHandler(ctx *gin.Context) {
	var signUpInput models.SignUpInput
	err := ctx.ShouldBind(&signUpInput)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Invalid request body",
		})
		return
	}

	user := &models.User{
		Name:     signUpInput.Name,
		Email:    signUpInput.Email,
		Password: signUpInput.Password,
	}
	err = user.Validate()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Invalid request body",
		})
		return
	}

	newUser, err := user.Create(handler.DB)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Failed to create user",
		})
		return
	}

	token, err := utils.GenerateToken(newUser.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Failed to generate token",
		})
		return
	}
	ctx.SetCookie("token", token, cookieMaxAge, "/", "localhost", false, true)

	ctx.JSON(http.StatusOK, gin.H{
		"user_id": newUser.ID,
		"message": "Successfully created user",
	})
}

func (handler *Handler) LoginHandler(ctx *gin.Context) {
	var loginInput models.LoginInput
	err := ctx.ShouldBind(&loginInput)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Invalid request body",
		})
		return
	}

	user, err := models.FindUserByEmail(handler.DB, loginInput.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Failed to find user",
		})
		return
	}

	if !user.VerifyPassword(loginInput.Password) {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Password is invalid",
		})
		return
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Failed to generate token",
		})
		return
	}
	ctx.SetCookie("token", token, cookieMaxAge, "/", "localhost", false, true)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Successfully logged in",
	})
}
