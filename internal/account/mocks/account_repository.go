// Code generated by mockery v2.51.1. DO NOT EDIT.

package mocks

import (
	context "context"
	entity "accounts-and-transactions/internal/entity"

	mock "github.com/stretchr/testify/mock"
)

// AccountRepository is an autogenerated mock type for the AccountRepository type
type AccountRepository struct {
	mock.Mock
}

// CreateAccount provides a mock function with given fields: ctx, req
func (_m *AccountRepository) CreateAccount(ctx context.Context, req *entity.Account) (*entity.Account, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for CreateAccount")
	}

	var r0 *entity.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Account) (*entity.Account, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Account) *entity.Account); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Account)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entity.Account) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAccountById provides a mock function with given fields: ctx, id
func (_m *AccountRepository) GetAccountById(ctx context.Context, id string) (*entity.Account, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetAccountById")
	}

	var r0 *entity.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*entity.Account, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *entity.Account); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Account)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewAccountRepository creates a new instance of AccountRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAccountRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *AccountRepository {
	mock := &AccountRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
