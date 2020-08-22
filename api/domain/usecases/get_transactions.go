package usecases

import (
	"context"
	"errors"
	"github.com/alefherrera/accounting/api/domain/account"
	"github.com/alefherrera/accounting/api/domain/models"
)

const (
	UnableToGetTransactions = "unable to get transactions"
)

type GetTransactions interface {
	Execute(ctx context.Context) ([]models.Transaction, error)
}

var _ GetTransactions = (*getTransactionsImpl)(nil)

type getTransactionsImpl struct {
	accountRepository account.Repository
}

func NewGetTransactionsImpl(accountRepository account.Repository) *getTransactionsImpl {
	return &getTransactionsImpl{accountRepository: accountRepository}
}

func (g getTransactionsImpl) Execute(ctx context.Context) ([]models.Transaction, error) {
	account, err := g.accountRepository.Get(ctx)

	if err != nil {
		return nil, errors.New(UnableToGetTransactions)
	}

	if account == nil {
		return nil, nil
	}

	return account.Transactions, nil
}
