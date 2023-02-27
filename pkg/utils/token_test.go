package utils

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/golang-jwt/jwt/v5"
)

func TestGenerateToken(t *testing.T) {
	var userId uint = 1
	token, err := GenerateToken(userId)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if token == "" {
		t.Errorf("Error: %v", err)
	}

	t.Log(fmt.Sprintf("Token: %v", token))
}

func TestExtractToken(t *testing.T) {
	testHeaders := []string{
		"Bearer samplesamplesample",
		"samplesample",
		"Bearer",
		"",
	}
	assert := assert.New(t)
	for _, header := range testHeaders {
		token, err := ExtractToken(header)
		if err != nil {
			assert.Equal(err, fmt.Errorf("Invalid token: %s", header))
			assert.Equal(token, "")
		}

		t.Log(fmt.Sprintf("Token: %v", token))
	}
}

func TestParseTokenInvalidSigningMethod(t *testing.T) {
    _, err := ParseToken("invalid-token")
    assert.Error(t, err)
}

func TestParseTokenValid(t *testing.T) {
    sendToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{})
    tokenString, err := sendToken.SignedString([]byte(os.Getenv("SECRET_KEY")))
    assert.NoError(t, err)

    actualToken, err := ParseToken(tokenString)
	t.Log(actualToken)
    assert.NoError(t, err)
    assert.NotNil(t, actualToken)
}