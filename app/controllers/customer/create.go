package customer_controller

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/taako-502/go-mongodb-sample/app/infrastructure"
	"github.com/taako-502/go-mongodb-sample/app/infrastructure/customer_infrastructure"

	"github.com/labstack/echo/v4"
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
	dbm, err := infrastructure.NewMongoDBManager(ctx, cc.ConnectionString)
	if err != nil {
		return fmt.Errorf("NewMongoDBManager: %w", err)
	}
	defer dbm.Client.Disconnect(ctx)

	ci := customer_infrastructure.NewCustomerRepository(ctx, dbm.Client.Database(cc.DBName))
	dto := customer_infrastructure.NewCustomerDTO(request.Name, request.Email, request.Address)
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
