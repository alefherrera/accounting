package transaction

import (
	"context"
	"github.com/alefherrera/accounting/api/domain/models"
)

type Repository interface {
	Add(ctx context.Context, transaction models.Transaction) error
}
