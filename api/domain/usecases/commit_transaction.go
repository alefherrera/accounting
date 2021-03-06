package usecases

import (
	"context"
	"errors"
	"github.com/alefherrera/accounting/api/domain/account"
	"github.com/alefherrera/accounting/api/domain/models"
)

type CommitTransactionInput struct {
	TransactionType models.TransactionType `json:"transaction_type"`
	Amount          float64                `json:"amount"`
}

type CommitTransactionOutput struct {
	Id      string  `json:"id"`
	Balance float64 `json:"balance"`
}

const (
	UnableToCommitTransaction = "unable to commit transaction"
	BalanceNotFound           = "balance not found"
	TransactionRefused        = "transaction refused"
)

type CommitTransaction interface {
	Execute(ctx context.Context, input CommitTransactionInput) (*CommitTransactionOutput, error)
}

var _ CommitTransaction = (*commitTransactionImpl)(nil)

type commitTransactionImpl struct {
	accountRepository account.Repository
}

func NewCommitTransactionImpl(accountRepository account.Repository) *commitTransactionImpl {
	return &commitTransactionImpl{accountRepository: accountRepository}
}

func (c commitTransactionImpl) Execute(ctx context.Context, input CommitTransactionInput) (*CommitTransactionOutput, error) {

	account, err := c.accountRepository.Get(ctx)

	if err != nil {
		return nil, errors.New(UnableToCommitTransaction)
	}

	if account == nil {
		return nil, errors.New(BalanceNotFound)
	}

	if input.TransactionType == models.TransactionTypeDebit {
		if account.GetBalance()-input.Amount < 0 {
			return nil, errors.New(TransactionRefused)
		}
	}

	newTransaction := models.NewTransaction(input.TransactionType, input.Amount)

	account.CommitTransaction(*newTransaction)

	if err := c.accountRepository.Save(ctx, *account); err != nil {
		return nil, errors.New(UnableToCommitTransaction)
	}

	balance := account.GetBalance()

	return &CommitTransactionOutput{
		Id:      newTransaction.Id,
		Balance: balance,
	}, nil
}
