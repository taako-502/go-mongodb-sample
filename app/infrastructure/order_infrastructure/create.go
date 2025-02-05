package order_infrastructure

import (
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func (c OrderRepository) Create(dto *OrderDTO) (*OrderDTO, error) {
	result, err := c.Collection.InsertOne(c.Ctx, dto)
	if err != nil {
		return nil, fmt.Errorf("c.Collection.InsertOne: %w", err)
	}

	return &OrderDTO{
		ID:           result.InsertedID.(bson.ObjectID),
		CustomerID:   dto.CustomerID,
		OrderDetails: dto.OrderDetails,
		OrderDate:    dto.OrderDate,
		TotalAmount:  dto.TotalAmount,
		Status:       dto.Status,
	}, nil
}
