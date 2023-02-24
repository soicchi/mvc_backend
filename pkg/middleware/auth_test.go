package middleware

import(
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAuthMiddleware(t *testing.T) {
	ginContext, _ := gin.CreateTestContext(httptest.NewRecorder())
	request, _ := http.NewRequest("GET", "/", nil)
	request.Header.Add("Authorization", "Bearer")
	ginContext.Request = request
	AuthMiddleware(ginContext)

	assert := assert.New(t)
	assert.Equal(ginContext.Writer.Status(), http.StatusUnauthorized)
}