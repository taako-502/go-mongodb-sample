package order_usecase

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderDetailDTO struct {
	ProductID primitive.ObjectID
	Quantity  int
	Price     float64
}

type CreateDTO struct {
	CustomerID   primitive.ObjectID
	OrderDetails []OrderDetailDTO
	OrderDate    time.Time
	Status       string
}
