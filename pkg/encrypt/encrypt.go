package encrypt

import (
	"fmt"
	"crypto/sha256"
)

func Encrypt(char string) string {
	encryptText := fmt.Sprintf("%x", sha256.Sum256([]byte(char)))
	return encryptText
}
