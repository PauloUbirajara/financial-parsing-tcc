package controllers_test

import (
	"errors"
	"fmt"
	"testing"

	controllers "financial-parsing/src/controllers"
	models "financial-parsing/src/domain/models"

	"net/http/httptest"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

type DatabaseAdapterStub[T interface{}] struct {
	shouldFail bool
}

func (d *DatabaseAdapterStub[T]) GetAll() (*[]T, error) {
	if d.shouldFail {
		return nil, errors.New("Failed on GetAll")
	}
	results := make([]T, 5)
	return &results, nil
}

func (d DatabaseAdapterStub[T]) GetById(id string) (*T, error) {
	if d.shouldFail {
		return nil, errors.New("Failed on GetByID")
	}
	results, _ := d.GetAll()
	return &(*results)[0], nil
}

func (d DatabaseAdapterStub[T]) Create(model *T, fieldNames []string) (*T, error) {
	if d.shouldFail {
		return nil, errors.New("Failed on Create")
	}
	return model, nil
}

func (d DatabaseAdapterStub[T]) DeleteByIds(ids []string) error {
	if d.shouldFail {
		return errors.New("Failed on DeleteByIds")
	}
	return nil
}

func (d DatabaseAdapterStub[T]) UpdateById(id string, updated *T, fields []string) (*T, error) {
	if d.shouldFail {
		return nil, errors.New("Failed on UpdateById")
	}
	return updated, nil
}

type CurrencyControllerTestData struct {
	sut             controllers.BaseController
	app             *fiber.App
	baseUrl         string
	databaseAdapter DatabaseAdapterStub[models.Currency]
}

func NewCurrencyControllerTestData(dbShouldFail bool) CurrencyControllerTestData {
	databaseAdapter := DatabaseAdapterStub[models.Currency]{
		shouldFail: dbShouldFail,
	}
	sut := controllers.CurrencyController{
		DatabaseAdapter: &databaseAdapter,
	}

	app := fiber.New()

	currenciesRouter := app.Group("/currencies")
	currenciesRouter.Get("/", sut.GetAll)
	currenciesRouter.Get("/:id", sut.GetById)
	currenciesRouter.Post("/", sut.Create)
	currenciesRouter.Delete("/", sut.Delete)
	currenciesRouter.Put("/:id", sut.Update)

	return CurrencyControllerTestData{
		baseUrl:         app.AcquireCtx(&fasthttp.RequestCtx{}).BaseURL(),
		databaseAdapter: databaseAdapter,
		sut:             sut,
		app:             app,
	}
}

func TestCurrencyControllerShouldReturnOKOnGetAll(t *testing.T) {
	testData := NewCurrencyControllerTestData(false)

	getAllUrl := fmt.Sprintf("%s/currencies", testData.baseUrl)
	req := httptest.NewRequest("GET", getAllUrl, nil)
	res, err := testData.app.Test(req)

	if err != nil {
		t.Fatal("Failed to call CurrencyController GetAll")
	}

	if res.StatusCode != fiber.StatusOK {
		t.Fatal("Did not return OK on CurrencyController GetAll")
	}
}

// // TODO TESTAR O GET ALL COM E SEM ERRO
// func TestCurrencyControllerShouldReturnInternalErrorOnGetAll(t *testing.T) {
// 	testData := NewCurrencyControllerTestData(true)
// 	getAllUrl := fmt.Sprintf("%s/currencies", testData.baseUrl)
// 	req := httptest.NewRequest("GET", getAllUrl, nil)
// 	res, err := testData.app.Test(req)
//
// 	if err == nil {
// 		t.Fatal("Did not throw error when failing to call CurrencyController GetAll")
// 	}
//
// 	if res.StatusCode != fiber.StatusInternalServerError {
// 		t.Fatal("Did not return InternalServerError on CurrencyController GetAll")
// 	}
// }
