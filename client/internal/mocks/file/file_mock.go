// Code generated by mockery v2.14.0. DO NOT EDIT.

package filemock

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// ThumbFile is an autogenerated mock type for the ThumbFile type
type ThumbFile struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, id, data
func (_m *ThumbFile) Create(ctx context.Context, id string, data []byte) (string, error) {
	ret := _m.Called(ctx, id, data)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string, []byte) string); ok {
		r0 = rf(ctx, id, data)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, []byte) error); ok {
		r1 = rf(ctx, id, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewThumbFile interface {
	mock.TestingT
	Cleanup(func())
}

// NewThumbFile creates a new instance of ThumbFile. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewThumbFile(t mockConstructorTestingTNewThumbFile) *ThumbFile {
	mock := &ThumbFile{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
