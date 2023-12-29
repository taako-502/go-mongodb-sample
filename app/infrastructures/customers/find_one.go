package customer_infrastructure

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (c Customer) Find(id primitive.ObjectID) (*CustomerDTO, error) {
	var customer *CustomerDTO
	if err := c.Collection.FindOne(c.Ctx, primitive.D{{Key: "_id", Value: id}}).Decode(&customer); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, ErrCustomerNotFound
		} else {
			return nil, err
		}
	}

	return customer, nil
}
