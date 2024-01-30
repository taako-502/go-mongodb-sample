package product_usecase

import (
	"context"
)

type ProductService struct {
	Ctx              context.Context
	DBName           string
	ConnectionString string
}

func NewProductService(ctx context.Context, connectionString string, DBName string) *ProductService {
	return &ProductService{
		Ctx:              ctx,
		DBName:           DBName,
		ConnectionString: connectionString,
	}
}
