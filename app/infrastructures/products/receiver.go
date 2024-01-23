package product_infrastructure

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository struct {
	Ctx        context.Context
	Collection *mongo.Collection
}

func NewProductRepository(ctx context.Context, DB *mongo.Database) ProductRepository {
	DBName := os.Getenv("PRODUCT_COLLECTION_NAME")
	if DBName == "" {
		DBName = "products"
	}
	collection := DB.Collection(DBName)
	return ProductRepository{Ctx: ctx, Collection: collection}
}
