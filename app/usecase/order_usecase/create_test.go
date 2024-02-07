package order_usecase

import (
	"testing"
	"time"

	"github.com/taako-502/go-mongodb-sample/app/infrastructure/customer_infrastructure/customer_infrastructure_fake"
	"github.com/taako-502/go-mongodb-sample/app/infrastructure/order_infrastructure/order_infrastructure_fake"
	"github.com/taako-502/go-mongodb-sample/app/infrastructure/transaction_manager/transaction_manager_fake"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestOrderService_Create(t *testing.T) {
	exist, _ := primitive.ObjectIDFromHex("000000000000000000000001")
	errorId, _ := primitive.ObjectIDFromHex("000000000000000000000400")
	emptyID, _ := primitive.ObjectIDFromHex("000000000000000000000404")
	historyErrorId, _ := primitive.ObjectIDFromHex("100000000000000000000400")
	type args struct{ dto CreateDTO }
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{
				dto: CreateDTO{
					CustomerID:   exist,
					OrderDetails: []OrderDetailDTO{{ProductID: exist}},
					OrderDate:    time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
					Status:       "test",
				},
			},
			wantErr: false,
		},
		{
			name:    "modelの作成に失敗",
			args:    args{dto: CreateDTO{}},
			wantErr: true,
		},
		{
			name: "カスタマーが存在しない",
			args: args{
				dto: CreateDTO{
					CustomerID:   emptyID,
					OrderDetails: []OrderDetailDTO{{ProductID: emptyID}},
					OrderDate:    time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
					Status:       "test",
				},
			},
			wantErr: true,
		},
		{
			name: "カスタマーの取得処理でエラー",
			args: args{
				dto: CreateDTO{
					CustomerID:   errorId,
					OrderDetails: []OrderDetailDTO{{ProductID: errorId}},
					OrderDate:    time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
					Status:       "test",
				},
			},
			wantErr: true,
		},
		{
			name: "オーダーの作成処理でエラー",
			args: args{
				dto: CreateDTO{
					CustomerID:   exist,
					OrderDetails: []OrderDetailDTO{{ProductID: exist}},
					OrderDate:    time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
					Status:       "error",
				},
			},
			wantErr: true,
		},
		{
			name: "カスタマーのオーダー履歴の作成処理でエラー",
			args: args{
				dto: CreateDTO{
					CustomerID:   historyErrorId,
					OrderDetails: []OrderDetailDTO{{ProductID: errorId}},
					OrderDate:    time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
					Status:       "test",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := orderService{}
			tmFake := transaction_manager_fake.NewFakeTransactionManager()
			customerFake := customer_infrastructure_fake.NewFakeCustomerRepositor()
			orderFake := order_infrastructure_fake.NewFakeOrderRepository()
			if err := o.Create(tmFake, customerFake, orderFake, tt.args.dto); (err != nil) != tt.wantErr {
				t.Errorf("OrderService.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
