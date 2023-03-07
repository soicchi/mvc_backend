package middleware

import (
	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/soicchi/chatapp_backend/pkg/utils"
)

func AuthMiddleware(ctx *gin.Context) {
	tokenString, err := ctx.Cookie("token")
	fmt.Println(ctx.Request)
	fmt.Println(tokenString)
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
