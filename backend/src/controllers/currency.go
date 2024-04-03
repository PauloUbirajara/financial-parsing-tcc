package controllers

import (
	models "financial-parsing/src/domain/models"
	usecases "financial-parsing/src/domain/usecases"
	helpers "financial-parsing/src/helpers"
	protocols "financial-parsing/src/protocols"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
)

type CurrencyController struct {
	Connection       *gorm.DB
	UUIDGenerator    protocols.UUIDGenerator
	ValidateCurrency usecases.ValidateModel[models.Currency]
}

func (c CurrencyController) GetAll(ctx *fiber.Ctx) error {
	log.Debug("Currency - GetAll")

	var (
		currencies []models.Currency
		user       models.User
	)

	result := c.Connection.First(&user, "username = ?", helpers.GetUsername(ctx))

	if result.Error != nil {
		return ctx.
			Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{
				"error": "Could not get user while getting all currencies",
			})
	}

	result = c.Connection.
		Table("currencies").
		Joins("JOIN currency_users ON currency_users.user_id = ?", user.ID).
		Find(&currencies)

	if result.Error != nil {
		return ctx.
			Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{
				"error": "Could not get all currencies",
			})
	}

	return ctx.
		Status(fiber.StatusOK).
		JSON(currencies)
}

func (c CurrencyController) GetById(ctx *fiber.Ctx) error {
	log.Debug("Currency - GetById")

	var (
		currency models.Currency
		user     models.User
		id       string = ctx.Params("id")
	)

	result := c.Connection.First(&user, "username = ?", helpers.GetUsername(ctx))

	if result.Error != nil {
		return ctx.
			Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{
				"error": "Could not get user while getting currency by id",
			})
	}

	result = c.Connection.
		Table("currencies").
		Joins("JOIN currency_users ON currency_users.user_id = ?", user.ID).
		First(&currency, "currencies.id = ?", id)

	if result.Error != nil {
		return ctx.
			Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{
				"error": "Could not get currency by id",
			})
	}

	return ctx.
		Status(fiber.StatusOK).
		JSON(currency)
}

func (c CurrencyController) Update(ctx *fiber.Ctx) error {
	log.Debug("Currency - UpdateById")

	return ctx.
		Status(fiber.StatusInternalServerError).
		JSON(fiber.Map{
			"error": "Could not update currency by id",
		})
}

func (c CurrencyController) Delete(ctx *fiber.Ctx) error {
	log.Debug("Currency - Delete")

	return ctx.
		Status(fiber.StatusInternalServerError).
		JSON(fiber.Map{
			"error": "Could not delete currencies by ids",
		})
}

func (c CurrencyController) Create(ctx *fiber.Ctx) error {
	log.Debug("Currency - Create")

	return ctx.
		Status(fiber.StatusInternalServerError).
		JSON(fiber.Map{
			"error": "Could not create currency",
		})
}
