package customer_infrastructure

import (
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func (c CustomerRepository) UpdateHistory(ID bson.ObjectID, orderID bson.ObjectID) error {
	update := bson.D{{Key: "$push", Value: bson.D{{Key: "order_history", Value: orderID}}}}
	result, err := c.Collection.UpdateByID(c.Ctx, ID, update)

	if err != nil {
		return fmt.Errorf("c.Collection.UpdateByID: %w", err)
	}

	if result.MatchedCount == 0 {
		return ErrCustomerNotFound
	}

	return nil
}
