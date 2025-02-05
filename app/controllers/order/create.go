package order_controller

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/taako-502/go-mongodb-sample/app/infrastructure"
	"github.com/taako-502/go-mongodb-sample/app/infrastructure/customer_infrastructure"
	"github.com/taako-502/go-mongodb-sample/app/infrastructure/order_infrastructure"
	"github.com/taako-502/go-mongodb-sample/app/infrastructure/transaction_manager"
	"github.com/taako-502/go-mongodb-sample/app/usecase/order_usecase"
	"go.mongodb.org/mongo-driver/v2/bson"

	"github.com/labstack/echo/v4"
)

type newCreate struct {
	CustomerId bson.ObjectID `json:"customerId" validate:"required"`
	// 2006-01-02
	OrderDetails []orderDetail `json:"orderDetails" validate:"required"`
	OrderDate    string        `json:"orderDate" validate:"required"`
	Status       string        `json:"status" validate:"required"`
}

type orderDetail struct {
	ProductID bson.ObjectID `json:"productId" validate:"required"`
	Quantity  int           `json:"quantity" validate:"required"`
	Price     float64       `json:"price" validate:"required"`
}

func (oo OrderController) Create(c echo.Context) error {
	request := new(newCreate)
	if err := c.Bind(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	orderDate, err := time.Parse("2006-01-02", request.OrderDate)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// TODO: DTOを作成する処理はメソッドに切り出す
	orderDetails := make([]order_usecase.OrderDetailDTO, len(request.OrderDetails))
	for i, v := range request.OrderDetails {
		orderDetails[i] = order_usecase.OrderDetailDTO{
			ProductID: v.ProductID,
			Quantity:  v.Quantity,
			Price:     v.Price,
		}
	}
	dto := order_usecase.CreateDTO{
		CustomerID:   request.CustomerId,
		OrderDetails: orderDetails,
		OrderDate:    orderDate,
		Status:       request.Status,
	}

	// コンテキストを設定
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	dbm, err := infrastructure.NewMongoDBManager(ctx, oo.ConnectionString)
	if err != nil {
		return fmt.Errorf("NewMongoDBManager: %w", err)
	}
	defer dbm.Client.Disconnect(ctx)

	o := order_usecase.NewOrderService()
	tm := transaction_manager.NewMongoTransactionManager(ctx, dbm.Client)
	cc := customer_infrastructure.NewCustomerRepository(ctx, dbm.Client.Database(oo.DBName))
	oi := order_infrastructure.NewOrderRepository(ctx, dbm.Client.Database(oo.DBName))
	if err := o.Create(tm, cc, oi, dto); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, "success")
}
