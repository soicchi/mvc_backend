package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateJWTToken(t *testing.T) {
	email := "test@test.com"
	password := "Password1234"
	jwsToken, err := GenerateJWTToken(email, password)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if jwsToken == "" {
		t.Errorf("Error: %v", err)
	}

	t.Log(fmt.Sprintf("Token: %v", jwsToken))
}

func TestExtractToken(t *testing.T) {
	testHeaders := []string{
		"Bearer samplesamplesample",
		"samplesample",
		"Bearer",
		"",
	}
	for _, header := range testHeaders {
		assert := assert.New(t)
		token, err := ExtractToken(header)
		if err != nil {
			assert.Equal(err, fmt.Errorf("Invalid token: %s", header))
			assert.Equal(token, "")
		}

		t.Log(fmt.Sprintf("Token: %v", token))
	}
}