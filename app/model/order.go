package model

import (
	"errors"
	"time"

	"github.com/taako-502/go-mongodb-sample/app/infrastructure/order_infrastructure"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type OrderDetail struct {
	ProductID bson.ObjectID
	Quantity  int
	Price     float64
}

type Order struct {
	ID           bson.ObjectID
	CustomerID   bson.ObjectID
	OrderDetails []OrderDetail
	OrderDate    time.Time
	Status       string
}

type OrderAdapter interface {
	Create(dto *order_infrastructure.OrderDTO) (*order_infrastructure.OrderDTO, error)
	FindByCustomerID(customerID bson.ObjectID) ([]order_infrastructure.OrderDTO, error)
	GetTotalAmountSpent(orderHistories []bson.ObjectID) (float64, error)
}

func NewOrder(customerID bson.ObjectID, orderDetails []OrderDetail, orderDate time.Time, status string) (*Order, error) {
	if customerID == bson.NilObjectID {
		return nil, errors.New("customerId is required")
	}
	if len(orderDetails) == 0 {
		return nil, errors.New("orderDetails is required")
	}
	if orderDate.IsZero() {
		return nil, errors.New("orderDate is required")
	}
	if status == "" {
		return nil, errors.New("status is required")
	}
	return &Order{
		CustomerID:   customerID,
		OrderDetails: orderDetails,
		OrderDate:    orderDate,
		Status:       status,
	}, nil
}

func NewOrderDetail(productID bson.ObjectID, quantity int, price float64) *OrderDetail {
	return &OrderDetail{
		ProductID: productID,
		Quantity:  quantity,
		Price:     price,
	}
}
