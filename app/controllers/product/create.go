package product_controller

import (
	"context"
	product_infrastructure "go-mongodb-sample/app/infrastructures/products"
	"time"

	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(pc.ConnectionString))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	defer client.Disconnect(ctx)

	collection := client.Database(pc.DBName).Collection(pc.CollectionName)
	pi := product_infrastructure.NewProduct(ctx, collection)
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
