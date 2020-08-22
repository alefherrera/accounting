package usecases

import (
	"context"
	"github.com/alefherrera/accounting/api/domain/transaction"
)

type CommitTransactionInput struct {
}

type CommitTransactionOutput struct {
}

type CommitTransaction interface {
	Execute(ctx context.Context, input CommitTransactionInput) (CommitTransactionOutput, error)
}

var _ CommitTransaction = (*commitTransactionImpl)(nil)

type commitTransactionImpl struct {
	transactionRepository transaction.Repository
}

func (c commitTransactionImpl) Execute(ctx context.Context, input CommitTransactionInput) (CommitTransactionOutput, error) {
	c.transactionRepository.Add(ctx, struct{}{})
	return CommitTransactionOutput{}, nil
}
