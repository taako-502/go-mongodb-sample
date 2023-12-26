package model

import (
	customer_infrastructure "go-mongodb-sample/infrastructures/customers"

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
}
