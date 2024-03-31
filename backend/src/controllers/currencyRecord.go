package controllers

import (
	// "fmt"
	// "strings"

	models "financial-parsing/src/domain/models"
	usecases "financial-parsing/src/domain/usecases"
	protocols "financial-parsing/src/protocols"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
)

type CurrencyRecordController struct {
	UUIDGenerator          protocols.UUIDGenerator
	ValidateCurrencyRecord usecases.ValidateModel[models.CurrencyRecord]

	Connection *gorm.DB
}

func (c CurrencyRecordController) GetAll(ctx *fiber.Ctx) error {
	log.Info("CurrencyRecord - Get All")

	currencyId := ctx.Params("currencyId")
	var currencyRecords []models.CurrencyRecord

	result := c.Connection.
		Table("currency_records").
		Joins("JOIN currency_currency_records ON currency_currency_records.currency_record_id = currency_records.id").
		Find(&currencyRecords, "currency_currency_records.currency_id = ?", currencyId)

	if result.Error != nil {
		log.Warn(result.Error)
		return ctx.
			Status(fiber.StatusInternalServerError).
			SendString("Error when getting all currency records for specified currency")
	}

	return ctx.JSON(currencyRecords)
}

func (c CurrencyRecordController) GetById(ctx *fiber.Ctx) error {
	log.Info("CurrencyRecord - Get By Id")

	currencyId := ctx.Params("currencyId")
	id := ctx.Params("id")

	var currencyRecord models.CurrencyRecord

	result := c.Connection.
		Table("currency_records").
		Joins("JOIN currency_currency_records ON currency_currency_records.currency_record_id = currency_records.id").
		First(&currencyRecord, "currency_currency_records.currency_id = ? AND currency_currency_records.currency_record_id = ?", currencyId, id)

	if result.Error != nil {
		log.Warn(result.Error)
		return ctx.
			Status(fiber.StatusInternalServerError).
			SendString("Error when getting currency record by id for specified currency")
	}

	return ctx.JSON(currencyRecord)
}

func (c CurrencyRecordController) Update(ctx *fiber.Ctx) error {
	log.Info("CurrencyRecord - Update")

	// Getting existing currency to update
	id := ctx.Params("id")
	currencyId := ctx.Params("currencyId")
	var existingCurrencyRecord models.CurrencyRecord

	result := c.Connection.
		Table("currency_records").
		Joins("JOIN currency_currency_records ON currency_currency_records.currency_record_id = currency_records.id").
		First(
			&existingCurrencyRecord,
			"currency_currency_records.currency_id = ? AND currency_currency_records.currency_record_id = ?",
			currencyId,
			id,
		)

	if existingCurrencyRecord.CreatedAt.IsZero() {
		return ctx.
			Status(fiber.StatusNotFound).
			SendString("Could not find currency record by id for update")
	}

	if result.Error != nil {
		log.Warn("Error when searching existing currency record to update", result.Error)
		return ctx.
			Status(fiber.StatusInternalServerError).
			SendString("Error when searching for existing currency record to update")
	}

	currencyRecord := new(models.CurrencyRecord)
	if err := ctx.BodyParser(&currencyRecord); err != nil {
		log.Warn("Error when parsing", err)
		return ctx.
			Status(fiber.StatusBadRequest).
			SendString("Could not parse request body to CurrencyRecord")
	}

	existingCurrencyRecord.RecordDate = currencyRecord.RecordDate
	existingCurrencyRecord.Value = currencyRecord.Value

	trx := c.Connection.Begin()
	updateResult := trx.Save(&existingCurrencyRecord)

	if updateResult.Error != nil {
		log.Warn("Error when updating", updateResult.Error)
		trx.Rollback()
		return ctx.
			Status(fiber.StatusInternalServerError).
			SendString("Error when updating CurrencyRecord by id")
	}
	trx.Commit()

	return ctx.
		Status(fiber.StatusOK).
		JSON(existingCurrencyRecord)
}

func (c CurrencyRecordController) Delete(ctx *fiber.Ctx) error {
	log.Info("CurrencyRecord - Delete")

	currencyId := ctx.Params("currencyId")
	ids := ctx.Query("ids")
	idsToDelete := strings.Split(ids, ",")
	log.Info("IDS to delete: ", idsToDelete)

	var deleted []models.CurrencyRecord

	// Delete currency record
	result := c.Connection.
		Table("currency_currency_records").
		Joins("JOIN currency_currency_records ON currency_currency_records.currency_record_id = currency_records.id").
		Joins("JOIN currency_user ON currency_currency_records.currency_id = currency_user.currency_id").
		Delete(&deleted, "currency_currency_records.currency_id = ?", currencyId)

	if result.Error != nil {
		log.Warn(result.Error)
		return ctx.
			Status(fiber.StatusInternalServerError).
			SendString("Error when deleting currency records by ids")
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func (c CurrencyRecordController) Create(ctx *fiber.Ctx) error {
	log.Info("CurrencyRecord - Create")

	currencyId := ctx.Params("currencyId")
	var body models.CurrencyRecord

	// Parse request body as currency record
	if err := ctx.BodyParser(&body); err != nil {
		log.Warn("Error when parsing body to currency record", err)
		return ctx.
			Status(fiber.StatusBadRequest).
			SendString("Could not parse specified request body to CurrencyRecord")
	}
	body.ID = c.UUIDGenerator.Generate()

	trx := c.Connection.Begin()

	// Create currency record
	result := trx.
		Table("currency_records").
		Create(&body)

	if result.Error != nil {
		log.Warn(result.Error)
		trx.AddError(result.Error)
		return ctx.
			Status(fiber.StatusInternalServerError).
			SendString("Error when creating currency record")
	}

	// Create relationship between currency record and currency
	result = trx.
		Table("currency_currency_records").
		Create(&models.Currency_CurrencyRecord{
			ID:               c.UUIDGenerator.Generate(),
			CurrencyId:       currencyId,
			CurrencyRecordId: body.ID,
		})

	if result.Error != nil {
		log.Warn(result.Error)
		trx.Rollback()
		return ctx.
			Status(fiber.StatusInternalServerError).
			SendString("Error when creating relationship between currency and currency record")
	}

	// Save changes
	trx.Commit()

	return ctx.
		Status(fiber.StatusCreated).
		JSON(body)
}
