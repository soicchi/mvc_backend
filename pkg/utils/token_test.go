package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateJWTToken(t *testing.T) {
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
