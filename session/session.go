package session

import (
	"time"

	"github.com/golang-jwt/jwt"
)

var hmacSampleSecret []byte = []byte("abcd")

type CustomClaimsUser struct {
	IDUser      string `json:"id_user"`
	AccessLevel int    `json:"access_level"`
	jwt.StandardClaims
}

func New(idUser string, accessLevel int) (string, error) {
	jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaimsUser{
		idUser,
		accessLevel,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 10).Unix(),
		},
	})
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id_user":      idUser,
		"access_level": accessLevel,
		"nbf":          time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	return token.SignedString(hmacSampleSecret)
}
