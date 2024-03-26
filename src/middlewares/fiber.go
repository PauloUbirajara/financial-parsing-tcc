package middlewares

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

var FiberMiddlewares []fiber.Handler = []fiber.Handler{
	// Rate limiter
	limiter.New(limiter.Config{
		Expiration: 2 * time.Second,
		Max:        50,
	}),
	// CORS
	cors.New(),

	// Error Recover
	recover.New(),
}
