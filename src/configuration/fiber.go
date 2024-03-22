package configuration

import "github.com/gofiber/fiber/v2"
import "github.com/goccy/go-json"

func FiberConfig() fiber.Config {
	return fiber.Config{
		AppName:       "Financial Parsing",
		CaseSensitive: true,
		JSONEncoder:   json.Marshal,
		JSONDecoder:   json.Unmarshal,
	}
}
