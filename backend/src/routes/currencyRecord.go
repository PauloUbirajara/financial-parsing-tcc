package routes

import (
	controllers "financial-parsing/src/controllers"
	// models "financial-parsing/src/domain/models"

	validatemodel "financial-parsing/src/data/usecases/validateModel"
	// databaseadapter "financial-parsing/src/utils/databaseAdapter"
	uuidgenerator "financial-parsing/src/utils/uuidGenerator"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CurrencyRecordsRoutes(router fiber.Router, connection *gorm.DB) fiber.Router {
	currencyRecordRouter := router.Group("/currencies/:currencyId/records")

	uuidGenerator := uuidgenerator.GoogleUUIDGenerator{}

	var currencyRecordController controllers.BaseController = controllers.CurrencyRecordController{
		UUIDGenerator: uuidGenerator,
		Connection:    connection,
		ValidateCurrencyRecord: validatemodel.ValidateCurrencyRecord{
			UUIDGenerator: uuidGenerator,
		},
	}

	currencyRecordRouter.Get("/", currencyRecordController.GetAll)
	currencyRecordRouter.Get("/:id", currencyRecordController.GetById)
	currencyRecordRouter.Post("/", currencyRecordController.Create)
	currencyRecordRouter.Delete("/", currencyRecordController.Delete)
	currencyRecordRouter.Put("/:id", currencyRecordController.Update)

	return currencyRecordRouter
}
