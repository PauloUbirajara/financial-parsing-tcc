package main

import (
	"fmt"
	"os"

	configuration "financial-parsing/src/configuration"
	helpers "financial-parsing/src/helpers"
	middlewares "financial-parsing/src/middlewares"
	routes "financial-parsing/src/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func main() {
	// DB
	connection, dbError := helpers.CreateConnection(
		configuration.DatabaseHost,
		configuration.DatabaseUsername,
		configuration.DatabasePassword,
		configuration.DatabasePort,
	)
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
	host := os.Getenv("APP_HOST")
	port := os.Getenv("APP_PORT")
	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "3000"
	}
	serverAddress := fmt.Sprintf("%s:%s", host, port)
	log.Info("Listening to: ", serverAddress)
	log.Fatal(app.Listen(serverAddress))
}
