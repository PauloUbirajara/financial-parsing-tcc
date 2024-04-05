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
		Joins("JOIN currency_record_users ON currency_record_users.currency_record_id = currency_records.id").
		Joins("JOIN users ON users.id = currency_record_users.user_id").
		Find(&currencyRecords, "currency_record_users.user_id = ? AND currency_record_users.currency_id = ?", user.ID, currencyId)

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

	var (
		currencyId string = ctx.Params("currencyId")
		id         string = ctx.Params("id")

		currencyRecord models.CurrencyRecord
		user           models.User
	)

	result := c.Connection.First(&user, "username = ?", helpers.GetUsername(ctx))

	if result.Error != nil {
		return ctx.
			Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{
				"error": "Could not get user while getting currency record by id",
			})
	}

	result = c.Connection.
		Joins("JOIN currency_record_users ON currency_record_users.currency_record_id = currency_records.id").
		Joins("JOIN users ON users.id = currency_record_users.user_id").
		First(&currencyRecord, "currency_record_users.user_id = ? AND currency_record_users.currency_id = ? AND currency_record_users.currency_record_id = ?", user.ID, currencyId, id)

	if result.Error != nil {
		return ctx.
			Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{
				"error": "Could not get currency record by id",
			})
	}

	return ctx.
		Status(fiber.StatusOK).
		JSON(currencyRecord)
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

	var (
		currencyId string = ctx.Params("currencyId")

		currencyRecord         models.CurrencyRecord
		currencyCurrencyRecord models.CurrencyRecord_User
		user                   models.User
	)

	result := c.Connection.First(&user, "username = ?", helpers.GetUsername(ctx))

	if result.Error != nil {
		return ctx.
			Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{
				"error": "Could not get user while creating currency record",
			})
	}

	// Parse body
	if err := ctx.BodyParser(&currencyRecord); err != nil {
		log.Warn("Error when parsing currency record body for create")
		log.Warn(err)
		return ctx.
			Status(fiber.StatusBadRequest).
			JSON(fiber.Map{
				"error": "Error when parsing body for create",
			})
	}
	currencyRecord.ID = c.UUIDGenerator.Generate()

	trx := c.Connection.Begin()

	// Create currency record
	result = trx.
		Create(&currencyRecord)

	if result.Error != nil {
		log.Warn(result.Error)
		trx.Rollback()
		return ctx.
			Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{
				"error": "Could not create currency record",
			})
	}

	// Create currency record relationship
	currencyCurrencyRecord = models.CurrencyRecord_User{
		ID:               c.UUIDGenerator.Generate(),
		CurrencyRecordId: currencyRecord.ID,
		CurrencyId:       currencyId,
		UserId:           user.ID,
	}

	result = trx.
		Create(&currencyCurrencyRecord)

	if result.Error != nil {
		log.Warn(result.Error)
		trx.Rollback()
		return ctx.
			Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{
				"error": "Could not create currency record relationship",
			})
	}

	trx.Commit()

	return ctx.
		Status(fiber.StatusOK).
		JSON(currencyRecord)
}
