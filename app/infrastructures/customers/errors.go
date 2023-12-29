package customer_infrastructure

import "errors"

var ErrCustomerNotFound = errors.New("customer not found")
var ErrCustomerDuplicate = errors.New("customer duplicate")
