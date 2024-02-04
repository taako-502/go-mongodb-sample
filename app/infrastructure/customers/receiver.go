package customer_infrastructure

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepository struct {
	Ctx        context.Context
	Collection *mongo.Collection
}

func NewCustomerRepository(ctx context.Context, DB *mongo.Database) *OrderRepository {
	DBName := os.Getenv("CUSTOMER_COLLECTION_NAME")
	if DBName == "" {
		DBName = "customers"
	}
	collection := DB.Collection(DBName)
	return &OrderRepository{Ctx: ctx, Collection: collection}
}
