package order_controller

import (
	"context"
	"errors"
	"go-mongodb-sample/app/infrastructure/order_infrastructure"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (oo OrderController) FindByCustomerID(c echo.Context) error {
	customerID, err := primitive.ObjectIDFromHex(c.Param("customer_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// コンテキストを設定
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(oo.ConnectionString))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	defer client.Disconnect(ctx)

	oi := order_infrastructure.NewOrderRepository(ctx, client.Database(oo.DBName))
	order, err := oi.FindByCustomerID(customerID)
	if errors.Is(err, order_infrastructure.ErrOrderNotFound) {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, order)
}
