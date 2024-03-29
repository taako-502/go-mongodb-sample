package model

import (
	"github.com/taako-502/go-mongodb-sample/app/infrastructure/customer_infrastructure"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Customer struct {
	ID           primitive.ObjectID
	Name         string
	Email        string
	Address      string
	OrderHistory []primitive.ObjectID
}

type CustomerAdapter interface {
	Create(dto *customer_infrastructure.CustomerDTO) (*customer_infrastructure.CustomerDTO, error)
	FindOne(id primitive.ObjectID) (*customer_infrastructure.CustomerDTO, error)
	UpdateHistory(ID primitive.ObjectID, orderID primitive.ObjectID) error
}
