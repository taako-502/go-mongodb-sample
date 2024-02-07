package order_controller

import (
	"context"
	"net/http"
	"time"

	"github.com/taako-502/go-mongodb-sample/app/infrastructure"
	"github.com/taako-502/go-mongodb-sample/app/infrastructure/order_infrastructure"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (oo OrderController) FindByCustomerID(c echo.Context) error {
	customerID, err := primitive.ObjectIDFromHex(c.Param("customer_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// コンテキストを設定
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	dbm, err := infrastructure.NewMongoDBManager(ctx, oo.ConnectionString)
	if err != nil {
		return errors.Wrap(err, "NewMongoDBManager")
	}
	defer dbm.Client.Disconnect(ctx)

	oi := order_infrastructure.NewOrderRepository(ctx, dbm.Client.Database(oo.DBName))
	order, err := oi.FindByCustomerID(customerID)
	if errors.Is(err, order_infrastructure.ErrOrderNotFound) {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, order)
}
