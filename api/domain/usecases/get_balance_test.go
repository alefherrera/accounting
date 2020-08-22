package usecases

import (
	"context"
	"errors"
	"github.com/alefherrera/accounting/api/domain/account/mocks"
	"github.com/alefherrera/accounting/api/domain/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_getBalanceImpl_Execute(t *testing.T) {
	ctx := context.TODO()

	t.Run("not found", func(t *testing.T) {
		accountRepository := new(mocks.Repository)
		defer accountRepository.AssertExpectations(t)

		accountRepository.On("Get", ctx).Return(nil, nil)

		getBalanceImpl := NewGetBalanceImpl(accountRepository)

		result, err := getBalanceImpl.Execute(ctx)

		assert.NoError(t, err)
		assert.Nil(t, result)
	})

	t.Run("unable to get transactions", func(t *testing.T) {
		accountRepository := new(mocks.Repository)
		defer accountRepository.AssertExpectations(t)

		accountRepository.On("Get", ctx).Return(nil, errors.New("error"))

		getBalanceImpl := NewGetBalanceImpl(accountRepository)

		result, err := getBalanceImpl.Execute(ctx)

		assert.Nil(t, result)
		assert.EqualError(t, err, UnableToGetBalance)
	})

	t.Run("success", func(t *testing.T) {
		accountRepository := new(mocks.Repository)
		defer accountRepository.AssertExpectations(t)

		transactions := []models.Transaction{
			{
				TransactionType: models.TransactionTypeCredit,
				Amount:          101,
			},
		}

		account := models.Account{
			Transactions: transactions,
		}

		accountRepository.On("Get", ctx).Return(&account, nil)

		getBalanceImpl := NewGetBalanceImpl(accountRepository)

		result, err := getBalanceImpl.Execute(ctx)

		assert.NoError(t, err)
		assert.Equal(t, float64(101), result.Balance)
	})
}
