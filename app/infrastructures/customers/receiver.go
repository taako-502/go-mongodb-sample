package customer_infrastructure

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type Customer struct {
	Ctx        context.Context
	Collection *mongo.Collection
}

func NewCustomerRepository(ctx context.Context, collection *mongo.Collection) Customer {
	return Customer{Ctx: ctx, Collection: collection}
}
