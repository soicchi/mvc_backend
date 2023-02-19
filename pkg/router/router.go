package router

import (
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func Run() {
	setupRouter()
	router.Run()
}

func setupRouter() {
	// TODO: SetTrustedProxiesの設定を検討
	api := router.Group("/api")

	v1 := api.Group("/v1")
	addUserRoutes(v1)
}