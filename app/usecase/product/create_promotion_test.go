package product_usecase

import (
	product_infrastructure "go-mongodb-sample/app/infrastructures/products"
	"testing"
	"time"
)

func TestProductService_CreatePromotion(t *testing.T) {
	type args struct {
		product *productlDTO
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{
				product: &productlDTO{
					Name:               "test",
					Description:        "test",
					Price:              100,
					Stock:              100,
					Category:           "test",
					PromotionExpiresAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
				},
			},
			wantErr: false,
		},
		{
			name: "正常系", args: args{product: &productlDTO{Name: "error"}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fake := product_infrastructure.NewFakeProductRepository()
			p := NewProductService()
			if err := p.CreatePromotion(fake, tt.args.product); (err != nil) != tt.wantErr {
				t.Errorf("ProductService.CreatePromotion() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
