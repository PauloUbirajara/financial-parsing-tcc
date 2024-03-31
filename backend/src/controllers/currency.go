package controllers

import (
	"strings"

	models "financial-parsing/src/domain/models"
	usecases "financial-parsing/src/domain/usecases"
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
	log.Info("Currency - Get All")

	var currencies []models.Currency

	result := c.Connection.
		Table("currencies").
		Find(&currencies)

	if result.Error != nil {
		log.Warn(result.Error)
		return ctx.
			Status(fiber.StatusInternalServerError).
			SendString("Error when getting all currencies")
	}

	return ctx.JSON(currencies)
}

func (c CurrencyController) GetById(ctx *fiber.Ctx) error {
	log.Info("Currency - Get By Id")

	id := ctx.Params("id")

	var currency models.Currency

	result := c.Connection.
		Table("currencies").
		First(&currency, "currencies.id = ?", id)

	if result.Error != nil {
		log.Warn(result.Error)
		return ctx.
			Status(fiber.StatusInternalServerError).
			SendString("Error when getting currency by id")
	}

	return ctx.JSON(currency)
}

func (c CurrencyController) Update(ctx *fiber.Ctx) error {
	log.Info("Currency - Update")

	// Getting existing currency to update
	id := ctx.Params("id")
	var existingCurrency models.Currency

	result := c.Connection.
		Table("currencies").
		First(&existingCurrency, "currencies.id = ?", id)

	if existingCurrency.CreatedAt.IsZero() {
		return ctx.
			Status(fiber.StatusNotFound).
			SendString("Could not find currency by id for update")
	}

	if result.Error != nil {
		log.Warn("Error when searching existing currency to update", result.Error)
		return ctx.
			Status(fiber.StatusInternalServerError).
			SendString("Error when searching for existing currency to update")
	}

	currency := new(models.Currency)
	if err := ctx.BodyParser(&currency); err != nil {
		log.Warn("Error when parsing", err)
		return ctx.
			Status(fiber.StatusBadRequest).
			SendString("Could not parse request body to Currency")
	}

	existingCurrency.Name = currency.Name
	existingCurrency.Representation = currency.Representation

	trx := c.Connection.Begin()
	updateResult := trx.Save(&existingCurrency)

	if updateResult.Error != nil {
		log.Warn("Error when updating", updateResult.Error)
		trx.Rollback()
		return ctx.
			Status(fiber.StatusInternalServerError).
			SendString("Error when updating Currency by id")
	}
	trx.Commit()

	return ctx.
		Status(fiber.StatusOK).
		JSON(existingCurrency)
}

func (c CurrencyController) Delete(ctx *fiber.Ctx) error {
	log.Info("Currency - Delete")

	ids := ctx.Query("ids")
	idsToDelete := strings.Split(ids, ",")
	log.Info("IDS to delete: ", idsToDelete)

	var deleted []models.Currency

	// Delete currency
	result := c.Connection.
		Table("currencies").
		Delete(&deleted, "currencies.id = ?", idsToDelete)

	if result.Error != nil {
		log.Warn(result.Error)
		return ctx.
			Status(fiber.StatusInternalServerError).
			SendString("Error when deleting currencies by id")
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func (c CurrencyController) Create(ctx *fiber.Ctx) error {
	log.Info("Currency - Create")

	var body models.Currency

	// Parse request body as currency
	if err := ctx.BodyParser(&body); err != nil {
		log.Warn("Error when parsing body to currency", err)
		return ctx.
			Status(fiber.StatusBadRequest).
			SendString("Could not parse specified request body to Currency")
	}
	body.ID = c.UUIDGenerator.Generate()

	if err := c.ValidateCurrency.Validate(body); err != nil {
		return ctx.
			Status(fiber.StatusBadRequest).
			SendString(err.Error())
	}

	trx := c.Connection.Begin()

	// Create currency
	result := trx.
		Table("currencies").
		Create(&body)

	if result.Error != nil {
		log.Warn(result.Error)
		trx.AddError(result.Error)
		return ctx.
			Status(fiber.StatusInternalServerError).
			SendString("Error when creating currency")
	}

	// TODO Create relationship between currency and user
	// result = trx.
	// 	Table("currency_users").
	// 	Create(&models.Currency_User{
	// 		ID:         c.UUIDGenerator.Generate(),
	// 		CurrencyId: body.ID,
	// 		UserId:     userId
	// 	})

	if result.Error != nil {
		log.Warn(result.Error)
		trx.AddError(result.Error)
		return ctx.
			Status(fiber.StatusInternalServerError).
			SendString("Error when creating relationship between currency and user")
	}

	// Save changes
	trx.Commit()

	return ctx.
		Status(fiber.StatusCreated).
		JSON(body)
}
