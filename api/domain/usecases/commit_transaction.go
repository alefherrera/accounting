package usecases

import (
	"context"
	"errors"
	"github.com/alefherrera/accounting/api/domain/account"
	"github.com/alefherrera/accounting/api/domain/models"
)

type CommitTransactionInput struct {
	TransactionType models.TransactionType
	Amount          float64
}

const UnableToCommitTransaction = "unable to commit transaction"

type CommitTransaction interface {
	Execute(ctx context.Context, input CommitTransactionInput) error
}

var _ CommitTransaction = (*commitTransactionImpl)(nil)

type commitTransactionImpl struct {
	accountRepository account.Repository
}

func NewCommitTransactionImpl(accountRepository account.Repository) *commitTransactionImpl {
	return &commitTransactionImpl{accountRepository: accountRepository}
}

func (c commitTransactionImpl) Execute(ctx context.Context, input CommitTransactionInput) error {
	newTransaction := models.NewTransaction(input.TransactionType, input.Amount)

	if err := c.accountRepository.CommitTransaction(ctx, *newTransaction); err != nil {
		return errors.New(UnableToCommitTransaction)
	}

	return nil
}
