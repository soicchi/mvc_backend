package router

import (
	"github.com/gin-gonic/gin"

	"github.com/soicchi/chatapp_backend/pkg/middleware"
	"github.com/soicchi/chatapp_backend/pkg/controllers"
)

func addUserRoutes(handler *controllers.Handler, routerGroup *gin.RouterGroup) {
	users := routerGroup.Group("/users")
	users.Use(middleware.AuthMiddleware)
	{
		users.GET("", handler.GetUsers)
	}
}
