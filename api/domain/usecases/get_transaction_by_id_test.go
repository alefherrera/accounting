package usecases

import (
	"context"
	"errors"
	"github.com/alefherrera/accounting/api/domain/account/mocks"
	"github.com/alefherrera/accounting/api/domain/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_getTransactionByIdImpl_Execute(t *testing.T) {
	ctx := context.TODO()

	t.Run("not found", func(t *testing.T) {
		accountRepository := new(mocks.Repository)
		defer accountRepository.AssertExpectations(t)

		accountRepository.On("Get", ctx).Return(nil, nil)

		getTransactionsImpl := NewGetTransactionByIdImpl(accountRepository)

		transaction, err := getTransactionsImpl.Execute(ctx, "id")

		assert.NoError(t, err)
		assert.Nil(t, transaction)
	})

	t.Run("unable to get transactions", func(t *testing.T) {
		accountRepository := new(mocks.Repository)
		defer accountRepository.AssertExpectations(t)

		accountRepository.On("Get", ctx).Return(nil, errors.New("error"))

		getTransactionsImpl := NewGetTransactionByIdImpl(accountRepository)

		transaction, err := getTransactionsImpl.Execute(ctx, "id")

		assert.Nil(t, transaction)
		assert.EqualError(t, err, UnableToGetTransaction)
	})

	t.Run("success", func(t *testing.T) {
		accountRepository := new(mocks.Repository)
		defer accountRepository.AssertExpectations(t)

		id := "id1"
		transaction := models.Transaction{
			Id:              id,
			TransactionType: models.TransactionTypeCredit,
			Amount:          101,
		}
		transactions := []models.Transaction{
			transaction,
		}

		account := models.Account{
			Transactions: transactions,
		}

		accountRepository.On("Get", ctx).Return(&account, nil)

		getTransactionsImpl := NewGetTransactionByIdImpl(accountRepository)

		result, err := getTransactionsImpl.Execute(ctx, id)

		assert.NoError(t, err)
		assert.Equal(t, transaction, *result)
	})

	t.Run("not found with another transactions", func(t *testing.T) {
		accountRepository := new(mocks.Repository)
		defer accountRepository.AssertExpectations(t)

		id := "id1"
		transaction := models.Transaction{
			Id:              id,
			TransactionType: models.TransactionTypeCredit,
			Amount:          101,
		}
		transactions := []models.Transaction{
			transaction,
		}

		account := models.Account{
			Transactions: transactions,
		}

		accountRepository.On("Get", ctx).Return(&account, nil)

		getTransactionsImpl := NewGetTransactionByIdImpl(accountRepository)

		result, err := getTransactionsImpl.Execute(ctx, "id")

		assert.NoError(t, err)
		assert.Nil(t, result)
	})
}
