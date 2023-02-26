package router

import (
	"github.com/gin-gonic/gin"

	"github.com/soicchi/chatapp_backend/pkg/controllers"
)

func addAuthRoutes(handler *controllers.Handler, routerGroup *gin.RouterGroup) {
	auth := routerGroup.Group("/auth")
	{
		auth.POST("/signup", handler.SignUpHandler)
		auth.POST("/login", handler.LoginHandler)
	}
}
