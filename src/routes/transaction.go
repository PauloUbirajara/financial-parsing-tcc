package routes

import (
	controllers "financial-parsing/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func TransactionRoutes(router fiber.Router) fiber.Router {
	transactionsRouter := router.Group("/transactions")
	var transactionController controllers.BaseController = controllers.TransactionController{}

	transactionsRouter.Get("/", transactionController.GetAll)
	transactionsRouter.Get("/:id", transactionController.GetById)
	transactionsRouter.Post("/", transactionController.Create)
	transactionsRouter.Put("/:id", transactionController.Update)
	transactionsRouter.Delete("/:id", transactionController.Delete)

	return transactionsRouter
}
