package order_usecase

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type OrderDetailDTO struct {
	ProductID bson.ObjectID
	Quantity  int
	Price     float64
}

type CreateDTO struct {
	CustomerID   bson.ObjectID
	OrderDetails []OrderDetailDTO
	OrderDate    time.Time
	Status       string
}
