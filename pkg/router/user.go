package router

import (
	"github.com/gin-gonic/gin"

	"github.com/soicchi/chatapp_backend/pkg/controllers"
	"github.com/soicchi/chatapp_backend/pkg/middleware"
)

func addUserRoutes(routerGroup *gin.RouterGroup, handler *controllers.Handler) {
	users := routerGroup.Group("/users")
	users.Use(middleware.AuthMiddleware)
	{
		users.GET("", handler.GetUsers)
		users.GET("/:id", handler.GetUser)
		users.PUT("/:id", handler.UpdateUser)
		users.DELETE("/:id", handler.DeleteUser)
	}
}
