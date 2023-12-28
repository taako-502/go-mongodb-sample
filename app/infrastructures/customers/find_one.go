package customer_infrastructure

import "go.mongodb.org/mongo-driver/bson/primitive"

func (c Customer) FindOne(id primitive.ObjectID) (*CustomerDTO, error) {
	var customer *CustomerDTO
	if err := c.Collection.FindOne(c.Ctx, primitive.D{{Key: "_id", Value: id}}).Decode(&customer); err != nil {
		return nil, err
	}
	return customer, nil
}
