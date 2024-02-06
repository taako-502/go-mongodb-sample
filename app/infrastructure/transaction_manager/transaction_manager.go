package transaction_manager

import (
	"context"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
)

type TransactionManager interface {
	WithSession(fn func(sc context.Context) error) error
}

type MongoTransactionManager struct {
	Ctx     context.Context
	Client  *mongo.Client
	Session mongo.Session
}

func NewMongoTransactionManager(ctx context.Context, client *mongo.Client) *MongoTransactionManager {
	return &MongoTransactionManager{Ctx: ctx, Client: client}
}

func (tm *MongoTransactionManager) WithSession(fn func(sc context.Context) error) error {
	sess, err := tm.Client.StartSession()
	if err != nil {
		return errors.Wrap(err, "tm.Client.StartSession")
	}
	defer sess.EndSession(tm.Ctx)

	return mongo.WithSession(tm.Ctx, sess, func(sc mongo.SessionContext) error {
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
