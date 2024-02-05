package product_usecase

import (
	"go-mongodb-sample/app/infrastructure/product_infrastructure"
	"go-mongodb-sample/app/model"

	"github.com/pkg/errors"
)

func (p ProductService) CreatePromotion(c model.ProductAdapter, dto *productlDTO) (*productlDTO, error) {
	model, err := model.NewProduct(
		dto.ID,
		dto.Name,
		dto.Description,
		dto.Price,
		dto.Stock,
		dto.Category,
		dto.PromotionExpiresAt,
	)
	if err != nil {
		return nil, errors.Wrap(err, "model.NewProduct")
	}

	product := product_infrastructure.NewPromotionProductDTO(
		model.Name,
		model.Description,
		model.Price,
		model.Stock,
		model.Category,
		model.PromotionExpiresAt,
	)

	createdDto, err := c.Create(product)
	if err != nil {
		return nil, errors.Wrap(err, "c.Create")
	}

	result := dto
	result.ID = createdDto.ID

	return result, nil
}
