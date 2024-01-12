// Code generated by mockery v2.39.2. DO NOT EDIT.

package mocks

import (
	model "main/src/domain/model"

	mock "github.com/stretchr/testify/mock"
)

// StreamService is an autogenerated mock type for the StreamService type
type StreamService struct {
	mock.Mock
}

// CreateStream provides a mock function with given fields: _a0
func (_m *StreamService) CreateStream(_a0 *model.Stream) (*model.Stream, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for CreateStream")
	}

	var r0 *model.Stream
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.Stream) (*model.Stream, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(*model.Stream) *model.Stream); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Stream)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.Stream) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllStream provides a mock function with given fields:
func (_m *StreamService) GetAllStream() ([]model.Stream, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAllStream")
	}

	var r0 []model.Stream
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]model.Stream, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []model.Stream); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Stream)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewStreamService creates a new instance of StreamService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStreamService(t interface {
	mock.TestingT
	Cleanup(func())
}) *StreamService {
	mock := &StreamService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
