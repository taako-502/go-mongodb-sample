package order_controller

import (
	"context"
	"log"
	"net/http"
	"time"

	order_infrastructure "go-mongodb-sample/app/infrastructures/orders"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type newCreate struct {
	customerId   primitive.ObjectID
	OrderDetails []order_infrastructure.OrderDetailDTO
	OrderDate    time.Time
	TotalAmount  float64
	Status       string
}

func (oo OrderController) Create(c echo.Context) error {
	request := new(newCreate)
	if err := c.Bind(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// コンテキストを設定
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(oo.ConnectionString))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	collection := client.Database(oo.DBName).Collection(oo.CollectionName)
	oi := order_infrastructure.NewOrderRepository(ctx, collection)
	dto := order_infrastructure.NewOrderDTO(
		request.customerId,
		request.OrderDetails,
		request.OrderDate,
		request.TotalAmount,
		request.Status,
	)
	order, err := oi.Create(dto)
	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(http.StatusOK, order)
}
