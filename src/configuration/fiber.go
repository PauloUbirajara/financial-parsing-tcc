package configuration

import "github.com/gofiber/fiber/v2"
import "github.com/goccy/go-json"

var FiberConfig fiber.Config = fiber.Config{
	AppName:       "Financial Parsing",
	CaseSensitive: true,
	JSONEncoder:   json.Marshal,
	JSONDecoder:   json.Unmarshal,
}
