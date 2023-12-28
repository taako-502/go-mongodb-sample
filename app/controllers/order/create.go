package order_controller

import (
	"context"
	"net/http"
	"time"

	order_infrastructure "go-mongodb-sample/app/infrastructures/orders"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type newCreate struct {
	CustomerId primitive.ObjectID `json:"customerId" validate:"required"`
	// 2006-01-02
	OrderDetails []orderDetail `json:"orderDetails" validate:"required"`
	OrderDate    string        `json:"orderDate" validate:"required"`
	TotalAmount  float64       `json:"totalAmount" validate:"required"`
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

	// コンテキストを設定
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(oo.ConnectionString))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	defer client.Disconnect(ctx)

	collection := client.Database(oo.DBName).Collection(oo.CollectionName)
	oi := order_infrastructure.NewOrderRepository(ctx, collection)
	orderDetails := make([]order_infrastructure.OrderDetailDTO, len(request.OrderDetails))
	for i, v := range request.OrderDetails {
		d := order_infrastructure.NewOrderDetailDTO(v.ProductID, v.Quantity, v.Price)
		orderDetails[i] = *d
	}
	dto := order_infrastructure.NewOrderDTO(
		request.CustomerId,
		orderDetails,
		orderDate,
		request.TotalAmount,
		request.Status,
	)
	order, err := oi.Create(dto)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, order)
}
