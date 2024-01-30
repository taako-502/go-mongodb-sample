package product_usecase

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductlDTO struct {
	ID                 primitive.ObjectID
	Name               string
	Description        string
	Price              float64
	Stock              int
	Category           string
	PromotionExpiresAt time.Time
}

func NewPromotionProductDTO(
	name string,
	description string,
	price float64,
	stock int,
	category string,
	promotionExpiresAt time.Time,
) *ProductlDTO {
	return &ProductlDTO{
		Name:               name,
		Description:        description,
		Price:              price,
		Stock:              stock,
		Category:           category,
		PromotionExpiresAt: promotionExpiresAt,
	}
}
