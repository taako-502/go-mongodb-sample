package order_infrastructure

import (
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (c OrderRepository) FindByCustomerID(customerID primitive.ObjectID) ([]OrderDTO, error) {
	var orders []OrderDTO

	cursor, err := c.Collection.Find(c.Ctx, bson.D{{Key: "customer_id", Value: customerID}})
	if err != nil {
		return nil, errors.Wrap(err, "c.Collection.Find")
	}
	defer cursor.Close(c.Ctx)

	for cursor.Next(c.Ctx) {
		var order OrderDTO
		if err := cursor.Decode(&order); err != nil {
			return nil, errors.Wrap(err, "cursor.Decode")
		}
		orders = append(orders, order)
	}

	if err := cursor.Err(); err != nil {
		return nil, errors.Wrap(err, "cursor.Err")
	}

	if len(orders) == 0 {
		return nil, ErrOrderNotFound
	}

	return orders, nil
}
