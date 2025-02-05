package customer_usecase

import (
	"testing"

	"github.com/taako-502/go-mongodb-sample/app/infrastructure/customer_infrastructure/customer_infrastructure_fake"
	"github.com/taako-502/go-mongodb-sample/app/infrastructure/order_infrastructure/order_infrastructure_fake"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func Test_customerService_GetTotalAmountSpent(t *testing.T) {
	hasOrderHistoryId, _ := bson.ObjectIDFromHex("000000000000000000000001")
	hasErrorOrderHistoryId, _ := bson.ObjectIDFromHex("000000000000000000400001")
	custmoerEmptyID, _ := bson.ObjectIDFromHex("000000000000000000000404")
	type args struct {
		ID bson.ObjectID
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
			args:    args{ID: bson.NewObjectID()},
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
