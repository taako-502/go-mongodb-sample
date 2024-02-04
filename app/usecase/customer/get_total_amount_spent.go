package customer_usecase

import (
	customer_infrastructure "go-mongodb-sample/app/infrastructure/customers"
	order_infrastructure "go-mongodb-sample/app/infrastructure/orders"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (c CustomerService) GetTotalAmountSpent(ID primitive.ObjectID) (float64, error) {
	client, err := mongo.Connect(c.Ctx, options.Client().ApplyURI(c.ConnectionString))
	if err != nil {
		return 0, errors.Wrap(err, "mongo.Connect")
	}
	defer client.Disconnect(c.Ctx)

	cr := customer_infrastructure.NewCustomerRepository(c.Ctx, client.Database(c.DBName))
	customer, err := cr.Find(ID)
	if err != nil {
		errors.Wrap(err, "cr.Find")
	}
	if customer.OrderHistory == nil {
		return 0, nil
	}

	or := order_infrastructure.NewOrderRepository(c.Ctx, client.Database(c.DBName))
	amount, err := or.GetTotalAmountSpent(customer.OrderHistory)
	if err != nil {
		return 0, err
	}

	return amount, nil
}
