package controllers

import (
	"fmt"
	"strings"

	models "financial-parsing/src/domain/models"
	usecases "financial-parsing/src/domain/usecases"
	protocols "financial-parsing/src/protocols"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type CurrencyController struct {
	DatabaseAdapter  protocols.DatabaseAdapter[models.Currency]
	UUIDGenerator    protocols.UUIDGenerator
	ValidateCurrency usecases.ValidateModel[models.Currency]
}

func (c CurrencyController) GetAll(ctx *fiber.Ctx) error {
	ctx.SendString("Currency - Get All")
	currencies, err := c.DatabaseAdapter.GetAll()

	if err != nil {
		return ctx.
			Status(fiber.StatusInternalServerError).
			SendString("Currency Controller - Error when getting all currencies")
	}

	return ctx.JSON(currencies)
}

func (c CurrencyController) GetById(ctx *fiber.Ctx) error {
	ctx.SendString("Currency - Get By Id")
	id := ctx.Params("id")
	currency, err := c.DatabaseAdapter.GetById(id)

	if err != nil {
		return ctx.
			Status(fiber.StatusInternalServerError).
			SendString(
				fmt.Sprintf(
					"Currency Controller - Error when getting currency by id - %s",
					err,
				),
			)
	}

	return ctx.JSON(currency)
}

func (c CurrencyController) Update(ctx *fiber.Ctx) error {
	ctx.SendString("Currency - Update")

	id := ctx.Params("id")
	currency := new(models.Currency)

	if err := ctx.BodyParser(&currency); err != nil {
		log.Warn("Error when parsing", err)
		return ctx.
			Status(fiber.StatusBadRequest).
			SendString("Currency Controller - Could not parse request body to currency")
	}

	// Validate before updating in DB
	if err := c.ValidateCurrency.Validate(*currency); err != nil {
		return ctx.
			Status(fiber.StatusBadRequest).
			SendString(err.Error())
	}

	fields := []string{"Name", "Representation"}
	updated, err := c.DatabaseAdapter.UpdateById(id, currency, fields)
	if err != nil {
		return ctx.
			Status(fiber.StatusInternalServerError).
			SendString(
				fmt.Sprintf(
					"Currency Controller - Error when updating currency by id - %s",
					err,
				),
			)
	}

	return ctx.
		Status(fiber.StatusOK).
		JSON(updated)
}

func (c CurrencyController) Delete(ctx *fiber.Ctx) error {
	ctx.SendString("Currency - Delete")

	ids := ctx.Query("ids")
	idsToDelete := strings.Split(ids, ",")
	log.Info("IDS to delete", ids, idsToDelete)
	err := c.DatabaseAdapter.DeleteByIds(idsToDelete)

	if err != nil {
		return ctx.
			Status(fiber.StatusInternalServerError).
			SendString(
				fmt.Sprintf(
					"Currency Controller - Error when deleting currency by ids - %s",
					err,
				),
			)
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func (c CurrencyController) Create(ctx *fiber.Ctx) error {
	ctx.SendString("Currency - Create")

	// Create currency struct from JSON body
	body := new(models.Currency)
	if err := ctx.BodyParser(&body); err != nil {
		log.Warn("Error when parsing", err)
		return ctx.
			Status(fiber.StatusBadRequest).
			SendString("Currency Controller - Could not parse request body to currency")
	}
	body.ID = c.UUIDGenerator.Generate()

	// Validate before creating in DB
	if err := c.ValidateCurrency.Validate(*body); err != nil {
		return ctx.
			Status(fiber.StatusBadRequest).
			SendString(err.Error())
	}

	fields := []string{"ID", "Name", "Representation"}
	created, err := c.DatabaseAdapter.Create(body, fields)

	if err != nil {
		return ctx.
			Status(fiber.StatusInternalServerError).
			SendString("Currency Controller - Error when creating currency")
	}

	return ctx.
		Status(fiber.StatusCreated).
		JSON(created)
}
