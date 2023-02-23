package router

import (
	"github.com/gin-gonic/gin"

	"github.com/soicchi/chatapp_backend/pkg/controllers"
)

func addAuthRoutes(routerGroup *gin.RouterGroup) {
	auth := routerGroup.Group("/auth")
	{
		auth.POST("/signup", controllers.SignUpHandler)
		auth.POST("/login", controllers.LoginHandler)
	}
}
