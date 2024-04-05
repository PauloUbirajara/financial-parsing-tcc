package controllers

import (
	"strings"

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
		Find(&currencies, "currency_users.user_id = ?", user.ID)

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
		id string = ctx.Params("id")

		currency models.Currency
		user     models.User
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
		First(&currency, "currency_users.user_id = ? AND currency_users.currency_id = ?", user.ID, id)

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
		id string = ctx.Params("id")

		currency models.Currency
		body     models.Currency
		user     models.User
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
	if err := ctx.BodyParser(&body); err != nil {
		log.Warn("Error when parsing currency body for update")
		log.Warn(err)
		return ctx.
			Status(fiber.StatusBadRequest).
			JSON(fiber.Map{
				"error": "Error when parsing body for update",
			})
	}

	result = c.Connection.
		Joins("JOIN currency_users ON currency_users.currency_id = currencies.id").
		Joins("JOIN users ON users.id = currency_users.user_id").
		First(&currency, "currency_users.user_id = ? AND currency_users.currency_id = ?", user.ID, id)

	if result.Error != nil {
		return ctx.
			Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{
				"error": "Could not update currency by id",
			})
	}

	trx := c.Connection.Begin()
	currency.Name = body.Name
	currency.Representation = body.Representation

	result = trx.Save(&currency)
	if result.Error != nil {
		trx.Rollback()
		return ctx.
			Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{
				"error": "Could not update currency using body",
			})
	}
	trx.Commit()

	return ctx.
		Status(fiber.StatusOK).
		JSON(currency)
}

func (c CurrencyController) Delete(ctx *fiber.Ctx) error {
	log.Debug("Currency - Delete")

	var (
		ids string = ctx.Query("ids")

		currenciesUser []models.Currency_User
		user           models.User
	)

	idsToDelete := strings.Split(ids, ",")

	result := c.Connection.First(&user, "username = ?", helpers.GetUsername(ctx))

	if result.Error != nil {
		return ctx.
			Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{
				"error": "Could not get user while deleting currencies by ids",
			})
	}

	trx := c.Connection.Begin()

	// Find currency relationships
	result = trx.
		Find(&currenciesUser, "currency_users.user_id = ? AND currency_users.currency_id IN ?", user.ID, idsToDelete)

	if result.Error != nil {
		trx.Rollback()
		return ctx.
			Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{
				"error": "Could not find currency relationships to delete",
			})
	}

	currenciesToDelete := make([]string, len(currenciesUser))
	for _, currencyUser := range currenciesUser {
		currenciesToDelete = append(currenciesToDelete, currencyUser.CurrencyId)
	}

	// Delete currencies
	result = trx.
		Where("currencies.id IN ?", currenciesToDelete).
		Delete(&models.Currency{})

	if result.Error != nil {
		trx.Rollback()
		return ctx.
			Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{
				"error": "Could not delete currencies by ids",
			})
	}

	// Delete currency relationships
	result = trx.
		Where("currency_users.currency_id IN ?", currenciesToDelete).
		Delete(&models.Currency_User{})

	if result.Error != nil {
		trx.Rollback()
		return ctx.
			Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{
				"error": "Could not delete currency relationships",
			})
	}

	trx.Commit()

	return ctx.
		Status(fiber.StatusOK).
		JSON(fiber.Map{
			"message": "Currencies deleted successfully",
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

	// Create currency relationship
	currencyUser = models.Currency_User{
		ID:         c.UUIDGenerator.Generate(),
		UserId:     user.ID,
		CurrencyId: currency.ID,
	}

	result = trx.
		Create(&currencyUser)

	if result.Error != nil {
		log.Warn(result.Error)
		trx.Rollback()
		return ctx.
			Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{
				"error": "Could not create currency relationship",
			})
	}

	trx.Commit()

	return ctx.
		Status(fiber.StatusOK).
		JSON(currency)
}
