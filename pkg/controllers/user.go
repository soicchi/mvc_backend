package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/soicchi/chatapp_backend/pkg/database"
	"github.com/soicchi/chatapp_backend/pkg/models"
)

func SignUpHandler(c *gin.Context) {
	var signUpInput models.SignUpInput
	err := c.ShouldBind(&signUpInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
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
	db := database.GetDB()
	user, err := newUser.Create(db)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Failed to create user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user":    user,
		"message": "Successfully created user",
	})
}
