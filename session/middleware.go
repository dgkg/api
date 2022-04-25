package session

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func NewMiddleware(accessLvlMin int) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		valAuth := ctx.GetHeader("Authorization")
		if len(valAuth) == 0 || !strings.Contains(valAuth, "Bearer ") {
			ctx.AbortWithStatus(http.StatusForbidden)
			return
		}

		token, err := jwt.Parse(valAuth[7:], func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("session: unexpected signing method: %w", token.Header["alg"].(error))
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return hmacSampleSecret, nil
		})
		if err != nil {
			ctx.AbortWithStatus(http.StatusForbidden)
			return
		}

		val, ok := token.Claims.(CustomClaimsUser)
		if !ok {
			ctx.AbortWithStatus(http.StatusForbidden)
			return
		}

		if val.AccessLevel < accessLvlMin {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		// set the claims into the context.
		ctx.Set("claims", val)
	}
}
