package order_usecase

import (
	"context"
	"go-mongodb-sample/app/infrastructure"
	"go-mongodb-sample/app/infrastructure/customers/customer_infrastructure_fake"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestOrderService_Create(t *testing.T) {
	emptyID, _ := primitive.ObjectIDFromHex("000000000000000000000001")
	type fields struct {
		Ctx              context.Context
		DBName           string
		ConnectionString string
	}
	type args struct {
		tm  *infrastructure.MongoTransactionManager
		dto CreateDTO
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "modelの作成に失敗",
			fields: fields{
				Ctx:              context.Background(),
				DBName:           "test",
				ConnectionString: "mongodb://localhost:27017",
			},
			args: args{
				tm:  &infrastructure.MongoTransactionManager{},
				dto: CreateDTO{},
			},
			wantErr: true,
		},
		{
			name: "カスタマーが存在しない",
			fields: fields{
				Ctx:              context.Background(),
				DBName:           "test",
				ConnectionString: "mongodb://localhost:27017",
			},
			args: args{
				tm: &infrastructure.MongoTransactionManager{},
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
			o := OrderService{
				Ctx:              tt.fields.Ctx,
				DBName:           tt.fields.DBName,
				ConnectionString: tt.fields.ConnectionString,
			}
			customerFake := customer_infrastructure_fake.NewFakeCustomerRepositor()
			if err := o.Create(tt.args.tm, customerFake, tt.args.dto); (err != nil) != tt.wantErr {
				t.Errorf("OrderService.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
