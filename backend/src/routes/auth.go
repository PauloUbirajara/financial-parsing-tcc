package routes

import (
	"financial-parsing/src/domain/models"
	"financial-parsing/src/helpers"
	"os"
	"time"

	// controllers "financial-parsing/src/controllers"
	// models "financial-parsing/src/domain/models"

	// validatemodel "financial-parsing/src/data/usecases/validateModel"
	// databaseadapter "financial-parsing/src/utils/databaseAdapter"
	// uuidgenerator "financial-parsing/src/utils/uuidGenerator"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

func Login(c *fiber.Ctx, connection *gorm.DB) error {
	log.Info("Login")

	email := c.FormValue("email")
	password := c.FormValue("password")

	// Validate within database
	var user models.User
	result := connection.
		Table("users").
		First(&user, "users.email = ? AND users.password_hash = ?", email, helpers.StringToSHA512Sum(password))

	if result.Error != nil {
		log.Warn("Error when logging user in", result.Error)
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// Create the Claims
	// TODO Use later with 15 min
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Minute * 2).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}

func Register(c *fiber.Ctx, connection *gorm.DB) error {
	log.Info("Register")

	email := c.FormValue("email")
	password := c.FormValue("password")
	confirmPassword := c.FormValue("confirmPassword")

	// Validate before adding to database
	if password != confirmPassword {
		return c.
			Status(fiber.StatusBadRequest).
			SendString("Senhas não são iguais")
	}

	// Add to database
	trx := connection.Begin()

	user := models.User{
		Email:        email,
		PasswordHash: helpers.StringToSHA512Sum(password),
	}

	result := trx.
		Table("users").
		Create(&user)

	if result.Error != nil {
		log.Warn(result.Error)
		trx.AddError(result.Error)
	}
	commitResult := trx.Commit()
	if commitResult.Error != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusCreated)
}

func AuthRoutes(router fiber.Router, connection *gorm.DB) fiber.Router {
	authRouter := router.Group("/auth")
	authRouter.Post("/login", func(c *fiber.Ctx) error { return Login(c, connection) })
	authRouter.Post("/register", func(c *fiber.Ctx) error { return Register(c, connection) })
	return authRouter
}
