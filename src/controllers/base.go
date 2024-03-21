package controllers

import "github.com/gofiber/fiber/v2"

type BaseController interface {
	GetAll(c *fiber.Ctx) error
	GetById(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}
