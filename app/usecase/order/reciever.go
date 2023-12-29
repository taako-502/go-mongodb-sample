package order_usecase

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type OrderService struct {
	Ctx              context.Context
	DB               *mongo.Database
	ConnectionString string
}

func NewOrderService(ctx context.Context, connectionString string, DB *mongo.Database) *OrderService {
	return &OrderService{
		Ctx:              ctx,
		DB:               DB,
		ConnectionString: connectionString,
	}
}
