package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	product_infrastructure "go-mongodb-sample/app/infrastructures/products"
)

type Product struct {
	ID          primitive.ObjectID
	Name        string
	Description string
	Price       float64
	Stock       int
	Category    string
}

type Producter interface {
	Create(dto *product_infrastructure.ProductDTO) (*product_infrastructure.ProductDTO, error)
}
