package customer_controller

import (
	"context"
	"errors"
	customer_infrastructure "go-mongodb-sample/app/infrastructures/customers"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type NewCreate struct {
	Name    string `json:"name" validate:"required"`
	Email   string `json:"email" validate:"required"`
	Address string `json:"address" validate:"required"`
}

func (cc CostumerController) Create(c echo.Context) error {
	request := new(NewCreate)
	if err := c.Bind(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// コンテキストを設定
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cc.ConnectionString))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	defer client.Disconnect(ctx)

	collection := client.Database(cc.DBName).Collection(cc.CollectionName)
	ci := customer_infrastructure.NewCustomer(ctx, collection)
	dto := customer_infrastructure.NewCustomerDTO(request.Name, request.Email, request.Address, nil)
	customer, err := ci.Create(dto)
	if err != nil {
		if errors.Is(err, customer_infrastructure.ErrCustomerDuplicate) {
			return echo.NewHTTPError(http.StatusConflict, err.Error())
		} else {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	return c.JSON(http.StatusOK, customer)
}
