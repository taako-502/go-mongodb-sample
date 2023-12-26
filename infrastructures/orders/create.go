package order_infrastructure

import (
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (c OrderRepository) Create(dto *OrderDTO) (*OrderDTO, error) {
	result, err := c.Collection.InsertOne(c.Ctx, dto)
	if err != nil {
		return nil, errors.Wrap(err, "c.Collection.InsertOne")
	}

	return &OrderDTO{
		ID:           result.InsertedID.(primitive.ObjectID),
		CustomerID:   dto.CustomerID,
		OrderDetails: dto.OrderDetails,
		OrderDate:    dto.OrderDate,
		TotalAmount:  dto.TotalAmount,
		Status:       dto.Status,
	}, nil
}
