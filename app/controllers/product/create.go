package customer_controller

import (
	product_infrastructure "go-mongodb-sample/app/infrastructures/products"

	"log"
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
	pi := product_infrastructure.NewProduct(pc.Ctx, pc.Collection)
	dto := product_infrastructure.NewProductDTO(
		request.Name,
		request.Description,
		request.Price,
		request.Stock,
		request.Category,
	)
	product, err := pi.Create(dto)
	if err != nil {
		log.Fatal(err)
	}
	return c.JSON(http.StatusOK, product)
}
