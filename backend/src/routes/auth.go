package routes

import (
	"os"

	controllers "financial-parsing/src/controllers"

	uuidgenerator "financial-parsing/src/utils/uuidGenerator"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AuthRoutes(router fiber.Router, connection *gorm.DB) fiber.Router {
	authRouter := router.Group("/auth")
	authController := controllers.AuthController{
		JwtSecret:     os.Getenv("JWT_SECRET"),
		Connection:    connection,
		UUIDGenerator: uuidgenerator.GoogleUUIDGenerator{},
	}
	authRouter.Post("/login", authController.Login)
	authRouter.Post("/register", authController.Register)
	return authRouter
}