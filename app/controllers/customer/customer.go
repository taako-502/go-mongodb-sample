package customer_controller

type CostumerController struct {
	ConnectionString string
	DBName           string
	CollectionName   string
}

func NewCostumerController(con string, DBName string, CollectionName string) CostumerController {
	return CostumerController{ConnectionString: con, DBName: DBName, CollectionName: CollectionName}
}
