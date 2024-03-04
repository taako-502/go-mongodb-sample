package customer_infrastructure

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

type CustomerRepository struct {
	Ctx        context.Context
	Collection *mongo.Collection
}

func NewCustomerRepository(ctx context.Context, DB *mongo.Database) *CustomerRepository {
	collection := DB.Collection(GetName())
	return &CustomerRepository{Ctx: ctx, Collection: collection}
}

func GetName() string {
	DBName := os.Getenv("CUSTOMER_COLLECTION_NAME")
	if DBName == "" {
		DBName = "customers"
	}

	return DBName
}
