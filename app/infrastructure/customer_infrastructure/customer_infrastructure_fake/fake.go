package customer_infrastructure_fake

import (
	"go-mongodb-sample/app/infrastructure/customer_infrastructure"
	"go-mongodb-sample/app/model"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
		return nil, errors.New("test error")
	}
	return &customer_infrastructure.CustomerDTO{
		Name: dto.Name,
	}, nil
}

func (r *fakeCustomerRepository) FindOne(id primitive.ObjectID) (*customer_infrastructure.CustomerDTO, error) {
	hasOrderHistoryId, _ := primitive.ObjectIDFromHex("000000000000000000000001")
	errorId, _ := primitive.ObjectIDFromHex("000000000000000000000400")
	emptyID, _ := primitive.ObjectIDFromHex("000000000000000000000404")
	hasErrorOrderHistoryId, _ := primitive.ObjectIDFromHex("000000000000000000400001")
	switch id {
	case hasOrderHistoryId:
		return &customer_infrastructure.CustomerDTO{
			ID:   id,
			Name: "test",
			OrderHistory: []primitive.ObjectID{
				primitive.NewObjectID(),
				primitive.NewObjectID(),
			},
		}, nil
	case hasErrorOrderHistoryId:
		return &customer_infrastructure.CustomerDTO{
			ID:   id,
			Name: "test",
			OrderHistory: []primitive.ObjectID{
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

func (r *fakeCustomerRepository) UpdateHistory(ID primitive.ObjectID, orderID primitive.ObjectID) error {
	errorId, _ := primitive.ObjectIDFromHex("100000000000000000000400")
	if ID == errorId {
		return errors.New("test error")
	}
	return nil
}
