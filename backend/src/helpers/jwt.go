package helpers

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
)

func ValidateTokenFunc(signingKey string) jwt.Keyfunc {
	return func(t *jwt.Token) (interface{}, error) {
		log.Debug("signing method name", jwt.SigningMethodHS256.Name)
		log.Debug("found token signing method name", t.Method.Alg())

		// Always check the signing method
		if t.Method.Alg() != jwt.SigningMethodHS256.Name {
			return nil, errors.New(fmt.Sprintf("Unexpected jwt signing method=%v", t.Header["alg"]))
		}

		return []byte(signingKey), nil
	}
}

func ParseJWTToken(signingKey string, tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, ValidateTokenFunc(signingKey))
}
