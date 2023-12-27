package order_controller

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type OrderController struct {
	Ctx        context.Context
	Collection *mongo.Collection
}

func NewOrderController(ctx context.Context, collection *mongo.Collection) OrderController {
	return OrderController{Ctx: ctx, Collection: collection}
}
