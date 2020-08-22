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

	t.Run("Credit transaction", func(t *testing.T) {
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

		t.Run("error saving", func(t *testing.T) {
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
	})

	t.Run("Debit Transaction", func(t *testing.T) {

		t.Run("accept transaction if balance is > 0", func(t *testing.T) {
			accountRepository := new(mocks.Repository)
			defer accountRepository.AssertExpectations(t)

			balance := float64(101)
			accountRepository.On("GetBalance", ctx).Return(&balance, nil)
			accountRepository.On("CommitTransaction", ctx, mock.Anything).Return(nil)

			commitTransactionImpl := NewCommitTransactionImpl(accountRepository)

			err := commitTransactionImpl.Execute(ctx, CommitTransactionInput{
				TransactionType: models.TransactionTypeDebit,
				Amount:          100,
			})

			assert.NoError(t, err)
		})

		t.Run("accept transaction if balance is = 0", func(t *testing.T) {
			accountRepository := new(mocks.Repository)
			defer accountRepository.AssertExpectations(t)

			amount := float64(100)
			balance := amount
			accountRepository.On("GetBalance", ctx).Return(&balance, nil)
			accountRepository.On("CommitTransaction", ctx, mock.Anything).Return(nil)

			commitTransactionImpl := NewCommitTransactionImpl(accountRepository)

			err := commitTransactionImpl.Execute(ctx, CommitTransactionInput{
				TransactionType: models.TransactionTypeDebit,
				Amount:          amount,
			})

			assert.NoError(t, err)
		})

		t.Run("refuse negative amount on account", func(t *testing.T) {
			accountRepository := new(mocks.Repository)
			defer accountRepository.AssertExpectations(t)

			balance := float64(0)
			accountRepository.On("GetBalance", ctx).Return(&balance, nil)

			commitTransactionImpl := NewCommitTransactionImpl(accountRepository)

			err := commitTransactionImpl.Execute(ctx, CommitTransactionInput{
				TransactionType: models.TransactionTypeDebit,
				Amount:          100,
			})

			assert.EqualError(t, err, TransactionRefused)
		})

		t.Run("refuse transaction when unable get balance", func(t *testing.T) {
			accountRepository := new(mocks.Repository)
			defer accountRepository.AssertExpectations(t)

			accountRepository.On("GetBalance", ctx).Return(nil, errors.New("error"))

			commitTransactionImpl := NewCommitTransactionImpl(accountRepository)

			err := commitTransactionImpl.Execute(ctx, CommitTransactionInput{
				TransactionType: models.TransactionTypeDebit,
				Amount:          100,
			})

			assert.EqualError(t, err, UnableToGetBalance)
		})

		t.Run("refuse transaction when balance not found", func(t *testing.T) {
			accountRepository := new(mocks.Repository)
			defer accountRepository.AssertExpectations(t)

			accountRepository.On("GetBalance", ctx).Return(nil, nil)

			commitTransactionImpl := NewCommitTransactionImpl(accountRepository)

			err := commitTransactionImpl.Execute(ctx, CommitTransactionInput{
				TransactionType: models.TransactionTypeDebit,
				Amount:          100,
			})

			assert.EqualError(t, err, BalanceNotFound)
		})

		t.Run("error saving", func(t *testing.T) {
			accountRepository := new(mocks.Repository)
			defer accountRepository.AssertExpectations(t)

			balance := float64(0)
			accountRepository.On("GetBalance", ctx).Return(&balance, nil)
			accountRepository.On("CommitTransaction", ctx, mock.Anything).Return(errors.New("error saving account"))

			commitTransactionImpl := NewCommitTransactionImpl(accountRepository)

			err := commitTransactionImpl.Execute(ctx, CommitTransactionInput{
				TransactionType: models.TransactionTypeDebit,
				Amount:          0,
			})

			assert.EqualError(t, err, UnableToCommitTransaction)
		})
	})

}
