// Code generated by mockery v2.34.2. DO NOT EDIT.

package mocks

import (
	model "main/src/domain/model"

	mock "github.com/stretchr/testify/mock"
)

// StreamRepository is an autogenerated mock type for the StreamRepository type
type StreamRepository struct {
	mock.Mock
}

// GetAllStream provides a mock function with given fields:
func (_m *StreamRepository) GetAllStream() ([]model.Stream, error) {
	ret := _m.Called()

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

// NewStreamRepository creates a new instance of StreamRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStreamRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *StreamRepository {
	mock := &StreamRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}