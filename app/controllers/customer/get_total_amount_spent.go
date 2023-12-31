package customer_controller

import (
	"context"
	customer_usecase "go-mongodb-sample/app/usecase/customer"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (cc CostumerController) GetTotalAmountSpent(c echo.Context) error {
	ID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// コンテキストを設定
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cc.ConnectionString))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	defer client.Disconnect(ctx)

	cs := customer_usecase.NewCustomerService(ctx, cc.ConnectionString, cc.DBName)
	amount, err := cs.GetTotalAmountSpent(ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, strconv.FormatFloat(amount, 'f', -1, 64))
}
