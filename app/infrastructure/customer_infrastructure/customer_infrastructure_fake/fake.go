package customer_infrastructure_fake

import (
	"errors"
	"fmt"

	"github.com/taako-502/go-mongodb-sample/app/infrastructure/customer_infrastructure"
	"github.com/taako-502/go-mongodb-sample/app/model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type fakeCustomerRepository struct {
	insertCustomer        map[*customer_infrastructure.CustomerDTO]error
	findCustomer          *customer_infrastructure.CustomerDTO
	updateCustomerHistory error
}

func NewFakeCustomerRepositor() model.CustomerAdapter {
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
		return nil, fmt.Errorf("test error")
	}
	return &customer_infrastructure.CustomerDTO{
		Name: dto.Name,
	}, nil
}

func (r *fakeCustomerRepository) FindOne(id bson.ObjectID) (*customer_infrastructure.CustomerDTO, error) {
	hasOrderHistoryId, _ := bson.ObjectIDFromHex("000000000000000000000001")
	errorId, _ := bson.ObjectIDFromHex("000000000000000000000400")
	emptyID, _ := bson.ObjectIDFromHex("000000000000000000000404")
	hasErrorOrderHistoryId, _ := bson.ObjectIDFromHex("000000000000000000400001")
	switch id {
	case hasOrderHistoryId:
		return &customer_infrastructure.CustomerDTO{
			ID:   id,
			Name: "test",
			OrderHistory: []bson.ObjectID{
				bson.NewObjectID(),
				bson.NewObjectID(),
			},
		}, nil
	case hasErrorOrderHistoryId:
		return &customer_infrastructure.CustomerDTO{
			ID:   id,
			Name: "test",
			OrderHistory: []bson.ObjectID{
				errorId,
			},
		}, nil
	case errorId:
		return nil, errors.New("test error")
	case emptyID:
		return nil, customer_infrastructure.ErrCustomerNotFound
	}
	return &customer_infrastructure.CustomerDTO{
		ID:           id,
		Name:         "test",
		OrderHistory: nil,
	}, nil
}

func (r *fakeCustomerRepository) UpdateHistory(ID bson.ObjectID, orderID bson.ObjectID) error {
	errorId, _ := bson.ObjectIDFromHex("100000000000000000000400")
	if ID == errorId {
		return errors.New("test error")
	}
	return nil
}
