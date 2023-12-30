package customer_infrastructure

import "go.mongodb.org/mongo-driver/bson/primitive"

type CustomerDTO struct {
	ID           primitive.ObjectID   `bson:"_id,omitempty"`
	Name         string               `bson:"name"`
	Email        string               `bson:"email"`
	Address      string               `bson:"address"`
	OrderHistory []primitive.ObjectID `bson:"order_history,omitempty"`
}

func NewCustomerDTO(name string, email string, address string) *CustomerDTO {
	return &CustomerDTO{
		Name:         name,
		Email:        email,
		Address:      address,
		OrderHistory: []primitive.ObjectID{},
	}
}
