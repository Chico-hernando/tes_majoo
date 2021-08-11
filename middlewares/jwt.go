package middlewares

import (
	"majoo/constants"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT(userId int) (string, error) {
	claims := jwt.MapClaims{}

	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(constants.JWT_SECRET))
}