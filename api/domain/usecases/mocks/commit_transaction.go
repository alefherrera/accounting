// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	usecases "github.com/alefherrera/accounting/api/domain/usecases"
	mock "github.com/stretchr/testify/mock"
)

// CommitTransaction is an autogenerated mock type for the CommitTransaction type
type CommitTransaction struct {
	mock.Mock
}

// Execute provides a mock function with given fields: ctx, input
func (_m *CommitTransaction) Execute(ctx context.Context, input usecases.CommitTransactionInput) (usecases.CommitTransactionOutput, error) {
	ret := _m.Called(ctx, input)

	var r0 usecases.CommitTransactionOutput
	if rf, ok := ret.Get(0).(func(context.Context, usecases.CommitTransactionInput) usecases.CommitTransactionOutput); ok {
		r0 = rf(ctx, input)
	} else {
		r0 = ret.Get(0).(usecases.CommitTransactionOutput)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, usecases.CommitTransactionInput) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
