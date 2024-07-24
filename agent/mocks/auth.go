// Code generated by mockery v2.42.3. DO NOT EDIT.

// Copyright (c) Ultraviolet
// SPDX-License-Identifier: Apache-2.0

package mocks

import (
	context "context"

	auth "github.com/ultravioletrs/cocos/agent/auth"

	mock "github.com/stretchr/testify/mock"
)

// Authenticator is an autogenerated mock type for the Authenticator type
type Authenticator struct {
	mock.Mock
}

// AuthenticateUser provides a mock function with given fields: ctx, role
func (_m *Authenticator) AuthenticateUser(ctx context.Context, role auth.UserRole) (context.Context, error) {
	ret := _m.Called(ctx, role)

	if len(ret) == 0 {
		panic("no return value specified for AuthenticateUser")
	}

	var r0 context.Context
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, auth.UserRole) (context.Context, error)); ok {
		return rf(ctx, role)
	}
	if rf, ok := ret.Get(0).(func(context.Context, auth.UserRole) context.Context); ok {
		r0 = rf(ctx, role)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, auth.UserRole) error); ok {
		r1 = rf(ctx, role)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewAuthenticator creates a new instance of Authenticator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAuthenticator(t interface {
	mock.TestingT
	Cleanup(func())
}) *Authenticator {
	mock := &Authenticator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}