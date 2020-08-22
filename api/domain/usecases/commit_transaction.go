package usecases

import (
	"context"
	"github.com/alefherrera/accounting/api/domain/account"
	"github.com/alefherrera/accounting/api/domain/models"
)

type CommitTransactionInput struct {
	TransactionType models.TransactionType
	Amount          float64
}

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
	return c.accountRepository.CommitTransaction(ctx, *newTransaction)
}
