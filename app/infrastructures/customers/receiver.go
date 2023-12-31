package customer_infrastructure

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepository struct {
	Ctx        context.Context
	Collection *mongo.Collection
}

func NewCustomerRepository(ctx context.Context, DB *mongo.Database) OrderRepository {
	collection := DB.Collection("customers")
	return OrderRepository{Ctx: ctx, Collection: collection}
}
