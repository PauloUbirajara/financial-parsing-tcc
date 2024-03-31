package controllers

import (
	"errors"
	"financial-parsing/src/domain/models"
	"financial-parsing/src/helpers"
	"financial-parsing/src/protocols"
	"regexp"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type AuthController struct {
	Connection    *gorm.DB
	JwtSecret     string
	UUIDGenerator protocols.UUIDGenerator
}

type RegisterBody struct {
	Email           string
	Password        string
	ConfirmPassword string
}

type LoginBody struct {
	Email    string
	Password string
}

func validateRegisterBody(registerBody RegisterBody) error {
	emailRegex, _ := regexp.Compile(`^[\w\-\.]+@([\w-]+\.)+[\w-]{2,}$`)
	if !emailRegex.MatchString(registerBody.Email) {
		return errors.New("E-mail inválido")
	}

	if registerBody.Password != registerBody.ConfirmPassword {
		return errors.New("Senhas não são iguais")
	}

	passwordRegex, _ := regexp.Compile(`^.{8,}$`)
	if !passwordRegex.MatchString(registerBody.Password) {
		return errors.New("Senha precisa conter ao menos 8 caracteres")
	}

	return nil
}

func (a AuthController) Register(ctx *fiber.Ctx) error {
	log.Info("Register")

	var registerBody RegisterBody

	if err := ctx.BodyParser(&registerBody); err != nil {
		log.Warn("Error when parsing body", err)
		return ctx.
			Status(fiber.StatusBadRequest).
			SendString("Erro ao obter dados para cadastro de usuário")
	}

	log.Debug(registerBody)

	// Validate before adding to database
	if err := validateRegisterBody(registerBody); err != nil {
		return ctx.
			Status(fiber.StatusBadRequest).
			SendString(err.Error())
	}

	// Check if already exists user with email
	var existingUser models.User
	result := a.Connection.
		Table("users").
		First(&existingUser, "email = ?", registerBody.Email)

	if !existingUser.CreatedAt.IsZero() {
		log.Warn(result.Error)
		return ctx.
			Status(fiber.StatusConflict).
			SendString("Usuário já existe com esse email")
	}

	// Add to database
	trx := a.Connection.Begin()

	user := models.User{
		ID:           a.UUIDGenerator.Generate(),
		Email:        registerBody.Email,
		PasswordHash: helpers.StringToSHA512Sum(registerBody.Password),
		Active:       false,
	}

	trx.Create(&user)
	commitResult := trx.Commit()
	if commitResult.Error != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.SendStatus(fiber.StatusCreated)

}

func (a AuthController) Login(ctx *fiber.Ctx) error {
	log.Info("Login")

	var loginBody LoginBody

	if err := ctx.BodyParser(&loginBody); err != nil {
		log.Warn("Error when parsing body", err)
		return ctx.
			Status(fiber.StatusBadRequest).
			SendString("Erro ao obter dados para login de usuário")
	}

	// Validate within database
	var user models.User
	result := a.Connection.
		Table("users").
		First(
			&user,
			"users.email = ? AND users.password_hash = ?",
			loginBody.Email,
			helpers.StringToSHA512Sum(loginBody.Password),
		)

	if result.Error != nil {
		log.Warn("Error when logging user in", result.Error)
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}

	// Create the Claims
	// TODO Use later with 15 min
	claims := jwt.MapClaims{
		"email": loginBody.Email,
		"exp":   time.Now().Add(time.Minute * 2).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(a.JwtSecret))
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.JSON(fiber.Map{"token": t})
}
