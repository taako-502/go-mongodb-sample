package product_controller

import (
	"context"
	product_usecase "go-mongodb-sample/app/usecase/product"
	"time"

	"net/http"

	"github.com/labstack/echo/v4"
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

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pu := product_usecase.NewProductService(ctx, pc.ConnectionString, pc.DBName)
	dto := product_usecase.NewPromotionProductDTO(
		request.Name,
		request.Description,
		request.Price,
		request.Stock,
		request.Category,
		request.PromotionExpiresAt,
	)
	if err := pu.CreatePromotion(dto); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "success")
}
