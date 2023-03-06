package router

import (
	"github.com/gin-gonic/gin"

	"github.com/soicchi/chatapp_backend/pkg/controllers"
	"github.com/soicchi/chatapp_backend/pkg/middleware"
)

func addPostRoutes(routerGroup *gin.RouterGroup, handler *controllers.Handler) {
	posts := routerGroup.Group("/posts")
	posts.Use(middleware.AuthMiddleware)
	{
		posts.POST("/create", handler.CreatePost)
	}
}