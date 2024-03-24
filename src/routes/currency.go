package routes

import (
	controllers "financial-parsing/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func CurrencyRoutes(router fiber.Router) fiber.Router {
	currenciesRouter := router.Group("/currencies")

	var currencyController controllers.BaseController = controllers.CurrencyController{}

	currenciesRouter.Get("/", currencyController.GetAll)
	currenciesRouter.Get("/:id", currencyController.GetById)
	currenciesRouter.Post("/", currencyController.Create)
	currenciesRouter.Delete("/", currencyController.Delete)
	currenciesRouter.Put("/:id", currencyController.Update)

	return currenciesRouter
}
