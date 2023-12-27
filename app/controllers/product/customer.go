package customer_controller

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type ProductController struct {
	Ctx        context.Context
	Collection *mongo.Collection
}

func NewProductController(ctx context.Context, collection *mongo.Collection) ProductController {
	return ProductController{Ctx: ctx, Collection: collection}
}
