package customer_infrastructure

import (
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (c Customer) Create(dto *CustomerDTO) (*CustomerDTO, error) {
	result, err := c.Collection.InsertOne(c.Ctx, dto)
	if err != nil {
		return nil, errors.Wrap(err, "c.Collection.InsertOne")
	}

	return &CustomerDTO{
		ID:           result.InsertedID.(primitive.ObjectID),
		Name:         dto.Name,
		Email:        dto.Email,
		Address:      dto.Address,
		OrderHistory: dto.OrderHistory,
	}, nil
}
