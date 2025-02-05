package order_infrastructure

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type OrderDetailDTO struct {
	ProductID bson.ObjectID `bson:"product_id"`
	Quantity  int           `bson:"quantity"`
	Price     float64       `bson:"price"`
}

type OrderDTO struct {
	ID           bson.ObjectID    `bson:"_id,omitempty"`
	CustomerID   bson.ObjectID    `bson:"customer_id"`
	OrderDetails []OrderDetailDTO `bson:"order_details"`
	OrderDate    time.Time        `bson:"order_date"`
	TotalAmount  float64          `bson:"total_amount"`
	Status       string           `bson:"status"`
}

func NewOrderDetailDTO(productID bson.ObjectID, quantity int, price float64) *OrderDetailDTO {
	return &OrderDetailDTO{
		ProductID: productID,
		Quantity:  quantity,
		Price:     price,
	}
}

func NewOrderDTO(CustomerID bson.ObjectID, OrderDetails []OrderDetailDTO, OrderDate time.Time, TotalAmount float64, Status string) *OrderDTO {
	return &OrderDTO{
		CustomerID:   CustomerID,
		OrderDetails: OrderDetails,
		OrderDate:    OrderDate,
		TotalAmount:  TotalAmount,
		Status:       Status,
	}
}
