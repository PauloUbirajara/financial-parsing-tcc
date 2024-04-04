package controllers

import (
	models "financial-parsing/src/domain/models"
	usecases "financial-parsing/src/domain/usecases"
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
	return ctx.
		Status(fiber.StatusInternalServerError).
		JSON(fiber.Map{
			"error": "Could not get all currency records",
		})
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
