package product_controller

import "github.com/labstack/echo/v4"

type ProductController struct {
	ConnectionString string
	DBName           string
	CollectionName   string
}

func NewProductController(con string, DBName string, CollectionName string) ProductController {
	return ProductController{ConnectionString: con, DBName: DBName, CollectionName: CollectionName}
}

func (p ProductController) Connect(c echo.Context) error {
	return nil
}
