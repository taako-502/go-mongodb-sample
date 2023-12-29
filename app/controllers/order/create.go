package order_controller

import (
	"context"
	"net/http"
	"time"

	order_usecase "go-mongodb-sample/app/usecase/order"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type newCreate struct {
	CustomerId primitive.ObjectID `json:"customerId" validate:"required"`
	// 2006-01-02
	OrderDetails []orderDetail `json:"orderDetails" validate:"required"`
	OrderDate    string        `json:"orderDate" validate:"required"`
	Status       string        `json:"status" validate:"required"`
}

type orderDetail struct {
	ProductID primitive.ObjectID `json:"productId" validate:"required"`
	Quantity  int                `json:"quantity" validate:"required"`
	Price     float64            `json:"price" validate:"required"`
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

	o := order_usecase.NewOrderService(ctx, oo.ConnectionString, oo.DBName)
	if err := o.Create(dto); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, "success")
}
