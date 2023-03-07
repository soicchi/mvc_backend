package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/soicchi/chatapp_backend/pkg/utils"
)

var authLogFile string = "auth.log"

func AuthMiddleware(ctx *gin.Context) {
	logger, err := utils.SetupLogger(authLogFile)
	if err != nil {
		panic(err)
	}

	tokenString, err := ctx.Cookie("token")
	if err != nil {
		logger.Error(err)
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		ctx.Abort()
		return
	}

	token, err := utils.ParseToken(tokenString)
	if err != nil {
		logger.Error(err)
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid token",
		})
		ctx.Abort()
		return
	}

	userId, err := utils.ExtractUserIdFromToken(token)
	if err != nil {
		logger.Error(err)
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid token",
		})
		ctx.Abort()
		return
	}

	ctx.Set("userId", userId)
	ctx.Next()
}
