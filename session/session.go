package session

import (
	"time"

	"github.com/golang-jwt/jwt"
)

var hmacSampleSecret []byte = []byte("abcd")

func New(idUser string, accessLevel int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id_user":      idUser,
		"access_level": accessLevel,
		"nbf":          time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	return token.SignedString(hmacSampleSecret)
}
