package usecases

import (
	"context"
	"errors"
	"github.com/alefherrera/accounting/api/domain/account"
	"github.com/alefherrera/accounting/api/domain/models"
)

const (
	UnableToGetTransaction = "unable to get transaction"
)

type GetTransactionById interface {
	Execute(ctx context.Context, id string) (*models.Transaction, error)
}

var _ GetTransactionById = (*getTransactionByIdImpl)(nil)

type getTransactionByIdImpl struct {
	accountRepository account.Repository
}

func NewGetTransactionByIdImpl(accountRepository account.Repository) *getTransactionByIdImpl {
	return &getTransactionByIdImpl{accountRepository: accountRepository}
}

func (g getTransactionByIdImpl) Execute(ctx context.Context, id string) (*models.Transaction, error) {
	account, err := g.accountRepository.Get(ctx)

	if err != nil {
		return nil, errors.New(UnableToGetTransaction)
	}

	if account == nil {
		return nil, nil
	}

	return account.GetTransactionById(id), nil
}
