package middlewares

import (
	"os"
	"time"

	jwtware "github.com/gofiber/contrib/jwt"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

var (
	FiberErrorRecovery = recover.New()
	FiberCORS          = cors.New(cors.Config{
		// Allow everyone
		// AllowOrigins: "*",
	})
	FiberJWT = jwtware.New(
		jwtware.Config{
			SigningKey: jwtware.SigningKey{
				Key: []byte(os.Getenv("JWT_SECRET")),
			},
		},
	)
	FiberRateLimiter = limiter.New(
		limiter.Config{
			Expiration: 2 * time.Second,
			Max:        50,
		},
	)
)
