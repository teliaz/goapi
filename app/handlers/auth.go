package handlers

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/teliaz/goapi/config"
)


func GenerateToken() (string, error){
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now.Add(time.Minute.Days * 7).Unix

	tokenString, err := token.SignedString(&authCongif.settings.)

	if err != {
		log.Fatal("Error generating JWT", err.Error())
		return "", err
	}
	re
}