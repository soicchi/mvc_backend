package controllers

import (
	"net/http"
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