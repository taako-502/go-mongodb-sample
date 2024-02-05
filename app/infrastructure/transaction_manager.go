package infrastructure

import (
	"context"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoTransactionManager struct {
	Ctx    context.Context
	Client *mongo.Client
}

func NewMongoTransactionManager(ctx context.Context, client *mongo.Client) *MongoTransactionManager {
	return &MongoTransactionManager{Ctx: ctx, Client: client}
}

func (tm *MongoTransactionManager) StartSession() (mongo.Session, error) {
	return tm.Client.StartSession()
}

func (tm *MongoTransactionManager) WithSession(
	ctx context.Context,
	sess mongo.Session,
	fn func(sc context.Context) error) error {
	return mongo.WithSession(ctx, sess, func(sc mongo.SessionContext) error {
		if err := sc.StartTransaction(); err != nil {
			return errors.Wrap(err, "sc.StartTransaction()")
		}

		if err := fn(sc); err != nil {
			return errors.Wrap(err, "fn(mongo.SessionContext)")
		}

		if erro := sc.CommitTransaction(sc); erro != nil {
			return errors.Wrap(erro, "sc.CommitTransaction(sc)")
		}

		return nil
	})
}
