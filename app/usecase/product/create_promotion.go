package product_usecase

import (
	product_infrastructure "go-mongodb-sample/app/infrastructures/products"

	"github.com/pkg/errors"
)

func (p ProductService) CreatePromotion(c product_infrastructure.ProductRepository, product *productlDTO) error {
	dto := product_infrastructure.NewPromotionProductDTO(
		product.Name,
		product.Description,
		product.Price,
		product.Stock,
		product.Category,
		product.PromotionExpiresAt,
	)

	if _, err := c.Create(dto); err != nil {
		return errors.Wrap(err, "c.Create")
	}
	return nil
}
