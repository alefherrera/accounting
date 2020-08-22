package usecases

import (
	"context"
	"errors"
	"github.com/alefherrera/accounting/api/domain/account/mocks"
	"github.com/alefherrera/accounting/api/domain/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func Test_commitTransactionImpl_Execute(t *testing.T) {
	ctx := context.TODO()

	t.Run("Must accept credit transaction", func(t *testing.T) {
		accountRepository := new(mocks.Repository)
		defer accountRepository.AssertExpectations(t)

		accountRepository.On("CommitTransaction", ctx, mock.Anything).Return(nil)

		commitTransactionImpl := NewCommitTransactionImpl(accountRepository)

		err := commitTransactionImpl.Execute(ctx, CommitTransactionInput{
			TransactionType: models.TransactionTypeCredit,
			Amount:          100,
		})

		assert.NoError(t, err)
	})

	t.Run("error saving returns error on commit", func(t *testing.T) {
		accountRepository := new(mocks.Repository)
		defer accountRepository.AssertExpectations(t)

		accountRepository.On("CommitTransaction", ctx, mock.Anything).Return(errors.New("error saving account"))

		commitTransactionImpl := NewCommitTransactionImpl(accountRepository)

		err := commitTransactionImpl.Execute(ctx, CommitTransactionInput{
			TransactionType: models.TransactionTypeCredit,
			Amount:          100,
		})

		assert.EqualError(t, err, UnableToCommitTransaction)
	})

	t.Run("refuse negative amount on account", func(t *testing.T) {
		accountRepository := new(mocks.Repository)
		defer accountRepository.AssertExpectations(t)

		accountRepository.On("GetBalance", ctx).Return(float64(0), nil)

		commitTransactionImpl := NewCommitTransactionImpl(accountRepository)

		err := commitTransactionImpl.Execute(ctx, CommitTransactionInput{
			TransactionType: models.TransactionTypeDebit,
			Amount:          100,
		})

		assert.EqualError(t, err, TransactionRefused)
	})

}
