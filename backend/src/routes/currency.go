package routes

import (
	controllers "financial-parsing/src/controllers"

	validatemodel "financial-parsing/src/data/usecases/validateModel"
	uuidgenerator "financial-parsing/src/utils/uuidGenerator"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CurrencyRoutes(router fiber.Router, connection *gorm.DB) fiber.Router {
	currenciesRouter := router.Group("/currencies")

	uuidGenerator := uuidgenerator.GoogleUUIDGenerator{}

	var currencyController controllers.BaseController = controllers.CurrencyController{
		UUIDGenerator: uuidGenerator,
		Connection:    connection,
		ValidateCurrency: validatemodel.ValidateCurrency{
			UUIDGenerator: uuidGenerator,
		},
	}

	currenciesRouter.Get("/", currencyController.GetAll)
	currenciesRouter.Get("/:id", currencyController.GetById)
	currenciesRouter.Post("/", currencyController.Create)
	currenciesRouter.Delete("/", currencyController.Delete)
	currenciesRouter.Put("/:id", currencyController.Update)

	return currenciesRouter
}
