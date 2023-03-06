package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userId uint) (string, error) {
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
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ExtractToken(authHeader string) (string, error) {
	if authHeader == "" {
		return "", fmt.Errorf("Invalid token: %s", authHeader)
	}

	splitedHeader := strings.Split(authHeader, " ")
	if len(splitedHeader) != 2 {
		return "", fmt.Errorf("Invalid token: %s", authHeader)
	}
	token := splitedHeader[1]
	return token, nil
}

func ParseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func ExtractUserIdFromToken(token *jwt.Token) (uint, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, fmt.Errorf("cannot convert claims to MapClaims")
	}

	userId, ok := claims["user_id"]
	if !ok {
		return 0, fmt.Errorf("Does not have user_id in token claims")
	}
	userIdUint := uint(userId.(float64))

	return userIdUint, nil
}
