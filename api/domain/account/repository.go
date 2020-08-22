package account

import (
	"context"
	"github.com/alefherrera/accounting/api/domain/models"
)

type Repository interface {
	CommitTransaction(ctx context.Context, transaction models.Transaction) error
}