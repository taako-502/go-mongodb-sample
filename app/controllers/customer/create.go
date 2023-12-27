package customer_controller

import (
	customer_infrastructure "go-mongodb-sample/app/infrastructures/customers"
	"log"
	"net/http"

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
	ci := customer_infrastructure.NewCustomer(cc.Ctx, cc.Collection)
	dto := customer_infrastructure.NewCustomerDTO(request.Name, request.Email, request.Address, nil)
	customer, err := ci.Create(dto)
	if err != nil {
		log.Fatal(err)
	}
	return c.JSON(http.StatusOK, customer)
}
