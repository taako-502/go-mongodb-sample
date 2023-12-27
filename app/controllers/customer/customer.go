package customer_controller

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type CostumerController struct {
	Ctx        context.Context
	Collection *mongo.Collection
}

func NewCostumerController(ctx context.Context, collection *mongo.Collection) CostumerController {
	return CostumerController{Ctx: ctx, Collection: collection}
}
