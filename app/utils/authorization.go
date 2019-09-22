package utils

import (
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/teliaz/goapi/config"
)

func GenerateToken() (string, error) {

	config := config.GetConfig()

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(config.AUTH.ExpirationMinutes)).Unix()
	tokenString, err := token.SignedString(config.AUTH.HmacSecret)
	if err != nil {
		log.Fatal("Error generating JWT", err.Error())
		return "", err
	}
	return tokenString, nil
}
