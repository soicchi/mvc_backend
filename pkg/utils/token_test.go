package utils

import (
	"fmt"
	"testing"
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