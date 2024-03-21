package main

import (
	"log"

	configuration "financial-parsing/src/configuration"
	middlewares "financial-parsing/src/middlewares"
	routes "financial-parsing/src/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config := configuration.FiberConfig()
	middlewares := middlewares.FiberMiddlewares()

	app := fiber.New(config)
	for _, middleware := range middlewares {
		app.Use(middleware)
	}

	rootRouter := app.Group("/api/v1")

	routes.TransactionRoutes(rootRouter)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Go!")
	})

	log.Fatal(app.Listen(":3000"))
}
