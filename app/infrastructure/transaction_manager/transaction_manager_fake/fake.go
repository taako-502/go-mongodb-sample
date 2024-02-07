package transaction_manager_fake

import (
	"context"

	"github.com/taako-502/go-mongodb-sample/app/infrastructure/transaction_manager"
)

type TransactionManagerFake struct{}

func NewFakeTransactionManager() transaction_manager.TransactionManager {
	return &TransactionManagerFake{}
}

func (tm *TransactionManagerFake) WithSession(fn func(sc context.Context) error) error {
	return fn(context.Background())
}
