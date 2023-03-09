package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/soicchi/chatapp_backend/pkg/models"
	"github.com/soicchi/chatapp_backend/pkg/utils"
)

func (handler *Handler) GetAllRooms(ctx *gin.Context) {
	logger, err := utils.SetupLogger()
	if err != nil {
		panic(err)
	}

	rooms, err := models.FindAllRooms(handler.DB)
	if err != nil {
		logger.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get rooms",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"rooms":   rooms,
		"message": "Rooms fetched successfully",
	})
}

func (handler *Handler) CreateRoom(ctx *gin.Context) {
	logger, err := utils.SetupLogger()
	if err != nil {
		panic(err)
	}

	var inputRoom models.CreateRoomInput
	if err = ctx.ShouldBindJSON(&inputRoom); err != nil {
		logger.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}

	room := &models.Room{
		Name: inputRoom.Name,
	}
	if err = room.Validate(); err != nil {
		logger.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	newRoom, err := room.Create(handler.DB)
	if err != nil {
		logger.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create room",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"room":    newRoom,
		"message": "Room created successfully",
	})
}