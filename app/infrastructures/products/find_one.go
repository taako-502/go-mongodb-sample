package product_infrastructure

import (
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (c ProductRepository) FindOne(id primitive.ObjectID) (*ProductDTO, error) {
	var dto *ProductDTO
	if err := c.Collection.FindOne(c.Ctx, bson.D{{Key: "_id", Value: id}}).Decode(&dto); err != nil {
		return nil, errors.Wrap(err, "c.Collection.FindOne")
	}

	if dto == nil {
		return nil, ErrProductNotFound
	}

	return dto, nil
}
