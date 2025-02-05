package customer_infrastructure

import (
	"errors"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (c CustomerRepository) FindOne(id bson.ObjectID) (*CustomerDTO, error) {
	var customer *CustomerDTO
	if err := c.Collection.FindOne(c.Ctx, bson.D{{Key: "_id", Value: id}}).Decode(&customer); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, ErrCustomerNotFound
		} else {
			return nil, err
		}
	}

	return customer, nil
}
