package order_controller

import (
	"errors"
	"net/http"

	order_infrastructure "go-mongodb-sample/app/infrastructures/orders"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (oo OrderController) FindByCustomerID(c echo.Context) error {
	customerID, err := primitive.ObjectIDFromHex(c.Param("customer_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	oi := order_infrastructure.NewOrderRepository(oo.Ctx, oo.Collection)
	order, err := oi.FindByCustomerID(customerID)
	if errors.Is(err, order_infrastructure.ErrOrderNotFound) {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, order)
}
