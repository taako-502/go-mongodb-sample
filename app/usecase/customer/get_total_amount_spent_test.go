package customer_usecase

import (
	"go-mongodb-sample/app/infrastructure/customer_infrastructure/customer_infrastructure_fake"
	"go-mongodb-sample/app/infrastructure/order_infrastructure/order_infrastructure_fake"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Test_customerService_GetTotalAmountSpent(t *testing.T) {
	hasOrderHistoryId, _ := primitive.ObjectIDFromHex("000000000000000000000001")
	hasErrorOrderHistoryId, _ := primitive.ObjectIDFromHex("000000000000000000400001")
	custmoerEmptyID, _ := primitive.ObjectIDFromHex("000000000000000000000404")
	type args struct {
		ID primitive.ObjectID
	}
	tests := []struct {
		name    string
		c       customerService
		args    args
		want    float64
		wantErr bool
	}{
		{
			name:    "正常系",
			args:    args{ID: hasOrderHistoryId},
			want:    0,
			wantErr: false,
		},
		{
			name:    "カスタマーが見つからない",
			args:    args{ID: custmoerEmptyID},
			want:    0,
			wantErr: true,
		},
		{
			name:    "カスタマーが履歴を持っていない",
			args:    args{ID: primitive.NewObjectID()},
			want:    0,
			wantErr: false,
		},
		{
			name:    "合計金額を取得しようとしてエラー",
			args:    args{ID: hasErrorOrderHistoryId},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := customerService{}
			customerFake := customer_infrastructure_fake.NewFakeCustomerRepositor()
			orderFake := order_infrastructure_fake.NewFakeOrderRepository()
			got, err := c.GetTotalAmountSpent(customerFake, orderFake, tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("customerService.GetTotalAmountSpent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("customerService.GetTotalAmountSpent() = %v, want %v", got, tt.want)
			}
		})
	}
}
