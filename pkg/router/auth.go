package router

import (
	"github.com/gin-gonic/gin"

	"github.com/soicchi/chatapp_backend/pkg/controllers"
)

func addAuthRoutes(routerGroup *gin.RouterGroup, handler *controllers.Handler) {
	auth := routerGroup.Group("/auth")
	{
		auth.POST("/signup", handler.SignUpHandler)
		auth.POST("/login", handler.LoginHandler)
	}
}
