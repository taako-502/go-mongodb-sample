package customer_infrastructure_fake

import (
	customer_infrastructure "go-mongodb-sample/app/infrastructure/customers"
	"go-mongodb-sample/app/model"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type fakeCustomerRepository struct {
	insertCustomer        map[*customer_infrastructure.CustomerDTO]error
	findCustomer          *customer_infrastructure.CustomerDTO
	updateCustomerHistory error
}

func NewFakeProductRepository() model.CustomerAdapter {
	customerMap := make(map[*customer_infrastructure.CustomerDTO]error)
	return &fakeCustomerRepository{
		insertCustomer:        customerMap,
		findCustomer:          nil,
		updateCustomerHistory: nil,
	}
}

func (r *fakeCustomerRepository) Create(dto *customer_infrastructure.CustomerDTO) (*customer_infrastructure.CustomerDTO, error) {
	if dto.Name == "error" {
		// エラーパターンのテスト用
		return nil, errors.New("test error")
	}
	return &customer_infrastructure.CustomerDTO{
		Name: dto.Name,
	}, nil
}

func (r *fakeCustomerRepository) FindOne(id primitive.ObjectID) (*customer_infrastructure.CustomerDTO, error) {
	emptyID, _ := primitive.ObjectIDFromHex("000000000000000000000000")
	if id == emptyID {
		return nil, customer_infrastructure.ErrCustomerNotFound
	}
	var customer *customer_infrastructure.CustomerDTO
	customer.ID = id
	return customer, nil
}

func (r *fakeCustomerRepository) UpdateHistory(ID primitive.ObjectID, orderID primitive.ObjectID) error {
	return nil
}
