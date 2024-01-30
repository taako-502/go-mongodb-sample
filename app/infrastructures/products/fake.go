package product_infrastructure

import (
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type fakeProductRepository struct {
	insertProduct map[*ProductDTO]error
	findProduct   *ProductDTO
}

func NewFakeProductRepository() ProductRepository {
	productMap := make(map[*ProductDTO]error)
	return &fakeProductRepository{
		insertProduct: productMap,
		findProduct:   nil,
	}
}

func (r *fakeProductRepository) Create(dto *ProductDTO) (*ProductDTO, error) {
	if dto.Name == "error" {
		// エラーパターンのテスト用
		return nil, errors.New("test error")
	}
	if err, exists := r.insertProduct[dto]; exists {
		return nil, errors.Wrap(err, "fakeUserRepository.insertUser")
	}
	_id := primitive.NewObjectID()
	result := &ProductDTO{
		ID:          _id,
		Name:        dto.Name,
		Description: dto.Description,
		Price:       dto.Price,
		Stock:       dto.Stock,
		Category:    dto.Category,
	}
	r.insertProduct[dto] = nil
	return result, nil
}

func (r *fakeProductRepository) FindOne(id primitive.ObjectID) (*ProductDTO, error) {
	product := r.findProduct
	product.ID = id
	return product, nil
}
