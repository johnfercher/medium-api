// Code generated by mockery v2.43.1. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/johnfercher/medium-api/internal/core/models"
	mock "github.com/stretchr/testify/mock"
)

// ProductService is an autogenerated mock type for the ProductService type
type ProductService struct {
	mock.Mock
}

type ProductService_Expecter struct {
	mock *mock.Mock
}

func (_m *ProductService) EXPECT() *ProductService_Expecter {
	return &ProductService_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, product
func (_m *ProductService) Create(ctx context.Context, product *models.Product) (*models.Product, error) {
	ret := _m.Called(ctx, product)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *models.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Product) (*models.Product, error)); ok {
		return rf(ctx, product)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *models.Product) *models.Product); ok {
		r0 = rf(ctx, product)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Product)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *models.Product) error); ok {
		r1 = rf(ctx, product)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProductService_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type ProductService_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - product *models.Product
func (_e *ProductService_Expecter) Create(ctx interface{}, product interface{}) *ProductService_Create_Call {
	return &ProductService_Create_Call{Call: _e.mock.On("Create", ctx, product)}
}

func (_c *ProductService_Create_Call) Run(run func(ctx context.Context, product *models.Product)) *ProductService_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*models.Product))
	})
	return _c
}

func (_c *ProductService_Create_Call) Return(_a0 *models.Product, _a1 error) *ProductService_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ProductService_Create_Call) RunAndReturn(run func(context.Context, *models.Product) (*models.Product, error)) *ProductService_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, id
func (_m *ProductService) Delete(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ProductService_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type ProductService_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *ProductService_Expecter) Delete(ctx interface{}, id interface{}) *ProductService_Delete_Call {
	return &ProductService_Delete_Call{Call: _e.mock.On("Delete", ctx, id)}
}

func (_c *ProductService_Delete_Call) Run(run func(ctx context.Context, id string)) *ProductService_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *ProductService_Delete_Call) Return(_a0 error) *ProductService_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ProductService_Delete_Call) RunAndReturn(run func(context.Context, string) error) *ProductService_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *ProductService) GetByID(ctx context.Context, id string) (*models.Product, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetByID")
	}

	var r0 *models.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*models.Product, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.Product); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Product)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProductService_GetByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByID'
type ProductService_GetByID_Call struct {
	*mock.Call
}

// GetByID is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *ProductService_Expecter) GetByID(ctx interface{}, id interface{}) *ProductService_GetByID_Call {
	return &ProductService_GetByID_Call{Call: _e.mock.On("GetByID", ctx, id)}
}

func (_c *ProductService_GetByID_Call) Run(run func(ctx context.Context, id string)) *ProductService_GetByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *ProductService_GetByID_Call) Return(_a0 *models.Product, _a1 error) *ProductService_GetByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ProductService_GetByID_Call) RunAndReturn(run func(context.Context, string) (*models.Product, error)) *ProductService_GetByID_Call {
	_c.Call.Return(run)
	return _c
}

// Search provides a mock function with given fields: ctx, productType
func (_m *ProductService) Search(ctx context.Context, productType string) ([]*models.Product, error) {
	ret := _m.Called(ctx, productType)

	if len(ret) == 0 {
		panic("no return value specified for Search")
	}

	var r0 []*models.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]*models.Product, error)); ok {
		return rf(ctx, productType)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []*models.Product); ok {
		r0 = rf(ctx, productType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Product)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, productType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProductService_Search_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Search'
type ProductService_Search_Call struct {
	*mock.Call
}

// Search is a helper method to define mock.On call
//   - ctx context.Context
//   - productType string
func (_e *ProductService_Expecter) Search(ctx interface{}, productType interface{}) *ProductService_Search_Call {
	return &ProductService_Search_Call{Call: _e.mock.On("Search", ctx, productType)}
}

func (_c *ProductService_Search_Call) Run(run func(ctx context.Context, productType string)) *ProductService_Search_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *ProductService_Search_Call) Return(_a0 []*models.Product, _a1 error) *ProductService_Search_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ProductService_Search_Call) RunAndReturn(run func(context.Context, string) ([]*models.Product, error)) *ProductService_Search_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, product
func (_m *ProductService) Update(ctx context.Context, product *models.Product) (*models.Product, error) {
	ret := _m.Called(ctx, product)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *models.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Product) (*models.Product, error)); ok {
		return rf(ctx, product)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *models.Product) *models.Product); ok {
		r0 = rf(ctx, product)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Product)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *models.Product) error); ok {
		r1 = rf(ctx, product)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProductService_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type ProductService_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - product *models.Product
func (_e *ProductService_Expecter) Update(ctx interface{}, product interface{}) *ProductService_Update_Call {
	return &ProductService_Update_Call{Call: _e.mock.On("Update", ctx, product)}
}

func (_c *ProductService_Update_Call) Run(run func(ctx context.Context, product *models.Product)) *ProductService_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*models.Product))
	})
	return _c
}

func (_c *ProductService_Update_Call) Return(_a0 *models.Product, _a1 error) *ProductService_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ProductService_Update_Call) RunAndReturn(run func(context.Context, *models.Product) (*models.Product, error)) *ProductService_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewProductService creates a new instance of ProductService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewProductService(t interface {
	mock.TestingT
	Cleanup(func())
},
) *ProductService {
	mock := &ProductService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}