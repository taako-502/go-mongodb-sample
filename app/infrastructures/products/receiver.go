package product_infrastructure

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type Producter struct {
	Ctx        context.Context
	Collection *mongo.Collection
}

func NewProduct(ctx context.Context, collection *mongo.Collection) Producter {
	return Producter{Ctx: ctx, Collection: collection}
}
