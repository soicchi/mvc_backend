package router

import (
	"github.com/gin-gonic/gin"

	"github.com/soicchi/chatapp_backend/pkg/controllers"
	"github.com/soicchi/chatapp_backend/pkg/database"
)

func Run() {
	router := setupRouter()
	router.Run()
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	handler := &controllers.Handler{
		DB: database.GetDB(),
	}
	// TODO: SetTrustedProxiesの設定を検討
	api := router.Group("/api")

	v1 := api.Group("/v1")
	addAuthRoutes(handler, v1)
	addUserRoutes(handler, v1)

	return router
}
