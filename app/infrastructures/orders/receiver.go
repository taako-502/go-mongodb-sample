package order_infrastructure

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepository struct {
	Ctx        context.Context
	Collection *mongo.Collection
}

func NewOrderRepository(ctx context.Context, collection *mongo.Collection) OrderRepository {
	return OrderRepository{Ctx: ctx, Collection: collection}
}
