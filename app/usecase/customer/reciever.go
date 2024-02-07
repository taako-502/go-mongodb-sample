package customer_usecase

type customerService struct{}

func NewCustomerService() *customerService {
	return &customerService{}
}
