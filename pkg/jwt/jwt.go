package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

type TokenString = string

var SecretKey = []byte("secret")

func GenerateToken(username string) (TokenString, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString(SecretKey)

	if err != nil {
		log.Fatalln("Error in generating jwt token")
		return "", err
	}

	return tokenString, nil
}

func ParseToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})

	if err != nil {
		log.Fatalln("failure parsing token")
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username := claims["username"].(string)
		return username, nil
	}

	return "", errors.New("failure retrieving claim")
}
