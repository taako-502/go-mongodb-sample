package model

import (
	"errors"
	product_infrastructure "go-mongodb-sample/app/infrastructure/products"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID                 primitive.ObjectID
	Name               string
	Description        string
	Price              float64
	Stock              int
	Category           string
	PromotionExpiresAt *time.Time
}

type ProductAdapter interface {
	Create(dto *product_infrastructure.ProductDTO) (*product_infrastructure.ProductDTO, error)
	FindOne(id primitive.ObjectID) (*product_infrastructure.ProductDTO, error)
}

func NewProduct(
	id primitive.ObjectID,
	name string,
	description string,
	price float64,
	stock int,
	category string,
	promotionExpiresAt *time.Time,
) (*Product, error) {
	model := &Product{
		ID:                 id,
		Name:               name,
		Description:        description,
		Price:              price,
		Stock:              stock,
		Category:           category,
		PromotionExpiresAt: promotionExpiresAt,
	}

	if err := model.validate(); err != nil {
		return nil, err
	}

	return model, nil
}

func (p Product) validate() error {
	if p.Name == "" {
		return errors.New("name is required")
	}
	if p.Price <= 0 {
		return errors.New("price is required")
	}
	if p.Stock <= 0 {
		return errors.New("stock is required")
	}
	if p.Category == "" {
		return errors.New("category is required")
	}
	return nil
}
