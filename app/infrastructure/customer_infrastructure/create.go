package customer_infrastructure

import (
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (c CustomerRepository) Create(dto *CustomerDTO) (*CustomerDTO, error) {
	result, err := c.Collection.InsertOne(c.Ctx, dto)
	if err != nil {
		if writeErr, ok := err.(mongo.WriteException); ok {
			for _, e := range writeErr.WriteErrors {
				if e.Code == 11000 {
					return nil, ErrCustomerDuplicate
				}
			}
		} else {
			return nil, errors.Wrap(err, "c.Collection.InsertOne")
		}
	}

	return &CustomerDTO{
		ID:           result.InsertedID.(primitive.ObjectID),
		Name:         dto.Name,
		Email:        dto.Email,
		Address:      dto.Address,
		OrderHistory: dto.OrderHistory,
	}, nil
}
