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
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["exp"] = time.Now().Add(time.Hour * 3).Unix()
	claims["userid"] = userId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

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

func GetUserIdByToken(t string) (int, error) {
	token, err := jwt.Parse(t, ReturnSecretKey)
	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId, err := strconv.ParseUint(fmt.Sprintf("%.d", claims["user_id"]), 10, 64)
		if err != nil {
			return 0, err
		}

		return int(userId), nil
	}

	return 0, errors.New("Error to get token userId")
}
