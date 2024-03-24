package routes

import (
	controllers "financial-parsing/src/controllers"
	models "financial-parsing/src/domain/models"

	databaseadapter "financial-parsing/src/utils/databaseAdapter"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CurrencyRoutes(router fiber.Router, connection *gorm.DB) fiber.Router {
	currenciesRouter := router.Group("/currencies")

	var currencyController controllers.BaseController = controllers.CurrencyController{
		DatabaseAdapter: databaseadapter.GormDatabaseAdapter[models.Currency]{
			Connection: connection,
		},
	}

	currenciesRouter.Get("/", currencyController.GetAll)
	currenciesRouter.Get("/:id", currencyController.GetById)
	currenciesRouter.Post("/", currencyController.Create)
	currenciesRouter.Delete("/", currencyController.Delete)
	currenciesRouter.Put("/:id", currencyController.Update)

	return currenciesRouter
}
