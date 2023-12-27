package model

import (
	order_infrastructure "go-mongodb-sample/app/infrastructures/orders"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderDetail struct {
	ProductID primitive.ObjectID
	Quantity  int
	Price     float64
}

type Order struct {
	ID           primitive.ObjectID
	CustomerID   primitive.ObjectID
	OrderDetails []OrderDetail
	OrderDate    time.Time
	TotalAmount  float64
	Status       string
}

type OrderAdapter interface {
	Create(dto *order_infrastructure.OrderDTO) (*order_infrastructure.OrderDTO, error)
}
