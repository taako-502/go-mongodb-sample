package order_controller

type OrderController struct {
	ConnectionString string
	DBName           string
	CollectionName   string
}

func NewOrderController(con string, DBName string, CollectionName string) OrderController {
	return OrderController{ConnectionString: con, DBName: DBName, CollectionName: CollectionName}
}
