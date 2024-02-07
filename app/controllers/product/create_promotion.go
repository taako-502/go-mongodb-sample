package product_controller

import (
	"context"
	"time"

	"github.com/taako-502/go-mongodb-sample/app/infrastructure/product_infrastructure"
	"github.com/taako-502/go-mongodb-sample/app/usecase/product_usecase"

	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type newCreatePromotion struct {
	Name               string     `json:"name" validate:"required"`
	Description        string     `json:"description" validate:"required"`
	Price              float64    `json:"price" validate:"required"`
	Stock              int        `json:"stock" validate:"required"`
	Category           string     `json:"category" validate:"required"`
	PromotionExpiresAt *time.Time `json:"promotionExpiresAt" validate:"required"`
}

type response struct {
	ID                 primitive.ObjectID
	Name               string
	Description        string
	Price              float64
	Stock              int
	Category           string
	PromotionExpiresAt *time.Time
}

func (pc ProductController) CreatePromotion(c echo.Context) error {
	request := new(newCreatePromotion)
	if err := c.Bind(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(pc.ConnectionString))
	if err != nil {
		return errors.Wrap(err, "mongo.Connect")
	}
	defer client.Disconnect(ctx)

	pi := product_infrastructure.NewProductRepository(ctx, client.Database(pc.DBName))
	pu := product_usecase.NewProductService()
	dto := product_usecase.NewPromotionProductDTO(
		request.Name,
		request.Description,
		request.Price,
		request.Stock,
		request.Category,
		request.PromotionExpiresAt,
	)
	result, err := pu.CreatePromotion(pi, dto)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	response := response{
		ID:                 result.ID,
		Name:               result.Name,
		Description:        result.Description,
		Price:              result.Price,
		Stock:              result.Stock,
		Category:           result.Category,
		PromotionExpiresAt: result.PromotionExpiresAt,
	}

	return c.JSON(http.StatusOK, response)
}
