package customer_infrastructure

import "go.mongodb.org/mongo-driver/v2/bson"

type CustomerDTO struct {
	ID           bson.ObjectID   `bson:"_id,omitempty"`
	Name         string          `bson:"name"`
	Email        string          `bson:"email"`
	Address      string          `bson:"address"`
	OrderHistory []bson.ObjectID `bson:"order_history,omitempty"`
}

func NewCustomerDTO(name string, email string, address string) *CustomerDTO {
	return &CustomerDTO{
		Name:         name,
		Email:        email,
		Address:      address,
		OrderHistory: []bson.ObjectID{},
	}
}
