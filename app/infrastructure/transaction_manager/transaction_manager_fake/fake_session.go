package transaction_manager_fake

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type FakeSession struct{}

func (fs *FakeSession) StartTransaction(opts ...*options.TransactionOptions) error {
	return nil // トランザクション開始処理の模倣
}

func (fs *FakeSession) AbortTransaction(ctx context.Context) error {
	return nil // トランザクション中止処理の模倣
}

func (fs *FakeSession) CommitTransaction(ctx context.Context) error {
	return nil // トランザクションコミット処理の模倣
}

func (fs *FakeSession) WithTransaction(
	ctx context.Context,
	fn func(sc mongo.SessionContext) (interface{}, error),
	opts ...*options.TransactionOptions) (interface{}, error) {
	return nil, nil
}

func (fs *FakeSession) EndSession(ctx context.Context) {
	// セッション終了処理の模倣
}

func (fs *FakeSession) ClusterTime() bson.Raw {
	return bson.Raw{}
}

func (fs *FakeSession) OperationTime() *primitive.Timestamp {
	return &primitive.Timestamp{}
}

func (fs *FakeSession) Client() *mongo.Client {
	return nil
}

func (fs *FakeSession) ID() bson.Raw {
	return bson.Raw{}
}

func (fs *FakeSession) AdvanceClusterTime(bson.Raw) error {
	return nil
}

func (fs *FakeSession) AdvanceOperationTime(*primitive.Timestamp) error {
	return nil
}

func (fs *FakeSession) session() {
}
