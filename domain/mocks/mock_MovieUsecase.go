// Code generated by mockery v2.46.2. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "recommand-chat-bot/domain"

	mock "github.com/stretchr/testify/mock"
)

// MockMovieUsecase is an autogenerated mock type for the MovieUsecase type
type MockMovieUsecase struct {
	mock.Mock
}

type MockMovieUsecase_Expecter struct {
	mock *mock.Mock
}

func (_m *MockMovieUsecase) EXPECT() *MockMovieUsecase_Expecter {
	return &MockMovieUsecase_Expecter{mock: &_m.Mock}
}

// Store provides a mock function with given fields: ctx, m
func (_m *MockMovieUsecase) Store(ctx context.Context, m *domain.Movie) (int64, error) {
	ret := _m.Called(ctx, m)

	if len(ret) == 0 {
		panic("no return value specified for Store")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Movie) (int64, error)); ok {
		return rf(ctx, m)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Movie) int64); ok {
		r0 = rf(ctx, m)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *domain.Movie) error); ok {
		r1 = rf(ctx, m)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockMovieUsecase_Store_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Store'
type MockMovieUsecase_Store_Call struct {
	*mock.Call
}

// Store is a helper method to define mock.On call
//   - ctx context.Context
//   - m *domain.Movie
func (_e *MockMovieUsecase_Expecter) Store(ctx interface{}, m interface{}) *MockMovieUsecase_Store_Call {
	return &MockMovieUsecase_Store_Call{Call: _e.mock.On("Store", ctx, m)}
}

func (_c *MockMovieUsecase_Store_Call) Run(run func(ctx context.Context, m *domain.Movie)) *MockMovieUsecase_Store_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*domain.Movie))
	})
	return _c
}

func (_c *MockMovieUsecase_Store_Call) Return(_a0 int64, _a1 error) *MockMovieUsecase_Store_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockMovieUsecase_Store_Call) RunAndReturn(run func(context.Context, *domain.Movie) (int64, error)) *MockMovieUsecase_Store_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockMovieUsecase creates a new instance of MockMovieUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockMovieUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockMovieUsecase {
	mock := &MockMovieUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
