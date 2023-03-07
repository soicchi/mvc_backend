package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAuthMiddleware(t *testing.T) {
	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)
	request, _ := http.NewRequest("GET", "/", nil)
	context.Request = request
	AuthMiddleware(context)

	assert := assert.New(t)
	assert.Equal(context.Writer.Status(), http.StatusUnauthorized)
}
