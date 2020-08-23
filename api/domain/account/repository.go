package account

import (
	"context"
	"github.com/alefherrera/accounting/api/domain/models"
	"sync"
)

type Repository interface {
	Save(ctx context.Context, account models.Account) error
	Get(ctx context.Context) (*models.Account, error)
}

type inmemoryRepository struct {
	mux     sync.RWMutex
	account models.Account
}

func NewInmemoryRepository() *inmemoryRepository {
	return &inmemoryRepository{}
}

func (i *inmemoryRepository) Save(ctx context.Context, account models.Account) error {
	i.mux.Lock()
	defer i.mux.Unlock()
	i.account = account
	return nil
}

func (i *inmemoryRepository) Get(ctx context.Context) (*models.Account, error) {
	i.mux.RLock()
	defer i.mux.RUnlock()
	return &i.account, nil
}
