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

			account := models.Account{}

			balance := float64(100)
			accountRepository.On("Get", ctx).Return(&account, nil)
			accountRepository.On("Save", ctx, mock.Anything).Return(nil)

			commitTransactionImpl := NewCommitTransactionImpl(accountRepository)

			result, err := commitTransactionImpl.Execute(ctx, CommitTransactionInput{
				TransactionType: models.TransactionTypeCredit,
				Amount:          100,
			})

			assert.NoError(t, err)
			assert.Equal(t, balance, *result)
		})

		t.Run("error saving", func(t *testing.T) {
			accountRepository := new(mocks.Repository)
			defer accountRepository.AssertExpectations(t)

			account := models.Account{}

			accountRepository.On("Get", ctx).Return(&account, nil)
			accountRepository.On("Save", ctx, mock.Anything).Return(errors.New("error saving account"))

			commitTransactionImpl := NewCommitTransactionImpl(accountRepository)

			result, err := commitTransactionImpl.Execute(ctx, CommitTransactionInput{
				TransactionType: models.TransactionTypeCredit,
				Amount:          100,
			})

			assert.Nil(t, result)
			assert.EqualError(t, err, UnableToCommitTransaction)
		})
	})

	t.Run("Debit Transaction", func(t *testing.T) {

		t.Run("accept transaction if balance is > 0", func(t *testing.T) {
			accountRepository := new(mocks.Repository)
			defer accountRepository.AssertExpectations(t)

			account := models.Account{
				Transactions: []models.Transaction{
					{
						TransactionType: models.TransactionTypeCredit,
						Amount:          101,
					},
				},
			}

			accountRepository.On("Get", ctx).Return(&account, nil)
			accountRepository.On("Save", ctx, mock.Anything).Return(nil)

			commitTransactionImpl := NewCommitTransactionImpl(accountRepository)

			transactionAmount := float64(100)
			balance, err := commitTransactionImpl.Execute(ctx, CommitTransactionInput{
				TransactionType: models.TransactionTypeDebit,
				Amount:          transactionAmount,
			})

			assert.NoError(t, err)
			assert.Equal(t, float64(1), *balance)

		})

		t.Run("accept transaction if balance is = 0", func(t *testing.T) {
			accountRepository := new(mocks.Repository)
			defer accountRepository.AssertExpectations(t)

			amount := float64(100)
			account := models.Account{
				Transactions: []models.Transaction{
					{
						TransactionType: models.TransactionTypeCredit,
						Amount:          amount,
					},
				},
			}

			accountRepository.On("Get", ctx).Return(&account, nil)
			accountRepository.On("Save", ctx, mock.Anything).Return(nil)

			commitTransactionImpl := NewCommitTransactionImpl(accountRepository)

			balance, err := commitTransactionImpl.Execute(ctx, CommitTransactionInput{
				TransactionType: models.TransactionTypeDebit,
				Amount:          amount,
			})

			assert.NoError(t, err)
			assert.Equal(t, float64(0), *balance)
		})

		t.Run("refuse negative amount on account", func(t *testing.T) {
			accountRepository := new(mocks.Repository)
			defer accountRepository.AssertExpectations(t)

			account := models.Account{}
			accountRepository.On("Get", ctx).Return(&account, nil)

			commitTransactionImpl := NewCommitTransactionImpl(accountRepository)

			result, err := commitTransactionImpl.Execute(ctx, CommitTransactionInput{
				TransactionType: models.TransactionTypeDebit,
				Amount:          100,
			})

			assert.Nil(t, result)
			assert.EqualError(t, err, TransactionRefused)
		})

		t.Run("refuse transaction when unable get balance", func(t *testing.T) {
			accountRepository := new(mocks.Repository)
			defer accountRepository.AssertExpectations(t)

			accountRepository.On("Get", ctx).Return(nil, errors.New("error"))

			commitTransactionImpl := NewCommitTransactionImpl(accountRepository)

			result, err := commitTransactionImpl.Execute(ctx, CommitTransactionInput{
				TransactionType: models.TransactionTypeDebit,
				Amount:          100,
			})

			assert.Nil(t, result)
			assert.EqualError(t, err, UnableToGetBalance)
		})

		t.Run("refuse transaction when balance not found", func(t *testing.T) {
			accountRepository := new(mocks.Repository)
			defer accountRepository.AssertExpectations(t)

			accountRepository.On("Get", ctx).Return(nil, nil)

			commitTransactionImpl := NewCommitTransactionImpl(accountRepository)

			result, err := commitTransactionImpl.Execute(ctx, CommitTransactionInput{
				TransactionType: models.TransactionTypeDebit,
				Amount:          100,
			})

			assert.Nil(t, result)
			assert.EqualError(t, err, BalanceNotFound)
		})

		t.Run("error saving", func(t *testing.T) {
			accountRepository := new(mocks.Repository)
			defer accountRepository.AssertExpectations(t)

			account := models.Account{}

			accountRepository.On("Get", ctx).Return(&account, nil)
			accountRepository.On("Save", ctx, mock.Anything).Return(errors.New("error saving account"))

			commitTransactionImpl := NewCommitTransactionImpl(accountRepository)

			result, err := commitTransactionImpl.Execute(ctx, CommitTransactionInput{
				TransactionType: models.TransactionTypeDebit,
				Amount:          0,
			})

			assert.Nil(t, result)
			assert.EqualError(t, err, UnableToCommitTransaction)
		})
	})

}
