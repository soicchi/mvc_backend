package utils

import (
	"fmt"
	"os"
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
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
	assert.NoError(t, err)
	assert.NotNil(t, actualToken)
}

func TestExtractUserIdFromToken(t *testing.T) {
	claims := jwt.MapClaims{
		"user_id": 1.0,
	}
	sendToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	userId, err := ExtractUserIdFromToken(sendToken)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	assert.Equal(t, userId, uint(1))
}

// TODO: Fix this test
// func TestVerifyUserId(t *testing.T) {
// 	claims := jwt.MapClaims{
// 		"user_id": uint(1),
// 	}
// 	sendToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	t.Log(sendToken.Raw)
// 	testHeader := "Bearer " + sendToken.Raw
// 	testUserId := uint(1)
// 	ok, err := VerifyUserId(testUserId, testHeader)
// 	if err != nil {
// 		t.Errorf("Error: %v", err)
// 	}
// 	assert.Equal(t, ok, true)
// }
