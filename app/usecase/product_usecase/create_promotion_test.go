package product_usecase

import (
	"go-mongodb-sample/app/infrastructure/product_infrastructure/product_infrastructure_fake"
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
			name: "バリデーションエラー：名前が空白", args: args{dto: &productlDTO{Name: "", Price: 100, Stock: 100, Category: "error"}},
			wantErr: true,
		}, {
			name: "MongoDBによるエラーが発生", args: args{dto: &productlDTO{Name: "error", Price: 100, Stock: 100, Category: "error"}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fake := product_infrastructure_fake.NewFakeProductRepository()
			p := productService{}
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
