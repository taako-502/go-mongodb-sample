package product_infrastructure

import (
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (c Producter) Create(dto *ProductDTO) (*ProductDTO, error) {
	result, err := c.Collection.InsertOne(c.Ctx, dto)
	if err != nil {
		return nil, errors.Wrap(err, "c.Collection.InsertOne")
	}

	return &ProductDTO{
		ID:          result.InsertedID.(primitive.ObjectID),
		Name:        dto.Name,
		Description: dto.Description,
		Price:       dto.Price,
		Stock:       dto.Stock,
		Category:    dto.Category,
	}, nil
}
