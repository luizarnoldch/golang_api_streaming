// Code generated by mockery v2.39.2. DO NOT EDIT.

package mocks

import (
	model "main/src/streams/domain/model"
	error "main/utils/error"

	mock "github.com/stretchr/testify/mock"
)

// StreamRepository is an autogenerated mock type for the StreamRepository type
type StreamRepository struct {
	mock.Mock
}

// CreateStream provides a mock function with given fields: _a0
func (_m *StreamRepository) CreateStream(_a0 *model.Stream) (*model.Stream, *error.Error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for CreateStream")
	}

	var r0 *model.Stream
	var r1 *error.Error
	if rf, ok := ret.Get(0).(func(*model.Stream) (*model.Stream, *error.Error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(*model.Stream) *model.Stream); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Stream)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.Stream) *error.Error); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*error.Error)
		}
	}

	return r0, r1
}

// DeleteStream provides a mock function with given fields: _a0
func (_m *StreamRepository) DeleteStream(_a0 string) *error.Error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for DeleteStream")
	}

	var r0 *error.Error
	if rf, ok := ret.Get(0).(func(string) *error.Error); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*error.Error)
		}
	}

	return r0
}

// GetAllStream provides a mock function with given fields:
func (_m *StreamRepository) GetAllStream() ([]model.Stream, *error.Error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAllStream")
	}

	var r0 []model.Stream
	var r1 *error.Error
	if rf, ok := ret.Get(0).(func() ([]model.Stream, *error.Error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []model.Stream); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Stream)
		}
	}

	if rf, ok := ret.Get(1).(func() *error.Error); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*error.Error)
		}
	}

	return r0, r1
}

// GetStreamById provides a mock function with given fields: _a0
func (_m *StreamRepository) GetStreamById(_a0 string) (*model.Stream, *error.Error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for GetStreamById")
	}

	var r0 *model.Stream
	var r1 *error.Error
	if rf, ok := ret.Get(0).(func(string) (*model.Stream, *error.Error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(string) *model.Stream); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Stream)
		}
	}

	if rf, ok := ret.Get(1).(func(string) *error.Error); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*error.Error)
		}
	}

	return r0, r1
}

// UpdateStreamById provides a mock function with given fields: stream_id, stream
func (_m *StreamRepository) UpdateStreamById(stream_id string, stream *model.Stream) (*model.Stream, *error.Error) {
	ret := _m.Called(stream_id, stream)

	if len(ret) == 0 {
		panic("no return value specified for UpdateStreamById")
	}

	var r0 *model.Stream
	var r1 *error.Error
	if rf, ok := ret.Get(0).(func(string, *model.Stream) (*model.Stream, *error.Error)); ok {
		return rf(stream_id, stream)
	}
	if rf, ok := ret.Get(0).(func(string, *model.Stream) *model.Stream); ok {
		r0 = rf(stream_id, stream)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Stream)
		}
	}

	if rf, ok := ret.Get(1).(func(string, *model.Stream) *error.Error); ok {
		r1 = rf(stream_id, stream)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*error.Error)
		}
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
