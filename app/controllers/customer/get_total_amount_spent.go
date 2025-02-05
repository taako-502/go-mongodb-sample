package customer_controller

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/taako-502/go-mongodb-sample/app/infrastructure"
	"github.com/taako-502/go-mongodb-sample/app/infrastructure/customer_infrastructure"
	"github.com/taako-502/go-mongodb-sample/app/infrastructure/order_infrastructure"
	customer_usecase "github.com/taako-502/go-mongodb-sample/app/usecase/customer"
	"go.mongodb.org/mongo-driver/v2/bson"

	"github.com/labstack/echo/v4"
)

func (cc CostumerController) GetTotalAmountSpent(c echo.Context) error {
	ID, err := bson.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// コンテキストを設定
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	dbm, err := infrastructure.NewMongoDBManager(ctx, cc.ConnectionString)
	if err != nil {
		return fmt.Errorf("NewMongoDBManager: %w", err)
	}
	defer dbm.Client.Disconnect(ctx)

	cs := customer_usecase.NewCustomerService()
	ci := customer_infrastructure.NewCustomerRepository(ctx, dbm.Client.Database(cc.DBName))
	or := order_infrastructure.NewOrderRepository(ctx, dbm.Client.Database(cc.DBName))
	amount, err := cs.GetTotalAmountSpent(ci, or, ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, strconv.FormatFloat(amount, 'f', -1, 64))
}
