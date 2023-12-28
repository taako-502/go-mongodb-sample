package product_controller

type ProductController struct {
	ConnectionString string
	DBName           string
	CollectionName   string
}

func NewProductController(con string, DBName string, CollectionName string) ProductController {
	return ProductController{ConnectionString: con, DBName: DBName, CollectionName: CollectionName}
}
