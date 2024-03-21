package controllers

import "github.com/gofiber/fiber/v2"

type TransactionController struct{}

func (t TransactionController) GetAll(c *fiber.Ctx) error {
	return c.SendString("Transaction - Get All")
}

func (t TransactionController) GetById(c *fiber.Ctx) error {
	return c.SendString("Transaction - Get By ID")
}

func (t TransactionController) Create(c *fiber.Ctx) error {
	return c.SendString("Transaction - Create")
}

func (t TransactionController) Update(c *fiber.Ctx) error {
	return c.SendString("Transaction - Update")
}

func (t TransactionController) Delete(c *fiber.Ctx) error {
	return c.SendString("Transaction - Delete")
}
