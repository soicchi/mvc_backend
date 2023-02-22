package router

import (
	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.Default()
	setupRouter(router)
	router.Run()
}

func setupRouter(router *gin.Engine) {
	// TODO: SetTrustedProxiesの設定を検討
	api := router.Group("/api")

	v1 := api.Group("/v1")
	addAuthRoutes(v1)
	// addUserRoutes(v1)
}
