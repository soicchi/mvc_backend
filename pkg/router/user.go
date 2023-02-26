package router

import (
	"github.com/gin-gonic/gin"

	"github.com/soicchi/chatapp_backend/pkg/controllers"
	"github.com/soicchi/chatapp_backend/pkg/middleware"
)

func addUserRoutes(routerGroup *gin.RouterGroup) {
	users := routerGroup.Group("/users")
	users.Use(middleware.AuthMiddleware)
	{
		users.GET("", controllers.GetUsers)
	}
}
