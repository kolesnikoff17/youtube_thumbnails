// Code generated by mockery v2.14.0. DO NOT EDIT.

package repomock

import (
	context "context"
	entity "thumbs/server/internal/entity"

	mock "github.com/stretchr/testify/mock"
)

// ThumbRepo is an autogenerated mock type for the ThumbRepo type
type ThumbRepo struct {
	mock.Mock
}

// Get provides a mock function with given fields: ctx, id
func (_m *ThumbRepo) Get(ctx context.Context, id string) (entity.Pic, error) {
	ret := _m.Called(ctx, id)

	var r0 entity.Pic
	if rf, ok := ret.Get(0).(func(context.Context, string) entity.Pic); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(entity.Pic)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Put provides a mock function with given fields: ctx, pic
func (_m *ThumbRepo) Put(ctx context.Context, pic entity.Pic) error {
	ret := _m.Called(ctx, pic)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.Pic) error); ok {
		r0 = rf(ctx, pic)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewThumbRepo interface {
	mock.TestingT
	Cleanup(func())
}

// NewThumbRepo creates a new instance of ThumbRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewThumbRepo(t mockConstructorTestingTNewThumbRepo) *ThumbRepo {
	mock := &ThumbRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}