package customer_infrastructure

import (
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (c Customer) UpdateHistory(ID primitive.ObjectID, orderID primitive.ObjectID) error {
	_, err := c.Collection.UpdateByID(
		c.Ctx,
		bson.D{{Key: "_id", Value: ID}},
		bson.D{{Key: "$push", Value: bson.D{{Key: "order_history", Value: orderID}}}},
	)
	if err != nil {
		return errors.Wrap(err, "c.Collection.UpdateByID")
	}
	return nil
}
