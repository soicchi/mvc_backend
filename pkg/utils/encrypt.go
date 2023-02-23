package utils

import (
	"crypto/sha256"
	"fmt"
)

func Encrypt(char string) string {
	encryptText := fmt.Sprintf("%x", sha256.Sum256([]byte(char)))
	return encryptText
}
