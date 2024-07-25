// Code generated by mockery v2.43.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// APIError is an autogenerated mock type for the APIError type
type APIError struct {
	mock.Mock
}

type APIError_Expecter struct {
	mock *mock.Mock
}

func (_m *APIError) EXPECT() *APIError_Expecter {
	return &APIError_Expecter{mock: &_m.Mock}
}

// Code provides a mock function with given fields:
func (_m *APIError) Code() int {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Code")
	}

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// APIError_Code_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Code'
type APIError_Code_Call struct {
	*mock.Call
}

// Code is a helper method to define mock.On call
func (_e *APIError_Expecter) Code() *APIError_Code_Call {
	return &APIError_Code_Call{Call: _e.mock.On("Code")}
}

func (_c *APIError_Code_Call) Run(run func()) *APIError_Code_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *APIError_Code_Call) Return(_a0 int) *APIError_Code_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *APIError_Code_Call) RunAndReturn(run func() int) *APIError_Code_Call {
	_c.Call.Return(run)
	return _c
}

// Name provides a mock function with given fields:
func (_m *APIError) Name() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Name")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// APIError_Name_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Name'
type APIError_Name_Call struct {
	*mock.Call
}

// Name is a helper method to define mock.On call
func (_e *APIError_Expecter) Name() *APIError_Name_Call {
	return &APIError_Name_Call{Call: _e.mock.On("Name")}
}

func (_c *APIError_Name_Call) Run(run func()) *APIError_Name_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *APIError_Name_Call) Return(_a0 string) *APIError_Name_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *APIError_Name_Call) RunAndReturn(run func() string) *APIError_Name_Call {
	_c.Call.Return(run)
	return _c
}

// NewAPIError creates a new instance of APIError. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAPIError(t interface {
	mock.TestingT
	Cleanup(func())
},
) *APIError {
	mock := &APIError{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
