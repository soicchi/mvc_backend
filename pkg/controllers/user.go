package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/soicchi/chatapp_backend/pkg/models"
)

func (handler *Handler) GetUsers(ctx *gin.Context) {
	users, err := models.FindAllUsers(handler.DB)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Failed to get users",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func (handler *Handler) GetUser(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Invalid user id",
		})
		return
	}

	user, err := models.FindUserById(handler.DB, uint(userId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Failed to get user",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func (handler *Handler) UpdateUser(ctx *gin.Context) {
	var updateInput models.UpdateUserInput
	if err := ctx.ShouldBind(&updateInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Invalid request body",
		})
		return
	}

	userId := ctx.GetUint("userId")
	user, err := models.FindUserById(handler.DB, userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Failed to get user",
		})
		return
	}

	updatedUser, err := user.Update(handler.DB, updateInput)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Failed to update user",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user": updatedUser,
	})
}

func (handler *Handler) DeleteUser(ctx *gin.Context) {
	userId := ctx.GetUint("userId")
	user, err := models.FindUserById(handler.DB, userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Failed to get user",
		})
		return
	}

	if err = user.Delete(handler.DB); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Failed to delete user",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User deleted",
	})
}
