package controllers

import (
	"time"

	models "financial-parsing/src/domain/models"
	protocols "financial-parsing/src/protocols"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type TransactionController struct {
	DatabaseAdapter protocols.DatabaseAdapter
	UUIDGenerator   protocols.UUIDGenerator
}

func (t TransactionController) GetAll(c *fiber.Ctx) error {
	c.SendString("Transaction - Get All")
	transactions, err := t.DatabaseAdapter.GetAll()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error when getting all transactions")
	}

	return c.JSON(transactions)
}

func (t TransactionController) GetById(c *fiber.Ctx) error {
	c.SendString("Transaction - Get By ID")

	id := c.Params("id")
	transaction, err := t.DatabaseAdapter.GetById(id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error when getting transaction by id")
	}

	return c.JSON(transaction)
}

func (t TransactionController) Create(c *fiber.Ctx) error {
	c.SendString("Transaction - Create")

	transaction := new(models.Transaction)
	if err := c.BodyParser(transaction); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Error when creating, could not parse request body to transaction")
	}

	transaction.Id = t.UUIDGenerator.Generate()
	transaction.CreatedAt = time.Now()

	err := t.DatabaseAdapter.Create(transaction.Id, *transaction)
	if err != nil {
		log.Info(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error when creating, error when creating transaction")
	}

	return c.SendStatus(fiber.StatusCreated)
}

func (t TransactionController) Update(c *fiber.Ctx) error {
	c.SendString("Transaction - Update")

	transaction := new(models.Transaction)
	if err := c.BodyParser(transaction); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Error when updating, could not parse request body to transaction")
	}

	id := c.Params("id")
	updatedTransaction, err := t.DatabaseAdapter.UpdateById(id, *transaction)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error when updating, error when updating transaction")
	}

	if updatedTransaction == nil {
		return c.Status(fiber.StatusNotFound).SendString("Error when updating, did not find transaction")
	}

	return c.SendStatus(fiber.StatusOK)
}

func (t TransactionController) Delete(c *fiber.Ctx) error {
	c.SendString("Transaction - Delete")
	id := c.Params("id")
	deletedTransaction, err := t.DatabaseAdapter.DeleteById(id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error when deleting, error when deleting transaction")
	}

	if deletedTransaction == nil {
		return c.Status(fiber.StatusNotFound).SendString("Error when deleting, did not find transaction")
	}

	return c.SendStatus(fiber.StatusOK)
}
