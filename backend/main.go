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
	app := fiber.New(configuration.FiberConfig)
	app.Use(middlewares.FiberCORS)
	app.Use(middlewares.FiberErrorRecovery)
	app.Use(middlewares.FiberRateLimiter)

	// Routes Setup
	rootRouter := app.Group("/api/v1")
	routes.AuthRoutes(rootRouter, connection)

	// JWT will be used on the next routes
	app.Use(middlewares.FiberJWT)
	routes.CurrencyRoutes(rootRouter, connection)
	routes.CurrencyRecordsRoutes(rootRouter, connection)

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
