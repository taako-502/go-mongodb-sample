package order_infrastructure

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderDetailDTO struct {
	ProductID primitive.ObjectID `bson:"product_id"`
	Quantity  int                `bson:"quantity"`
	Price     float64            `bson:"price"`
}

type OrderDTO struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	CustomerID   primitive.ObjectID `bson:"customer_id"`
	OrderDetails []OrderDetailDTO   `bson:"order_details"`
	OrderDate    time.Time          `bson:"order_date"`
	TotalAmount  float64            `bson:"total_amount"`
	Status       string             `bson:"status"`
}

func NewOrderDetailDTO(productID primitive.ObjectID, quantity int, price float64) *OrderDetailDTO {
	return &OrderDetailDTO{
		ProductID: productID,
		Quantity:  quantity,
		Price:     price,
	}
}

func NewOrderDTO(CustomerID primitive.ObjectID, OrderDetails []OrderDetailDTO, OrderDate time.Time, TotalAmount float64, Status string) *OrderDTO {
	return &OrderDTO{
		CustomerID:   CustomerID,
		OrderDetails: OrderDetails,
		OrderDate:    OrderDate,
		TotalAmount:  TotalAmount,
		Status:       Status,
	}
}
