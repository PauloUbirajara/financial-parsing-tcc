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
		Joins("JOIN currency_users ON currency_users.currency_id = currencies.id").
		Joins("JOIN users ON users.id = currency_users.user_id").
		Find(&currencies, "users.id = ?", user.ID)

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
		Joins("JOIN currency_users ON currency_users.currency_id = currencies.id").
		Joins("JOIN users ON users.id = currency_users.user_id").
		First(&currency, "users.id = ? AND currencies.id = ?", user.ID, id)

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
				"error": "Could not get user while updating currency by id",
			})
	}

	// Parse updated body
	updatedCurrency := new(models.Currency)
	if err := ctx.BodyParser(&updatedCurrency); err != nil {
		log.Warn("Error when parsing currency body for update")
		log.Warn(err)
		return ctx.
			Status(fiber.StatusBadRequest).
			JSON(fiber.Map{
				"error": "Error when parsing body for update",
			})
	}

	trx := c.Connection.Begin()

	result = trx.
		Joins("JOIN currency_users ON currency_users.currency_id = currencies.id").
		Joins("JOIN users ON users.id = currency_users.user_id").
		First(&currency, "users.id = ? AND currencies.id = ?", user.ID, id)

	if result.Error != nil {
		trx.Rollback()
		return ctx.
			Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{
				"error": "Could not update currency by id",
			})
	}

	currency.Name = updatedCurrency.Name
	currency.Representation = updatedCurrency.Representation

	trx.Save(&currency)

	trx.Commit()

	return ctx.
		Status(fiber.StatusOK).
		JSON(currency)
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

	var (
		currency     models.Currency
		user         models.User
		currencyUser models.Currency_User
	)

	result := c.Connection.First(&user, "username = ?", helpers.GetUsername(ctx))

	if result.Error != nil {
		return ctx.
			Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{
				"error": "Could not get user while creating currency",
			})
	}

	// Parse body
	if err := ctx.BodyParser(&currency); err != nil {
		log.Warn("Error when parsing currency body for create")
		log.Warn(err)
		return ctx.
			Status(fiber.StatusBadRequest).
			JSON(fiber.Map{
				"error": "Error when parsing body for create",
			})
	}
	currency.ID = c.UUIDGenerator.Generate()

	trx := c.Connection.Begin()

	// Create currency
	result = trx.
		Create(&currency)

	if result.Error != nil {
		log.Warn(result.Error)
		trx.Rollback()
		return ctx.
			Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{
				"error": "Could not create currency",
			})
	}

	currencyUser = models.Currency_User{
		ID:         c.UUIDGenerator.Generate(),
		UserId:     user.ID,
		CurrencyId: currency.ID,
	}

	// Create relationship between user and currency
	result = trx.
		Create(&currencyUser)

	if result.Error != nil {
		log.Warn(result.Error)
		trx.Rollback()
		return ctx.
			Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{
				"error": "Could not create relationship between user and currency",
			})
	}

	trx.Commit()

	return ctx.
		Status(fiber.StatusOK).
		JSON(currency)
}
