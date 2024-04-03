package middlewares

import (
	"errors"
	"strings"

	helpers "financial-parsing/src/helpers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
)

var JWTInfoMiddleware func(signingKey string) fiber.Handler = func(signingKey string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get header
		header := &c.Request().Header
		if header == nil {
			return errors.New("Could not get API header object")
		}

		// Retrieve token string
		authorizationString, hasToken := c.GetReqHeaders()[fiber.HeaderAuthorization]

		if !hasToken {
			return errors.New("Could not get Bearer token")
		}

		log.Debug(authorizationString)
		tokenString := strings.Split(authorizationString[0], " ")

		// Parse string for JWT token object
		token, err := helpers.ParseJWTToken(signingKey, tokenString[1])

		if err != nil {
			return err
		}

		// Get claims from token
		claims := token.Claims.(jwt.MapClaims)
		username, hasUsername := claims["username"]

		if !hasUsername {
			return errors.New("Could not get username from authenticated user")
		}

		header.Add("username", username.(string))
		log.Debug("username was added in request header")
		log.Debug(username)

		return c.Next()
	}
}
