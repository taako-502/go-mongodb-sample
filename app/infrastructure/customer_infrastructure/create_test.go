package customer_infrastructure

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/taako-502/go-mongodb-sample/test"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestCustomerRepository_Create(t *testing.T) {
	type args struct {
		dto *CustomerDTO
	}
	tests := []struct {
		name    string
		args    args
		want    *CustomerDTO
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				dto: NewCustomerDTO("test", "test@test.com", "Egypt Cairo"),
			},
			want:    NewCustomerDTO("test", "test@test.com", "Egypt Cairo"),
			wantErr: false,
		},
	}

	// MongoDBに接続
	ctx := context.Background()
	ci, err := test.NewCIDatabaseConfig(ctx)
	require.NoError(t, err, "MongoDBに接続できません")
	defer ci.Client.Disconnect(ctx)

	// テストデータを削除
	_, err = ci.GetDatabaseInstance().Collection(GetName()).DeleteOne(ctx, bson.M{"name": "test"})
	require.NoError(t, err, "テストデータの削除に失敗しました")

	t.Cleanup(func() {
		// テストデータを削除
		ctx := context.Background()
		ci, err := test.NewCIDatabaseConfig(ctx)
		require.NoError(t, err, "MongoDBに接続できません")
		_, err = ci.GetDatabaseInstance().Collection(GetName()).DeleteOne(ctx, bson.M{"name": "test"})
		require.NoError(t, err, "テストデータの削除に失敗しました")
		defer ci.Client.Disconnect(ctx)
	})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCustomerRepository(ctx, ci.GetDatabaseInstance())
			got, err := c.Create(tt.args.dto)
			if (err != nil) != tt.wantErr {
				t.Errorf("CustomerRepository.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.EqualValues(t, tt.want.Name, got.Name)
			assert.EqualValues(t, tt.want.Email, got.Email)
			assert.EqualValues(t, tt.want.Address, got.Address)
		})
	}
}
