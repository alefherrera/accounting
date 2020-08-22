package account

import (
	"context"
	"github.com/alefherrera/accounting/api/domain/models"
)

type Repository interface {
	Save(ctx context.Context, account models.Account) error
	Get(ctx context.Context) (*models.Account, error)
}

type inmemoryRepository struct {
	account models.Account
}

func NewInmemoryRepository() *inmemoryRepository {
	return &inmemoryRepository{}
}

func (i *inmemoryRepository) Save(ctx context.Context, account models.Account) error {
	i.account = account
	return nil
}

func (i inmemoryRepository) Get(ctx context.Context) (*models.Account, error) {
	return &i.account, nil
}
