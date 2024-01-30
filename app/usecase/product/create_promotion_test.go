package product_usecase

import (
	product_infrastructure_fake "go-mongodb-sample/app/infrastructures/products/fake"
	"reflect"
	"testing"
	"time"
)

func TestProductService_CreatePromotion(t *testing.T) {
	promotionDate := time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local)

	type args struct {
		dto *productlDTO
	}
	tests := []struct {
		name    string
		args    args
		want    *productlDTO
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{
				dto: &productlDTO{
					Name:               "test",
					Description:        "test",
					Price:              100,
					Stock:              100,
					Category:           "test",
					PromotionExpiresAt: &promotionDate,
				},
			},
			want: &productlDTO{
				Name:               "test",
				Description:        "test",
				Price:              100,
				Stock:              100,
				Category:           "test",
				PromotionExpiresAt: &promotionDate,
			},
			wantErr: false,
		}, {
			name: "正常系", args: args{dto: &productlDTO{Name: "error"}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fake := product_infrastructure_fake.NewFakeProductRepository()
			p := ProductService{}
			got, err := p.CreatePromotion(fake, tt.args.dto)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProductService.CreatePromotion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductService.CreatePromotion() = %v, want %v", got, tt.want)
			}
		})
	}
}