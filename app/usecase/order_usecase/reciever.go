package order_usecase

import (
	"context"
)

type OrderService struct {
	Ctx              context.Context
	DBName           string
	ConnectionString string
}

func NewOrderService(ctx context.Context, connectionString string, DBName string) *OrderService {
	return &OrderService{
		Ctx:              ctx,
		DBName:           DBName,
		ConnectionString: connectionString,
	}
}
