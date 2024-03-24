package main

import (
	"fmt"
	"log"
	"os"

	configuration "financial-parsing/src/configuration"
	helpers "financial-parsing/src/helpers"
	middlewares "financial-parsing/src/middlewares"
	routes "financial-parsing/src/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// DB
	connection, dbError := helpers.CreateConnection()
	if dbError != nil {
		panic("Could not establish connection to the database")
	}

	// Configuration
	config := configuration.FiberConfig()
	middlewares := middlewares.FiberMiddlewares()

	app := fiber.New(config)
	for _, middleware := range middlewares {
		app.Use(middleware)
	}

	// Routes Setup
	rootRouter := app.Group("/api/v1")

	routes.CurrencyRoutes(rootRouter, connection)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Go")
	})

	// Server start
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))
}
