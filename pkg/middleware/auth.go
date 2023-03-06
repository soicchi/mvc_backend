package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/soicchi/chatapp_backend/pkg/utils"
)

func AuthMiddleware(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	tokenString, err := utils.ExtractToken(authHeader)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   err.Error(),
			"message": "Invalid token",
		})
		ctx.Abort()
		return
	}

	token, err := utils.ParseToken(tokenString)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   err.Error(),
			"message": "Invalid token",
		})
		ctx.Abort()
		return
	}

	userId, err := utils.ExtractUserIdFromToken(token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   err.Error(),
			"message": "Invalid token",
		})
		ctx.Abort()
		return
	}

	ctx.Set("userId", userId)
	ctx.Next()
}
