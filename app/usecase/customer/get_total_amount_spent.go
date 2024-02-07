package customer_usecase

import (
	"github.com/taako-502/go-mongodb-sample/app/model"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (c customerService) GetTotalAmountSpent(ci model.CustomerAdapter, or model.OrderAdapter, ID primitive.ObjectID) (float64, error) {
	customer, err := ci.FindOne(ID)
	if err != nil {
		return 0, errors.Wrap(err, "customer_infrastructure.CustomerRepository.Find")
	}
	if customer.OrderHistory == nil {
		return 0, nil
	}

	amount, err := or.GetTotalAmountSpent(customer.OrderHistory)
	if err != nil {
		return 0, errors.Wrap(err, "order_infrastructure.OrderRepository.GetTotalAmountSpent")
	}

	return amount, nil
}
