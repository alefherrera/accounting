package usecases

import (
	"context"
	"errors"
	"github.com/alefherrera/accounting/api/domain/account"
)

const (
	UnableToGetBalance = "unable to get balance"
)

type GetBalanceOutput struct {
	Balance float64 `json:"balance"`
}

type GetBalance interface {
	Execute(ctx context.Context) (*GetBalanceOutput, error)
}

var _ GetBalance = (*getBalanceImpl)(nil)

type getBalanceImpl struct {
	accountRepository account.Repository
}

func NewGetBalanceImpl(accountRepository account.Repository) *getBalanceImpl {
	return &getBalanceImpl{accountRepository: accountRepository}
}

func (g getBalanceImpl) Execute(ctx context.Context) (*GetBalanceOutput, error) {
	account, err := g.accountRepository.Get(ctx)

	if err != nil {
		return nil, errors.New(UnableToGetBalance)
	}

	if account == nil {
		return nil, nil
	}

	return &GetBalanceOutput{
		Balance: account.GetBalance(),
	}, nil
}
