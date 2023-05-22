package services

import (
	"errors"
	"fmt"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

func CreateToken(userId int) (string, error) {
	permission := jwt.MapClaims{}
	permission["authorized"] = true
	permission["exp"] = time.Now().Add(time.Hour * 3).Unix()
	permission["userid"] = userId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permission)

	if err := godotenv.Load(); err != nil {
		return "", errors.New("error to load .env")
	}

	secretKey := os.Getenv("SECRET_KEY")

	return token.SignedString([]byte(secretKey))
}

func ValidateToken(token string) bool {
	if err := godotenv.Load(); err != nil {
		return false
	}

	secretKey := os.Getenv("SECRET_KEY")

	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("invalid token: %v", token)
		}

		return []byte(secretKey), nil
	})

	return err == nil
}
