package product_infrastructure

import (
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func (c *ProductReciever) FindOne(id bson.ObjectID) (*ProductDTO, error) {
	var dto *ProductDTO
	if err := c.Collection.FindOne(c.Ctx, bson.D{{Key: "_id", Value: id}}).Decode(&dto); err != nil {
		return nil, fmt.Errorf("c.Collection.FindOne: %w", err)
	}

	if dto == nil {
		return nil, ErrProductNotFound
	}

	return dto, nil
}
