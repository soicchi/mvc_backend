package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/soicchi/chatapp_backend/pkg/utils"
)

func AuthMiddleware(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")
	tokenString, err := utils.ExtractToken(authHeader)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"error":   err.Error(),
			"message": "Invalid token",
		})
		context.Abort()
		return
	}

	_, err = utils.ParseToken(tokenString)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"error":   err.Error(),
			"message": "Invalid token",
		})
		context.Abort()
		return
	}

	context.Next()
}
