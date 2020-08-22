package account

import (
	"context"
	"github.com/alefherrera/accounting/api/domain/models"
)

type Repository interface {
	Save(ctx context.Context, account models.Account) error
	Get(ctx context.Context) (*models.Account, error)
}
