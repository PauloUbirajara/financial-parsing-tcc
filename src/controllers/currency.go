package controllers

import (
	"fmt"
	"strings"

	models "financial-parsing/src/domain/models"
	protocols "financial-parsing/src/protocols"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type CurrencyController struct {
	DatabaseAdapter protocols.DatabaseAdapter[models.Currency]
}

func (c CurrencyController) GetAll(ctx *fiber.Ctx) error {
	ctx.SendString("Currency - Get All")
	currencies, err := c.DatabaseAdapter.GetAll()

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString("Currency Controller - Error when getting all currencies")
	}

	return ctx.JSON(currencies)
}

func (c CurrencyController) GetById(ctx *fiber.Ctx) error {
	ctx.SendString("Currency - Get By Id")
	id := ctx.Params("id")
	currency, err := c.DatabaseAdapter.GetById(id)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(
			fmt.Sprintf("Currency Controller - Error when getting currency by id - %s", err),
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
		return ctx.Status(fiber.StatusBadRequest).SendString("Currency Controller - Could not parse request body to currency")
	}

	updated, err := c.DatabaseAdapter.UpdateById(id, currency)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(
			fmt.Sprintf("Currency Controller - Error when updating currency by id - %s", err),
		)
	}

	return ctx.Status(fiber.StatusOK).JSON(updated)
}

func (c CurrencyController) Delete(ctx *fiber.Ctx) error {
	ctx.SendString("Currency - Delete")

	ids := ctx.Query("ids")
	idsToDelete := strings.Split(ids, ",")
	log.Info("IDS to delete", ids, idsToDelete)
	err := c.DatabaseAdapter.DeleteByIds(idsToDelete)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(
			fmt.Sprintf("Currency Controller - Error when deleting currency by ids - %s", err),
		)
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func (c CurrencyController) Create(ctx *fiber.Ctx) error {
	ctx.SendString("Currency - Create")

	body := new(models.Currency)
	if err := ctx.BodyParser(&body); err != nil {
		log.Warn("Error when parsing", err)
		return ctx.Status(fiber.StatusBadRequest).SendString("Currency Controller - Could not parse request body to currency")
	}

	fields := []string{"Name", "Representation"}
	created, err := c.DatabaseAdapter.Create(body, fields)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString("Currency Controller - Error when creating currency")
	}

	return ctx.Status(fiber.StatusCreated).JSON(created)
}
