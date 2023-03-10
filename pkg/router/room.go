package router

import (
	"github.com/gin-gonic/gin"

	"github.com/soicchi/chatapp_backend/pkg/controllers"
	"github.com/soicchi/chatapp_backend/pkg/middleware"
)

func addRoomRoutes(routerGroup *gin.RouterGroup, handler *controllers.Handler) {
	rooms := routerGroup.Group("/rooms")
	rooms.Use(middleware.AuthMiddleware)
	{
		rooms.GET("", handler.GetAllRooms)
		rooms.GET("/:id", handler.GetRoom)
		rooms.POST("/create", handler.CreateRoom)
	}
}
