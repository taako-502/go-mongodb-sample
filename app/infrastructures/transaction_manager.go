package infrastructures

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongoTransactionManager struct {
	Client *mongo.Client
}

func NewMongoTransactionManager(client *mongo.Client) *MongoTransactionManager {
	return &MongoTransactionManager{Client: client}
}

func (tm *MongoTransactionManager) StartSession() (mongo.Session, error) {
	return tm.Client.StartSession()
}

func (tm *MongoTransactionManager) WithSession(ctx context.Context, sess mongo.Session, fn func(sc mongo.SessionContext) error) error {
	return mongo.WithSession(ctx, sess, fn)
}
