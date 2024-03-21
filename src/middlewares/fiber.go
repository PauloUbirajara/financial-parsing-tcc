package middlewares

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func FiberMiddlewares() [](func(*fiber.Ctx) error) {
	rateLimiterHandler := limiter.New(limiter.Config{
		Expiration: 10 * time.Second,
		Max:        3,
	})
	corsHandler := cors.New()

	return [](func(*fiber.Ctx) error){
		rateLimiterHandler,
		corsHandler,
	}
}
