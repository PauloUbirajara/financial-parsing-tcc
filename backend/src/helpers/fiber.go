package helpers

import "github.com/gofiber/fiber/v2"

// Gets username stored from JWT token and stored in backend request header
func GetUsername(ctx *fiber.Ctx) string {
  username := (&ctx.Request().Header).Peek("username")
  return string(username)
}
