package transaction_manager_fake

import (
	"context"
	"go-mongodb-sample/app/infrastructure/transaction_manager"
)

type TransactionManagerFake struct{}

func NewFakeTransactionManager() transaction_manager.TransactionManager {
	return &TransactionManagerFake{}
}

func (tm *TransactionManagerFake) StartSession() error {
	return nil
}

func (tm *TransactionManagerFake) EndSession() {
}

func (tm *TransactionManagerFake) WithSession(fn func(sc context.Context) error) error {
	return nil
}
