package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWTToken(userId uint) (string, error) {
	secretKey := os.Getenv("SECRET_KEY")
	tokenLifeTime, err := strconv.Atoi(os.Getenv("TOKEN_LIFETIME"))
	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(time.Hour * time.Duration(tokenLifeTime)).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwsToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return jwsToken, nil
}

func ExtractToken(authHeader string) (string, error) {
	splitedHeader := strings.Split(authHeader, " ")
	if len(splitedHeader) != 2 {
		return "", fmt.Errorf("Invalid token: %s", authHeader)
	}
	token := splitedHeader[1]
	return token, nil
}

// func VerifyToken(tokenString string) (string, error) {
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*signingMethod); !ok {
// 			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
// 		}
// 	})
// }
