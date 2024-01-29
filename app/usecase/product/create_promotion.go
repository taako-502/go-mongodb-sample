package product_usecase

import (
	product_infrastructure "go-mongodb-sample/app/infrastructures/products"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (p ProductService) CreatePromotion(product *ProductlDTO) error {
	client, err := mongo.Connect(p.Ctx, options.Client().ApplyURI(p.ConnectionString))
	if err != nil {
		return errors.Wrap(err, "mongo.Connect")
	}
	defer client.Disconnect(p.Ctx)

	c := product_infrastructure.NewProductRepository(p.Ctx, client.Database(p.DBName))
	dto := product_infrastructure.NewPromotionProductDTO(
		product.Name,
		product.Description,
		product.Price,
		product.Stock,
		product.Category,
		product.PromotionExpiresAt,
	)

	_, err = c.Create(dto)
	if err != nil {
		return errors.Wrap(err, "c.Create")
	}
	return nil
}
