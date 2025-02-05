package order_infrastructure_fake

import (
	"errors"

	"github.com/taako-502/go-mongodb-sample/app/infrastructure/customer_infrastructure"
	"github.com/taako-502/go-mongodb-sample/app/infrastructure/order_infrastructure"
	"github.com/taako-502/go-mongodb-sample/app/model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type fakeOrderRepository struct {
	insertOrder           map[*order_infrastructure.OrderDTO]error
	finOrderdByCustomerID *order_infrastructure.OrderDTO
	getTotalAmountSpent   map[float64]error
}

func NewFakeOrderRepository() model.OrderAdapter {
	orderMap := make(map[*order_infrastructure.OrderDTO]error)
	getTotalAmountSpentMap := make(map[float64]error)
	return &fakeOrderRepository{
		insertOrder:           orderMap,
		finOrderdByCustomerID: nil,
		getTotalAmountSpent:   getTotalAmountSpentMap,
	}
}

func (r *fakeOrderRepository) Create(dto *order_infrastructure.OrderDTO) (*order_infrastructure.OrderDTO, error) {
	if dto.Status == "error" {
		// エラーパターンのテスト用
		return nil, errors.New("test error")
	}
	return &order_infrastructure.OrderDTO{}, nil
}

func (r *fakeOrderRepository) FindByCustomerID(id bson.ObjectID) ([]order_infrastructure.OrderDTO, error) {
	emptyID, _ := bson.ObjectIDFromHex("000000000000000000000000")
	if id == emptyID {
		return nil, customer_infrastructure.ErrCustomerNotFound
	}
	var customers []order_infrastructure.OrderDTO
	return customers, nil
}

func (r *fakeOrderRepository) GetTotalAmountSpent(orderHistories []bson.ObjectID) (float64, error) {
	errorId, _ := bson.ObjectIDFromHex("000000000000000000000400")
	if orderHistories[0] == errorId {
		return 0, errors.New("test error")
	}
	return 0, nil
}
