package product_controller

import (
	"context"
	product_infrastructure "go-mongodb-sample/app/infrastructures/products"
	product_usecase "go-mongodb-sample/app/usecase/product"
	"time"

	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type newCreatePromotion struct {
	Name               string    `json:"name" validate:"required"`
	Description        string    `json:"description" validate:"required"`
	Price              float64   `json:"price" validate:"required"`
	Stock              int       `json:"stock" validate:"required"`
	Category           string    `json:"category" validate:"required"`
	PromotionExpiresAt time.Time `json:"promotionExpiresAt" validate:"required"`
}

func (pc ProductController) CreatePromotion(c echo.Context) error {
	request := new(newCreatePromotion)
	if err := c.Bind(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// TODO: タイムアウトしないようにした方がいいか検討
	// NOTE: タイムアウトさせないならcontext.TODOを使えばよい？
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(pc.ConnectionString))
	if err != nil {
		return errors.Wrap(err, "mongo.Connect")
	}
	defer client.Disconnect(ctx)

	pi := product_infrastructure.NewProductRepository(ctx, client.Database(pc.DBName))
	pu := product_usecase.NewProductService(ctx, pc.ConnectionString, pc.DBName)
	dto := product_usecase.NewPromotionProductDTO(
		request.Name,
		request.Description,
		request.Price,
		request.Stock,
		request.Category,
		request.PromotionExpiresAt,
	)
	if err := pu.CreatePromotion(pi, dto); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "success")
}
