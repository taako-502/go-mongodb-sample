package order_infrastructure

import (
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func (c OrderRepository) FindByCustomerID(customerID bson.ObjectID) ([]OrderDTO, error) {
	var orders []OrderDTO

	cursor, err := c.Collection.Find(c.Ctx, bson.D{{Key: "customer_id", Value: customerID}})
	if err != nil {
		return nil, fmt.Errorf("c.Collection.Find: %w", err)
	}
	defer cursor.Close(c.Ctx)

	for cursor.Next(c.Ctx) {
		var order OrderDTO
		if err := cursor.Decode(&order); err != nil {
			return nil, fmt.Errorf("cursor.Decode: %w", err)
		}
		orders = append(orders, order)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor.Err: %w", err)
	}

	if len(orders) == 0 {
		return nil, ErrOrderNotFound
	}

	return orders, nil
}
