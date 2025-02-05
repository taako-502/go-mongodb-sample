package product_usecase

import (
	"fmt"

	"github.com/taako-502/go-mongodb-sample/app/infrastructure/product_infrastructure"
	"github.com/taako-502/go-mongodb-sample/app/model"
)

func (p productService) CreatePromotion(c model.ProductAdapter, dto *productlDTO) (*productlDTO, error) {
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
		return nil, fmt.Errorf("model.NewProduct: %w", err)
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
		return nil, fmt.Errorf("c.Create: %w", err)
	}

	result := dto
	result.ID = createdDto.ID

	return result, nil
}
