// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// RoleManager is an autogenerated mock type for the RoleManager type
type RoleManager struct {
	mock.Mock
}

type RoleManager_Expecter struct {
	mock *mock.Mock
}

func (_m *RoleManager) EXPECT() *RoleManager_Expecter {
	return &RoleManager_Expecter{mock: &_m.Mock}
}

// CreateOrImport provides a mock function with given fields: ctx, clusterName
func (_m *RoleManager) CreateOrImport(ctx context.Context, clusterName string) (string, error) {
	ret := _m.Called(ctx, clusterName)

	if len(ret) == 0 {
		panic("no return value specified for CreateOrImport")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (string, error)); ok {
		return rf(ctx, clusterName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, clusterName)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, clusterName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RoleManager_CreateOrImport_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateOrImport'
type RoleManager_CreateOrImport_Call struct {
	*mock.Call
}

// CreateOrImport is a helper method to define mock.On call
//   - ctx context.Context
//   - clusterName string
func (_e *RoleManager_Expecter) CreateOrImport(ctx interface{}, clusterName interface{}) *RoleManager_CreateOrImport_Call {
	return &RoleManager_CreateOrImport_Call{Call: _e.mock.On("CreateOrImport", ctx, clusterName)}
}

func (_c *RoleManager_CreateOrImport_Call) Run(run func(ctx context.Context, clusterName string)) *RoleManager_CreateOrImport_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *RoleManager_CreateOrImport_Call) Return(_a0 string, _a1 error) *RoleManager_CreateOrImport_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *RoleManager_CreateOrImport_Call) RunAndReturn(run func(context.Context, string) (string, error)) *RoleManager_CreateOrImport_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteIfRequired provides a mock function with given fields: ctx
func (_m *RoleManager) DeleteIfRequired(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for DeleteIfRequired")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RoleManager_DeleteIfRequired_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteIfRequired'
type RoleManager_DeleteIfRequired_Call struct {
	*mock.Call
}

// DeleteIfRequired is a helper method to define mock.On call
//   - ctx context.Context
func (_e *RoleManager_Expecter) DeleteIfRequired(ctx interface{}) *RoleManager_DeleteIfRequired_Call {
	return &RoleManager_DeleteIfRequired_Call{Call: _e.mock.On("DeleteIfRequired", ctx)}
}

func (_c *RoleManager_DeleteIfRequired_Call) Run(run func(ctx context.Context)) *RoleManager_DeleteIfRequired_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *RoleManager_DeleteIfRequired_Call) Return(_a0 error) *RoleManager_DeleteIfRequired_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RoleManager_DeleteIfRequired_Call) RunAndReturn(run func(context.Context) error) *RoleManager_DeleteIfRequired_Call {
	_c.Call.Return(run)
	return _c
}

// NewRoleManager creates a new instance of RoleManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRoleManager(t interface {
	mock.TestingT
	Cleanup(func())
}) *RoleManager {
	mock := &RoleManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}