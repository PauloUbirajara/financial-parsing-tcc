package configuration

import "github.com/gofiber/fiber/v2"

func FiberConfig() fiber.Config {
	return fiber.Config{
		AppName:       "Financial Parsing",
		CaseSensitive: true,
	}
}
