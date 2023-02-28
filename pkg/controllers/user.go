package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/soicchi/chatapp_backend/pkg/models"
)

func (handler *Handler) GetUsers(context *gin.Context) {
	users, err := models.FindAllUsers(handler.DB)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
			"message": "Failed to get users",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func (handler *Handler) GetUser(context *gin.Context) {
	userId, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"message": "Invalid user id",
		})
		return
	}

	user, err := models.FindUserById(handler.DB, userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
			"message": "Failed to get user",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}