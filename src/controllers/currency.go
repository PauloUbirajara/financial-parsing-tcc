package controllers

import (
	"fmt"
	"strings"

	models "financial-parsing/src/domain/models"
	helpers "financial-parsing/src/helpers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type CurrencyController struct {
}

func (t CurrencyController) GetAll(c *fiber.Ctx) error {
	c.SendString("Currency - Get All")
	db, _ := helpers.CreateConnection()
	var currencies []models.Currency
	result := db.Find(&currencies)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Currency Controller - Error when getting all currencies")
	}
	return c.JSON(currencies)
}

func (t CurrencyController) GetById(c *fiber.Ctx) error {
	c.SendString("Currency - Get By Id")
	db, _ := helpers.CreateConnection()
	var currency models.Currency
	id := c.Params("id")
	result := db.First(&currency, "id = ?", id)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(
			fmt.Sprintf("Currency Controller - Error when getting currency by id - %s", result.Error),
		)
	}

	return c.JSON(currency)
}

func (t CurrencyController) Update(c *fiber.Ctx) error {
	c.SendString("Currency - Update")
	db, _ := helpers.CreateConnection()
	var currency models.Currency
	id := c.Params("id")
	result := db.First(&currency, "id = ?", id)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(
			fmt.Sprintf("Currency Controller - Error when searching currency by id for update - %s", result.Error),
		)
	}

	updated := new(models.Currency)
	if err := c.BodyParser(&updated); err != nil {
		log.Warn("Error when parsing", err)
		return c.Status(fiber.StatusBadRequest).SendString("Currency Controller - Could not parse request body to currency")
	}
	updateResult := db.Model(&currency).Updates(&updated)
	if updateResult.Error != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(
			fmt.Sprintf("Currency Controller - Error when updating currency by id - %s", result.Error),
		)
	}

	return c.Status(fiber.StatusOK).JSON(currency)
}

func (t CurrencyController) Delete(c *fiber.Ctx) error {
	c.SendString("Currency - Delete")
	db, _ := helpers.CreateConnection()

	ids := c.Query("ids")
	idsToDelete := strings.Split(ids, ",")
	log.Info("IDS to delete", ids, idsToDelete)
	result := db.Delete(&models.Currency{}, idsToDelete)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(
			fmt.Sprintf("Currency Controller - Error when deleting currency by ids - %s", result.Error),
		)
	}

	return c.SendStatus(fiber.StatusOK)
}

func (t CurrencyController) Create(c *fiber.Ctx) error {
	c.SendString("Currency - Create")
	db, _ := helpers.CreateConnection()

	body := new(models.Currency)
	if err := c.BodyParser(&body); err != nil {
		log.Warn("Error when parsing", err)
		return c.Status(fiber.StatusBadRequest).SendString("Currency Controller - Could not parse request body to currency")
	}

	result := db.Select("Name", "Representation").Create(&body)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Currency Controller - Error when creating currency")
	}

	return c.Status(fiber.StatusCreated).JSON(body)
}
