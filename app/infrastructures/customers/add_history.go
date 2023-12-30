package customer_infrastructure

import (
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (c Customer) UpdateHistory(ID primitive.ObjectID, orderID primitive.ObjectID) error {
	update := bson.D{{Key: "$push", Value: bson.D{{Key: "order_history", Value: orderID}}}}
	result, err := c.Collection.UpdateByID(c.Ctx, ID, update)

	if err != nil {
		return errors.Wrap(err, "c.Collection.UpdateByID")
	}

	if result.MatchedCount == 0 {
		return ErrCustomerNotFound
	}

	return nil
}
