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

type CurrencyRecordController struct {
	Connection             *gorm.DB
	UUIDGenerator          protocols.UUIDGenerator
	ValidateCurrencyRecord usecases.ValidateModel[models.CurrencyRecord]
}

func (c CurrencyRecordController) GetAll(ctx *fiber.Ctx) error {
	log.Debug("CurrencyRecord - GetAll")

	var (
		currencyId string = ctx.Params("currencyId")

		currencyRecords []models.CurrencyRecord
		user            models.User
	)

	result := c.Connection.First(&user, "username = ?", helpers.GetUsername(ctx))

	if result.Error != nil {
		return ctx.
			Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{
				"error": "Could not get user while getting all currency records",
			})
	}

	result = c.Connection.
		Joins("JOIN currency_currency_records ON currency_currency_records.currency_record_id = currency_records.id").
		Joins("JOIN currency_users ON currency_users.currency_id = currency_currency_records.currency_id").
		Joins("JOIN users ON users.id = currency_users.user_id").
		Find(&currencyRecords, "currency_users.user_id = ? AND currency_currency_records.currency_id = ?", user.ID, currencyId)

	if result.Error != nil {
		return ctx.
			Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{
				"error": "Could not get all currency records",
			})
	}

	return ctx.
		Status(fiber.StatusOK).
		JSON(currencyRecords)
}

func (c CurrencyRecordController) GetById(ctx *fiber.Ctx) error {
	log.Debug("CurrencyRecord - GetById")
	return ctx.
		Status(fiber.StatusInternalServerError).
		JSON(fiber.Map{
			"error": "Could not get currency record by id",
		})
}

func (c CurrencyRecordController) Update(ctx *fiber.Ctx) error {
	log.Debug("CurrencyRecord - Update")
	return ctx.
		Status(fiber.StatusInternalServerError).
		JSON(fiber.Map{
			"error": "Could not update currency record by id",
		})
}

func (c CurrencyRecordController) Delete(ctx *fiber.Ctx) error {
	log.Debug("CurrencyRecord - Delete")
	return ctx.
		Status(fiber.StatusInternalServerError).
		JSON(fiber.Map{
			"error": "Could not delete currency records by ids",
		})
}

func (c CurrencyRecordController) Create(ctx *fiber.Ctx) error {
	log.Debug("CurrencyRecord - Create")
	return ctx.
		Status(fiber.StatusInternalServerError).
		JSON(fiber.Map{
			"error": "Could not create currency record",
		})
}
