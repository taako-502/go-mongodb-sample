package order_infrastructure

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepository struct {
	Ctx        context.Context
	Collection *mongo.Collection
}

func NewOrderRepository(ctx context.Context, DB *mongo.Database) OrderRepository {
	collection := DB.Collection("orders")
	return OrderRepository{Ctx: ctx, Collection: collection}
}
