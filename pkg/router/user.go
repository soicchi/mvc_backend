package router

import (
	"github.com/gin-gonic/gin"

	"github.com/soicchi/chatapp_backend/pkg/middleware"
	"github.com/soicchi/chatapp_backend/pkg/controllers"
)

func addUserRoutes(routerGroup *gin.RouterGroup, handler *controllers.Handler) {
	users := routerGroup.Group("/users")
	users.Use(middleware.AuthMiddleware)
	{
		users.GET("", handler.GetUsers)
		users.GET("/:id", handler.GetUser)
		users.PUT("/:id", handler.UpdateUser)
	}
}
