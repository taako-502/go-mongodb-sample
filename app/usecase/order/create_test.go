package order_usecase

import (
	"context"
	"go-mongodb-sample/app/infrastructure"
	customer_infrastructure "go-mongodb-sample/app/infrastructure/customers"
	"go-mongodb-sample/app/infrastructure/customers/customer_infrastructure_fake"
	"testing"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestOrderService_Create(t *testing.T) {
	emptyID, _ := primitive.ObjectIDFromHex("000000000000000000000000")
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
			name: "カスタマーが存在しない",
			fields: fields{
				Ctx:              context.Background(),
				DBName:           "test",
				ConnectionString: "mongodb://localhost:27017",
			},
			args: args{
				tm:  &infrastructure.MongoTransactionManager{},
				dto: CreateDTO{CustomerID: emptyID},
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
			customerFake := customer_infrastructure_fake.NewFakeProductRepository()
			customerRepo, ok := customerFake.(*customer_infrastructure.CustomerRepository) // TODO: 本来は必要なかったと思う
			require.True(t, ok)                                                            // NOTE: やはりここでエラーになる
			if err := o.Create(tt.args.tm, customerRepo, tt.args.dto); (err != nil) != tt.wantErr {
				t.Errorf("OrderService.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
