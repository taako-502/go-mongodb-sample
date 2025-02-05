package product_infrastructure

import (
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func (c *ProductReciever) Create(dto *ProductDTO) (*ProductDTO, error) {
	result, err := c.Collection.InsertOne(c.Ctx, dto)
	if err != nil {
		return nil, fmt.Errorf("c.Collection.InsertOne: %w", err)
	}

	return &ProductDTO{
		ID:          result.InsertedID.(bson.ObjectID),
		Name:        dto.Name,
		Description: dto.Description,
		Price:       dto.Price,
		Stock:       dto.Stock,
		Category:    dto.Category,
	}, nil
}
