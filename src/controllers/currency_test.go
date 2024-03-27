package controllers_test

import (
	"fmt"
	"testing"

	controllers "financial-parsing/src/controllers"
	models "financial-parsing/src/domain/models"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
	"net/http/httptest"
)

type CurrencyControllerTestData struct {
	sut     controllers.BaseController
	app     *fiber.App
	baseUrl string
}

type DatabaseAdapterStub[T interface{}] struct{}

func (d *DatabaseAdapterStub[T]) GetAll() (*[]T, error) {
	results := make([]T, 5)
	return &results, nil
}

func (d DatabaseAdapterStub[T]) GetById(id string) (*T, error) {
	results, _ := d.GetAll()
	return &(*results)[0], nil
}

func (d DatabaseAdapterStub[T]) Create(model *T, fieldNames []string) (*T, error) {
	return model, nil
}

func (d DatabaseAdapterStub[T]) DeleteByIds(ids []string) error {
	return nil
}

func (d DatabaseAdapterStub[T]) UpdateById(id string, updated *T, fields []string) (*T, error) {
	return updated, nil
}

func NewCurrencyControllerTestData() CurrencyControllerTestData {
	app := fiber.New()
	databaseAdapter := DatabaseAdapterStub[models.Currency]{}

	return CurrencyControllerTestData{
		baseUrl: app.AcquireCtx(&fasthttp.RequestCtx{}).BaseURL(),
		sut: controllers.CurrencyController{
			DatabaseAdapter: &databaseAdapter,
		},
		app: app,
	}
}

func TestCurrencyControllerShouldPassOnGetAll(t *testing.T) {
	testData := NewCurrencyControllerTestData()

	currenciesRouter := testData.app.Group("/currencies")
	currenciesRouter.Get("/", testData.sut.GetAll)

	req := httptest.NewRequest("GET", fmt.Sprintf("%s/currencies", testData.baseUrl), nil)
	res, err := testData.app.Test(req)

	if err != nil {
		t.Fatal("Failed to call CurrencyController GetAll")
	}

	if res.StatusCode != fiber.StatusOK {
		t.Fatal("Did not return OK on CurrencyController GetAll")
	}
}
