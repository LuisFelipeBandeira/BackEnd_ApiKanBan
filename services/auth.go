package services

import (
	"errors"
	"fmt"
	"os"
	"strconv"
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

func ValidateToken(t string) error {
	token, err := jwt.Parse(t, ReturnSecretKey)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("Inválid token!")
}

func ReturnSecretKey(t *jwt.Token) (interface{}, error) {

	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	secretKey := os.Getenv("SECRET_KEY")

	if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Método de assinatura inexperado: %v", t.Header["alg"])
	}

	return secretKey, nil
}

func GetUserIdByToken(t string) (interface{}, error) {
	token, err := jwt.Parse(t, ReturnSecretKey)
	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId := strconv.ParseInt(claims["userid"], 10, 64)
		fmt.Println(userId)
	}

}
