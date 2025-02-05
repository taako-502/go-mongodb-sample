package customer_infrastructure

import (
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
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
			return nil, fmt.Errorf("c.Collection.InsertOne: %w", err)
		}
	}

	return &CustomerDTO{
		ID:           result.InsertedID.(bson.ObjectID),
		Name:         dto.Name,
		Email:        dto.Email,
		Address:      dto.Address,
		OrderHistory: dto.OrderHistory,
	}, nil
}
