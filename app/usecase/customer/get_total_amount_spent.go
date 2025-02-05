package customer_usecase

import (
	"fmt"

	"github.com/taako-502/go-mongodb-sample/app/model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (c customerService) GetTotalAmountSpent(ci model.CustomerAdapter, or model.OrderAdapter, ID bson.ObjectID) (float64, error) {
	customer, err := ci.FindOne(ID)
	if err != nil {
		return 0, fmt.Errorf("customer_infrastructure.CustomerRepository.Find: %w", err)
	}
	if customer.OrderHistory == nil {
		return 0, nil
	}

	amount, err := or.GetTotalAmountSpent(customer.OrderHistory)
	if err != nil {
		return 0, fmt.Errorf("order_infrastructure.OrderRepository.GetTotalAmountSpent: %w", err)
	}

	return amount, nil
}
