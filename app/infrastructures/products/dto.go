package product_infrastructure

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductDTO struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty"`
	Name               string             `bson:"name"`
	Description        string             `bson:"description"`
	Price              float64            `bson:"price"`
	Stock              int                `bson:"stock"`
	Category           string             `bson:"category"`
	PromotionExpiresAt *time.Time         `bson:"promotion_expires_at"`
}

func NewProductDTO(name string, description string, price float64, stock int, category string) *ProductDTO {
	return &ProductDTO{
		Name:               name,
		Description:        description,
		Price:              price,
		Stock:              stock,
		Category:           category,
		PromotionExpiresAt: nil, // nilを入れることで、TTLインデックスが無効になる
	}
}

func NewPromotionProductDTO(name string, description string, price float64, stock int, category string, promotionExpiresAt time.Time) *ProductDTO {
	return &ProductDTO{
		Name:               name,
		Description:        description,
		Price:              price,
		Stock:              stock,
		Category:           category,
		PromotionExpiresAt: &promotionExpiresAt,
	}
}
