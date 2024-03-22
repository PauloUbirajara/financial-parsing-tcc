package middlewares

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func FiberMiddlewares() []fiber.Handler {
	rateLimiterHandler := limiter.New(limiter.Config{
		Expiration: 2 * time.Second,
		Max:        50,
	})
	corsHandler := cors.New()
	recoverHandler := recover.New()

	return []fiber.Handler{
		rateLimiterHandler,
		corsHandler,
		recoverHandler,
	}
}
