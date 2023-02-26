package router

import (
	"github.com/gin-gonic/gin"
)

func Run() {
	router := setupRouter()
	router.Run()
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	// TODO: SetTrustedProxiesの設定を検討
	api := router.Group("/api")

	v1 := api.Group("/v1")
	addAuthRoutes(v1)
	addUserRoutes(v1)

	return router
}
