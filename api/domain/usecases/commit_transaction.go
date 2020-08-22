package usecases

import "context"

type CommitTransactionInput struct {
}

type CommitTransactionOutput struct {
}

type CommitTransaction interface {
	Execute(ctx context.Context, input CommitTransactionInput) (CommitTransactionOutput, error)
}
