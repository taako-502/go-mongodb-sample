package transaction_manager

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/mongo"
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
		return fmt.Errorf("tm.Client.StartSession: %w", err)
	}
	defer sess.EndSession(tm.Ctx)

	return mongo.WithSession(tm.Ctx, sess, func(sc mongo.SessionContext) error {
		if err := sc.StartTransaction(); err != nil {
			return fmt.Errorf("sc.StartTransaction(): %w", err)
		}

		if err := fn(sc); err != nil {
			return fmt.Errorf("fn(mongo.SessionContext): %w", err)
		}

		if erro := sc.CommitTransaction(sc); erro != nil {
			return fmt.Errorf("sc.CommitTransaction(sc): %w", erro)
		}

		return nil
	})
}
