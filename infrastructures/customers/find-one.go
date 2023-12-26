package customer_infrastructure

import "go.mongodb.org/mongo-driver/bson/primitive"

func (c Customer) FindOne(id primitive.ObjectID) (*CustomerDTO, error) {
	var customer *CustomerDTO
	err := c.Collection.FindOne(c.Ctx, primitive.D{{Key: "_id", Value: id}}).Decode(&customer)
	if err != nil {
		return nil, err
	}
	return customer, nil
}
