package product_infrastructure

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type Product struct {
	Ctx        context.Context
	Collection *mongo.Collection
}

func NewProduct(ctx context.Context, collection *mongo.Collection) Product {
	return Product{Ctx: ctx, Collection: collection}
}
