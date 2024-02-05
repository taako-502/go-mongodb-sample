package order_usecase

import (
	"go-mongodb-sample/app/infrastructure/customer_infrastructure/customer_infrastructure_fake"
	"go-mongodb-sample/app/infrastructure/order_infrastructure/order_infrastructure_fake"
	"go-mongodb-sample/app/infrastructure/transaction_manager/transaction_manager_fake"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestOrderService_Create(t *testing.T) {
	exist, _ := primitive.ObjectIDFromHex("000000000000000000000001")
	emptyID, _ := primitive.ObjectIDFromHex("000000000000000000000099")
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := OrderService{}
			tmFake := transaction_manager_fake.NewFakeTransactionManager()
			customerFake := customer_infrastructure_fake.NewFakeCustomerRepositor()
			orderFake := order_infrastructure_fake.NewFakeOrderRepository()
			if err := o.Create(tmFake, customerFake, orderFake, tt.args.dto); (err != nil) != tt.wantErr {
				t.Errorf("OrderService.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
