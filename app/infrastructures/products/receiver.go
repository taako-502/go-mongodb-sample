package product_infrastructure

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

type ProductReciever struct {
	Ctx        context.Context
	Collection *mongo.Collection
}

func NewProductRepository(ctx context.Context, DB *mongo.Database) *ProductReciever {
	DBName := os.Getenv("PRODUCT_COLLECTION_NAME")
	if DBName == "" {
		DBName = "products"
	}
	collection := DB.Collection(DBName)
	return &ProductReciever{Ctx: ctx, Collection: collection}
}
