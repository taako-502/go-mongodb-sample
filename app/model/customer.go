package model

import (
	"github.com/taako-502/go-mongodb-sample/app/infrastructure/customer_infrastructure"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Customer struct {
	ID           bson.ObjectID
	Name         string
	Email        string
	Address      string
	OrderHistory []bson.ObjectID
}

type CustomerAdapter interface {
	Create(dto *customer_infrastructure.CustomerDTO) (*customer_infrastructure.CustomerDTO, error)
	FindOne(id bson.ObjectID) (*customer_infrastructure.CustomerDTO, error)
	UpdateHistory(ID bson.ObjectID, orderID bson.ObjectID) error
}
