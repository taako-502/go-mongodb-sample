package customer_usecase

import (
	"context"
)

type CustomerService struct {
	Ctx              context.Context
	DBName           string
	ConnectionString string
}

func NewCustomerService(ctx context.Context, connectionString string, DBName string) *CustomerService {
	return &CustomerService{
		Ctx:              ctx,
		DBName:           DBName,
		ConnectionString: connectionString,
	}
}
