package transaction_manager

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
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

	return mongo.WithSession(tm.Ctx, sess, func(ctx context.Context) error {
		if err := sess.StartTransaction(options.Transaction()); err != nil {
			return fmt.Errorf("sc.StartTransaction(): %w", err)
		}

		if err := fn(ctx); err != nil {
			return fmt.Errorf("fn(mongo.SessionContext): %w", err)
		}

		if erro := sess.CommitTransaction(ctx); erro != nil {
			return fmt.Errorf("sc.CommitTransaction(sc): %w", erro)
		}

		return nil
	})
}
