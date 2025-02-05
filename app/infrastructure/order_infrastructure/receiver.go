package order_infrastructure

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type OrderRepository struct {
	Ctx        context.Context
	Collection *mongo.Collection
}

func NewOrderRepository(ctx context.Context, DB *mongo.Database) OrderRepository {
	DBName := os.Getenv("ORDER_COLLECTION_NAME")
	if DBName == "" {
		DBName = "orders"
	}
	collection := DB.Collection(DBName)
	return OrderRepository{Ctx: ctx, Collection: collection}
}
