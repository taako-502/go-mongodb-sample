package product_infrastructure_fake

import (
	"errors"
	"fmt"

	"github.com/taako-502/go-mongodb-sample/app/infrastructure/product_infrastructure"
	"github.com/taako-502/go-mongodb-sample/app/model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type fakeProductRepository struct {
	insertProduct map[*product_infrastructure.ProductDTO]error
	findProduct   *product_infrastructure.ProductDTO
}

func NewFakeProductRepository() model.ProductAdapter {
	productMap := make(map[*product_infrastructure.ProductDTO]error)
	return &fakeProductRepository{
		insertProduct: productMap,
		findProduct:   nil,
	}
}

func (r *fakeProductRepository) Create(dto *product_infrastructure.ProductDTO) (*product_infrastructure.ProductDTO, error) {
	if dto.Name == "error" {
		// エラーパターンのテスト用
		return nil, errors.New("test error")
	}
	if err, exists := r.insertProduct[dto]; exists {
		return nil, fmt.Errorf("fakeUserRepository.insertUser: %w", err)
	}
	result := &product_infrastructure.ProductDTO{
		Name:        dto.Name,
		Description: dto.Description,
		Price:       dto.Price,
		Stock:       dto.Stock,
		Category:    dto.Category,
	}
	r.insertProduct[dto] = nil
	return result, nil
}

func (r *fakeProductRepository) FindOne(id bson.ObjectID) (*product_infrastructure.ProductDTO, error) {
	product := r.findProduct
	product.ID = id
	return product, nil
}
