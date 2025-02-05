package product_controller

import (
	"context"
	"fmt"
	"time"

	"github.com/taako-502/go-mongodb-sample/app/infrastructure"
	"github.com/taako-502/go-mongodb-sample/app/infrastructure/product_infrastructure"

	"net/http"

	"github.com/labstack/echo/v4"
)

type newCreate struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
	Stock       int     `json:"stock" validate:"required"`
	Category    string  `json:"category" validate:"required"`
}

func (pc ProductController) Create(c echo.Context) error {
	request := new(newCreate)
	if err := c.Bind(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// コンテキストを設定
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	dbm, err := infrastructure.NewMongoDBManager(ctx, pc.ConnectionString)
	if err != nil {
		return fmt.Errorf("NewMongoDBManager: %w", err)
	}
	defer dbm.Client.Disconnect(ctx)

	pi := product_infrastructure.NewProductRepository(ctx, dbm.Client.Database(pc.DBName))
	dto := product_infrastructure.NewProductDTO(
		request.Name,
		request.Description,
		request.Price,
		request.Stock,
		request.Category,
	)
	product, err := pi.Create(dto)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, product)
}
